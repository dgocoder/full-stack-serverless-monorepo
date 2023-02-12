package awsconf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewBaseS3Conf(ctx context.Context) (*s3.Client, error) {
	conf, err := newAWSConfig(ctx)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(conf), nil
}
