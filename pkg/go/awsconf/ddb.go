package awsconf

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"

	ddbconversions "github.com/aereal/go-dynamodb-attribute-conversions/v2"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	snstypes "github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/dgocoder/full-stack-serverless-monorepo/pkg/go/logger"
	"github.com/dgocoder/full-stack-serverless-monorepo/pkg/go/stdconv"
)

func NewBaseDynamoDBConf(ctx context.Context) (*dynamodb.Client, error) {
	conf, err := NewAWSConfig(ctx)
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(conf), nil
}

type (
	DBStreamEntities struct {
		Entities map[string]DBStreamEntity
	}
	DBStreamEntity struct {
		EntityTopic string
		ObjectType  DomainEntity
		EventSource *string
		Events      map[string]TopicEventName
	}
	DomainEntity interface {
		GetEntityType() string
	}
	TopicEventName struct {
		EventName       string
		TopicARN        string
		MessageAttrFunc MessageAttrFunc
	}
	MessageAttrFunc   func(data interface{}) (*map[string]snstypes.MessageAttributeValue, error)
	StreamEntityInput struct {
		EventSource  string
		DomainEntity DomainEntity
		EventTypes   []DDBStreamEventType
		InsertEvent  *EntityEvent
		ModifyEvent  *EntityEvent
		RemoveEvent  *EntityEvent
	}
	StreamEntity struct {
		StreamEntityInput
		EventSourceServiceName string
		eventMap               *map[string]TopicEventName
	}
	DDBStreamEventType int
	EntityEvent        struct {
		MessageAttrFunc *MessageAttrFunc
		EventName       string
		TopicARN        string
		DomainEntity    DomainEntity
	}
)

const (
	DDBStreamEventTypeInsert DDBStreamEventType = iota
	DDBStreamEventTypeModify
	DDBStreamEventTypeRemove
	DDBStreamEventTypeAll
)

// DDBEntityToSNS takes the event records from a dynamodb stream
// converts them to entity shape and publishes them to sns.
func DDBEntityToSNS(ctx context.Context, event events.DynamoDBEvent, entities []StreamEntity) error {
	entityMap := map[string]DBStreamEntity{}

	for _, streamEntity := range entities {
		entry := DBStreamEntity{
			EventSource: stdconv.Ptr(streamEntity.EventSourceServiceName),
			EntityTopic: streamEntity.DomainEntity.GetEntityType(),
			Events:      *streamEntity.eventMap,
			ObjectType:  streamEntity.DomainEntity,
		}

		entityMap[entry.EntityTopic] = entry
	}

	entityMaps := DBStreamEntities{Entities: entityMap}
	for _, record := range event.Records {
		entity := entityMaps.getByRecord(record)
		topic := entity.getTopicEvent(record.EventName)

		// Skip records whose topic handling isn't defined on entity
		if topic.isEmpty() || entity.IsEmpty() {
			// logger.Warn("Skipping DynamoDBEvent Record", map[string]interface{}{"isEmpty": topic.isEmpty(), "entityIsEmpty": entity.IsEmpty()})

			continue
		}

		changes, err := unmarshalChanges(entity, record)
		if err != nil {
			return err
		}

		err = sendSNSEntityChanges(ctx, record, topic, entity, changes.previous, changes.update)
		if err != nil {
			return err
		}
	}

	return nil
}

type unmarshaledChanges struct {
	update   interface{}
	previous interface{}
}

func unmarshalChanges(entity DBStreamEntity, record events.DynamoDBEventRecord) (unmarshaledChanges, error) {
	changes := unmarshaledChanges{}

	if record.Change.NewImage == nil && record.Change.OldImage == nil {
		return changes, nil
	}

	t := reflect.TypeOf(entity.ObjectType)

	if record.Change.NewImage != nil {
		upAttr := ddbconversions.AttributeValueMapFrom(record.Change.NewImage)
		update := reflect.New(t).Interface()

		err := attributevalue.UnmarshalMapWithOptions(upAttr, &update, func(options *attributevalue.DecoderOptions) {
			options.TagKey = "json"
		})
		if err != nil {
			return unmarshaledChanges{}, errors.New("failed to unmarshal new image change for db stream")
		}

		changes.update = update
	}

	if record.Change.OldImage != nil {
		prevAttr := ddbconversions.AttributeValueMapFrom(record.Change.OldImage)
		previous := reflect.New(t).Interface()

		err := attributevalue.UnmarshalMapWithOptions(prevAttr, &previous, func(options *attributevalue.DecoderOptions) {
			options.TagKey = "json"
		})
		if err != nil {
			return unmarshaledChanges{}, errors.New("failed to unmarshal old image change for db stream")
		}

		changes.previous = previous
	}

	return changes, nil
}

func (dse DBStreamEntities) getByRecord(record events.DynamoDBEventRecord) DBStreamEntity {
	var recordImage map[string]events.DynamoDBAttributeValue

	switch events.DynamoDBOperationType(record.EventName) {
	case events.DynamoDBOperationTypeInsert, events.DynamoDBOperationTypeModify:
		recordImage = record.Change.NewImage
	case events.DynamoDBOperationTypeRemove:
		recordImage = record.Change.OldImage
	default:
		return DBStreamEntity{}
	}

	if entityType, ok := recordImage["entity_type"]; ok {
		if entryFromEntityMap, ok := dse.Entities[entityType.String()]; ok {
			return entryFromEntityMap
		}
	}

	return DBStreamEntity{}
}

func (de *DBStreamEntity) IsEmpty() bool {
	objectType := de.ObjectType

	return len(de.Events) == 0 || (objectType == nil || (reflect.ValueOf(objectType).Kind() == reflect.Ptr && reflect.ValueOf(objectType).IsNil()))
}

