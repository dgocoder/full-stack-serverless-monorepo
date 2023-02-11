package awsapigw

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
)

var headerJSON = map[string]string{"Content-Type": "application/json"}

// SendValidationError returns an APIGatewayProxyResponse given a
// status and validation error are provided. This function should be
// leveraged for standard responses for validation errors.
func SendValidationError(
	status StatusCode,
	valError map[string]interface{},
) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(valError)
	if err != nil {
		return marshalAPIGatewayError()
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: int(status),
		Headers:    headerJSON,
	}, errors.New("validation error")
}

// SendError returns an APIGatewayProxyResponse given a
// status and error message are provided. This function should be
// leveraged for standard responses for errors during request.
func SendError(
	status StatusCode,
	errMsg string,
) (events.APIGatewayProxyResponse, error) {
	errMap := map[string]string{
		"error": errMsg,
	}

	body, err := json.Marshal(errMap)
	if err != nil {
		return marshalAPIGatewayError()
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: int(status),
		Headers:    headerJSON,
	}, errors.New(errMsg)
}

// SendResponse returns an APIGatewayProxyResponse given a
// status and data object. This function should be
// leveraged for standard responses for requests.
func SendResponse(
	status StatusCode,
	data interface{},
) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return marshalAPIGatewayError()
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: int(status),
		Headers:    headerJSON,
	}, nil
}

// GetParamPath returns the string value of the request unescaped.
func GetParamPath(param string, req *events.APIGatewayProxyRequest) (string, error) {
	rawParam1, found := req.PathParameters[param]
	if !found {
		return "", fmt.Errorf("param not found: %s", param)
	}

	qp, err := url.QueryUnescape(rawParam1)
	if err != nil {
		return "", fmt.Errorf("param invalid format: %s", param)
	}

	return qp, nil
}

func marshalAPIGatewayError() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers:    headerJSON,
	}, errors.New("failed to marshal validation errors")
}
