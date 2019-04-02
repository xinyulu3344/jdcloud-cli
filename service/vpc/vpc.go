package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
)

func NewVpc() *cli.Command {
	vpc := &cli.Command{
		Name: "vpc",
		Short: i18n.T("", "vpc cli 子命令，VPC相关API。"),
		Usage: "vpc [-h] [--output {json}] [-v] [subcommand]",
		Run: func(ctx *cli.Context, args []string) error {

			if len(args) > 0 {
				return cli.NewInvalidCommandError(args[0], ctx)
			}
			return doVpc(ctx, args)
		},
		Hidden: false,
	}

	vpc.AddSubCommand(NewDescribeElasticIp())
	vpc.AddSubCommand(NewDescribeElasticIps())
	vpc.AddSubCommand(NewDescribeNetworkAcl())
	vpc.AddSubCommand(NewDescribeNetworkAcls())
	vpc.AddSubCommand(NewDescribeNetworkInterface())
	vpc.AddSubCommand(NewDescribeNetworkInterfaces())
	vpc.AddSubCommand(NewDescribeNetworkSecurityGroup())
	vpc.AddSubCommand(NewDescribeNetworkSecurityGroups())
	vpc.AddSubCommand(NewDescribeQuota())
	vpc.AddSubCommand(NewDescribeRouteTable())
	vpc.AddSubCommand(NewDescribeRouteTables())
	vpc.AddSubCommand(NewDescribeSubnet())
	vpc.AddSubCommand(NewDescribeSubnets())
	vpc.AddSubCommand(NewDescribeVpc())
	vpc.AddSubCommand(NewDescribeVpcs())
	vpc.AddSubCommand(NewDescribeVpcPeering())
	vpc.AddSubCommand(NewDescribeVpcPeerings())
	return vpc
}

func doVpc(ctx *cli.Context, args []string) error{

	if len(args) == 0 {
		NewVpc().ExecuteHelp(ctx, args)
	}
	return nil
}
