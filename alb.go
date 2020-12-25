package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func createAlb(ctx *pulumi.Context, opt *Opt, sg *Sg) error {
	// alb(subnet)
	alb, err := lb.NewLoadBalancer(ctx, "app-dev-alb", &lb.LoadBalancerArgs{
		SecurityGroups: pulumi.StringArray{sg.SecurityGroupAlb.(*ec2.SecurityGroup).ID()},
		Subnets: pulumi.StringArray{
			opt.subnet.(*ec2.Subnet).ID(),
			opt.subnet2.(*ec2.Subnet).ID(),
		},
	})
	if err != nil {
		return err
	}

	// tg
	tg, err := lb.NewTargetGroup(ctx, "app-dev-tg", &lb.TargetGroupArgs{
		Port: pulumi.Int(80),
		Protocol: pulumi.String("HTTP"),
		VpcId: opt.vpc.(*ec2.Vpc).ID(),
	})
	if err != nil {
		return err
	}

	// listener
	_, err = lb.NewListener(ctx, "app-dev-listener", &lb.ListenerArgs{
		LoadBalancerArn: alb.Arn,
		Port: pulumi.Int(80),
		Protocol: pulumi.String("HTTP"),
		DefaultActions: lb.ListenerDefaultActionArray{
				&lb.ListenerDefaultActionArgs{
						Type: pulumi.String("forward"),
						TargetGroupArn: tg.Arn,
				},
		},
	})
	if err != nil {
		return err
	}

	return nil
}