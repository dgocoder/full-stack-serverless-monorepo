package awsconf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func NewSNSConf(ctx context.Context) (*sns.Client, error) {
	conf, err := newAWSConfig(ctx)
	if err != nil {
		return nil, err
	}

	return sns.NewFromConfig(conf), nil
}
