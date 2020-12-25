package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Sg struct {
	SecurityGroupEc2 interface{}
	SecurityGroupAlb interface{}
}

func createSg(ctx *pulumi.Context, opt *Opt) (*Sg, error) {
	// EC2 SecurityGroup
	sg1, err := ec2.NewSecurityGroup(ctx, "app-dev-ec2-sg", &ec2.SecurityGroupArgs{
		VpcId: opt.vpc.(*ec2.Vpc).ID(),
		Ingress: ec2.SecurityGroupIngressArray{
			ec2.SecurityGroupIngressArgs{
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				Protocol:   pulumi.String("TCP"),
			},
			ec2.SecurityGroupIngressArgs{
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				FromPort:   pulumi.Int(80),
				ToPort:     pulumi.Int(80),
				Protocol:   pulumi.String("TCP"),
			},
		},
		Egress: ec2.SecurityGroupEgressArray{
			ec2.SecurityGroupEgressArgs{
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				FromPort:   pulumi.Int(0),
				ToPort:     pulumi.Int(0),
				Protocol:   pulumi.String("-1"),
			},
		},
		Tags: pulumi.StringMap{
			"Name": pulumi.String("app-dev-sg1"),
		},
	})
	if err != nil {
		return nil, err
	}

	// ALB SecurityGroup
	sg2, err := ec2.NewSecurityGroup(ctx, "app-dev-lb-sg", &ec2.SecurityGroupArgs{
		VpcId: opt.vpc.(*ec2.Vpc).ID(),
		Ingress: ec2.SecurityGroupIngressArray{
			ec2.SecurityGroupIngressArgs{
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				FromPort:   pulumi.Int(80),
				ToPort:     pulumi.Int(80),
				Protocol:   pulumi.String("TCP"),
			},
		},
		Egress: ec2.SecurityGroupEgressArray{
			ec2.SecurityGroupEgressArgs{
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				FromPort:   pulumi.Int(0),
				ToPort:     pulumi.Int(0),
				Protocol:   pulumi.String("-1"),
			},
		},
		Tags: pulumi.StringMap{
			"Name": pulumi.String("app-dev-sg2"),
		},
	})
	if err != nil {
		return nil, err
	}

	sg := new(Sg)
	sg.SecurityGroupEc2 = sg1
	sg.SecurityGroupAlb = sg2

	return sg, nil
}
