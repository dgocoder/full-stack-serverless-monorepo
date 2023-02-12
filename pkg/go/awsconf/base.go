package awsconf

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func newAWSConfig(ctx context.Context) (aws.Config, error) {
	region := os.Getenv("AWS_REGION")

	awsConfig, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return aws.Config{}, err
	}

	return awsConfig, nil
}
