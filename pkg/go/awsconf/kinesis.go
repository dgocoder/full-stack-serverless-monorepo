package awsconf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

func NewBaseKinesisConf(ctx context.Context) (*kinesis.Client, error) {
	conf, err := newAWSConfig(ctx)
	if err != nil {
		return nil, err
	}

	return kinesis.NewFromConfig(conf), nil
}