func (de DBStreamEntity) getTopicEvent(eventName string) TopicEventName {
	if topicEventName, ok := de.Events[eventName]; ok {
		return topicEventName
	}

	return TopicEventName{}
}

func (t *TopicEventName) isEmpty() bool {
	return t.EventName == "" && t.TopicARN == ""
}

type DBBChange struct {
	// The item in the DynamoDB table as it appeared after it was modified.
	NewImage interface{} `json:"new_image"`

	// The item in the DynamoDB table as it appeared before it was modified.
	OldImage interface{} `json:"old_image"`
}
type DBStreamSNS struct {
	EventID     string    `json:"event_id"`
	EventName   string    `json:"event_name"`
	EventSource string    `json:"event_source"`
	Change      DBBChange `json:"change"`
}

func sendSNSEntityChanges(ctx context.Context, record events.DynamoDBEventRecord, topic TopicEventName, entity DBStreamEntity, previous interface{}, update interface{}) error {
	event := DBStreamSNS{
		EventID:     record.EventID,
		EventName:   topic.EventName,
		EventSource: *entity.EventSource,
		Change: DBBChange{
			NewImage: update,
			OldImage: previous,
		},
	}

	// z.Info("publishing event", zap.Any("metadata", map[string]interface{}{"topic": topic.TopicARN, "event": event}))
	// log.Info().Any("metadata", map[string]interface{}{"topic": topic.TopicARN, "event": event}).Msg("publishing event")
	logger.Info("publishing event", nil)
	logger.Info("publishing event", map[string]interface{}{"topic": topic.TopicARN, "event": event})

	msg, err := json.Marshal(event)
	if err != nil {
		return err
	}

	ms := string(msg)

	input := sns.PublishInput{
		Message:           &ms,
		TopicArn:          &topic.TopicARN,
		MessageAttributes: map[string]snstypes.MessageAttributeValue{},
	}

	if topic.MessageAttrFunc != nil {
		obj, err := topic.MessageAttrFunc(update)
		if err != nil {
			return err
		}

		input.MessageAttributes = *obj
	}

	client, err := NewSNSConf(ctx)
	if err != nil {
		return err
	}

	_, err = client.Publish(ctx, &input)
	if err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}

	return nil
}

func NewStreamEntity(input StreamEntityInput) StreamEntity {
	se := input.ToStreamEntity()

	eventCount := len(se.EventTypes)
	if eventCount == 0 {
		return StreamEntity{}
	}

	for _, eventType := range input.EventTypes {
		if eventCount == 1 && eventType == DDBStreamEventTypeAll {
			se.setInsertEvent().setModifyEvent().setRemoveEvent()

			continue
		}

		if eventType == DDBStreamEventTypeInsert {
			se.setInsertEvent()

			continue
		}

		if eventType == DDBStreamEventTypeModify {
			se.setModifyEvent()

			continue
		}

		if eventType == DDBStreamEventTypeRemove {
			se.setRemoveEvent()

			continue
		}
	}

	return se
}

func (sei StreamEntityInput) ToStreamEntity() StreamEntity {
	return StreamEntity{
		StreamEntityInput: StreamEntityInput{
			EventSource:  sei.EventSource,
			DomainEntity: sei.DomainEntity,
			EventTypes:   sei.EventTypes,
			InsertEvent:  sei.InsertEvent,
			ModifyEvent:  sei.ModifyEvent,
			RemoveEvent:  sei.RemoveEvent,
		},
		EventSourceServiceName: fmt.Sprintf("%s_SERVICE", sei.EventSource),
		eventMap:               &map[string]TopicEventName{},
	}
}

func (se *StreamEntity) setInsertEvent() *StreamEntity {
	se.InsertEvent = NewEntityEvent(*se, se.InsertEvent, "CREATED")

	return se.setTopicEventName(se.InsertEvent, DDBStreamEventTypeInsert)
}

func (se *StreamEntity) setModifyEvent() *StreamEntity {
	se.ModifyEvent = NewEntityEvent(*se, se.ModifyEvent, "UPDATED")

	return se.setTopicEventName(se.ModifyEvent, DDBStreamEventTypeModify)
}

func (se *StreamEntity) setRemoveEvent() *StreamEntity {
	se.RemoveEvent = NewEntityEvent(*se, se.RemoveEvent, "DELETED")

	return se.setTopicEventName(se.RemoveEvent, DDBStreamEventTypeRemove)
}

func NewEntityEvent(se StreamEntity, event *EntityEvent, action string) *EntityEvent {
	eventData := &EntityEvent{
		EventName: fmt.Sprintf("%s_%s", se.DomainEntity.GetEntityType(), action),
		TopicARN:  os.Getenv(fmt.Sprintf("%s_%s_TOPIC_ARN", se.DomainEntity.GetEntityType(), action)),
	}

	if event != nil {
		eventData.MessageAttrFunc = event.MessageAttrFunc
	}

	return eventData
}

func (data DDBStreamEventType) String() string {
	return [...]string{"INSERT", "MODIFY", "REMOVE", "ALL"}[data]
}

func (se *StreamEntity) setTopicEventName(entityEvent *EntityEvent, eventType DDBStreamEventType) *StreamEntity {
	if entityEvent == nil {
		return se
	}

	topicEvent := TopicEventName{
		EventName: entityEvent.EventName,
		TopicARN:  entityEvent.TopicARN,
	}

	if entityEvent.MessageAttrFunc != nil {
		topicEvent.MessageAttrFunc = *entityEvent.MessageAttrFunc
	}

	(*se.eventMap)[eventType.String()] = topicEvent

	return se
}
