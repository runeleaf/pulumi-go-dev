package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func createEc2(ctx *pulumi.Context, opt *Opt, sg *Sg) error {
	// Instance
	_, err := ec2.NewInstance(ctx, "app-dev-web", &ec2.InstanceArgs{
		Ami: pulumi.String("ami-01748a72bed07727c"), // Amazon Linux 2 AMI
		InstanceType: pulumi.String("t3.micro"),
		SubnetId: opt.subnet.(*ec2.Subnet).ID(),
		VpcSecurityGroupIds: pulumi.StringArray{sg.SecurityGroupEc2.(*ec2.SecurityGroup).ID()},
		KeyName: pulumi.String("app-dev"),
		AssociatePublicIpAddress: pulumi.Bool(true),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("app-dev-web1"),
		},
	})
	if err != nil {
		return err
	}

	return nil
}
