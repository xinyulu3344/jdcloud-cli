package disk

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
)

func NewDisk() *cli.Command{
	disk := &cli.Command{
		Name: "disk",
		Short: i18n.T("", "云硬盘API"),
		Usage: "disk [-h] [--output {json}] [-v] [subcommand]",
		Run: func(ctx *cli.Context, args []string) error {

			if len(args) > 0 {
				return cli.NewInvalidCommandError(args[0], ctx)
			}
			return doDisk(ctx, args)
		},
		Hidden: false,
	}
	disk.AddSubCommand(NewDescribeDisk())
	disk.AddSubCommand(NewDescribeDisks())
	disk.AddSubCommand(NewDescribeSnapshot())
	disk.AddSubCommand(NewDescribeSnapshots())
	return disk
}

func doDisk(ctx *cli.Context, args []string) error {

	if len(args) == 0 {
		NewDisk().ExecuteHelp(ctx, args)
	}
	return nil
}
