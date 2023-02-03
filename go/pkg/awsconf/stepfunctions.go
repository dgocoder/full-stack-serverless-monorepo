package awsconf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

func NewBaseStepFunctionConf(ctx context.Context) (*sfn.Client, error) {
	conf, err := NewAWSConfig(ctx)
	if err != nil {
		return nil, err
	}

	return sfn.NewFromConfig(conf), nil
}
