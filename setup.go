package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func setup(ctx *pulumi.Context) error {
	// VPC
	opt, err := createVpc(ctx)
	if err != nil {
		return err
	}

	// SecurityGroup
	sg, err := createSg(ctx, opt)
	if err != nil {
		return err
	}

	// S3
	err = createS3(ctx)
	if err != nil {
		return err
	}

	// EC2
	err = createEc2(ctx, opt, sg)
	if err != nil {
		return err
	}

	// ALB
	err = createAlb(ctx, opt, sg)
	if err != nil {
		return err
	}

	return nil
}
