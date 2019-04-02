package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
)

func NewVm() *cli.Command {
	vm := &cli.Command{
		Name: "vm",
		Short: i18n.T("", "云主机"),
		Usage: "vm ",
		Run: func(ctx *cli.Context, args []string) error {
			if len(args) > 0 {
				return cli.NewInvalidCommandError(args[0], ctx)
			}
			//profileName, _ := config.ProfileFlag(ctx.Flags()).GetValue()
			//return doVm(ctx, profileName, args)
			return doVm(ctx, args)
		},
		Hidden: false,
	}
	vm.AddSubCommand(NewDescribeImage())
	vm.AddSubCommand(NewDescribeImageMembers())
	vm.AddSubCommand(NewDescribeImageConstraints())
	vm.AddSubCommand(NewDescribeImageConstraintsBatch())
	vm.AddSubCommand(NewDescribeImages())
	vm.AddSubCommand(NewDescribeInstance())
	vm.AddSubCommand(NewDescribeInstancePrivateIpAddress())
	vm.AddSubCommand(NewDescribeInstanceStatus())
	vm.AddSubCommand(NewDescribeInstanceType())
	vm.AddSubCommand(NewDescribeInstanceVncUrl())
	vm.AddSubCommand(NewDescribeInstances())
	vm.AddSubCommand(NewDescribeKeypairs())
	vm.AddSubCommand(NewDescribeQuotas())
	return vm
}

func doVm(ctx *cli.Context, args []string) error{
	if len(args) == 0 {
		NewVm().ExecuteHelp(ctx, args)
	}
	return nil
}
