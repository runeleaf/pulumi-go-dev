package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func createS3(ctx *pulumi.Context) error {
	_, err := s3.NewBucket(ctx, "app-dev-bucket", nil)
	if err != nil {
		return err
	}

	return nil
}
