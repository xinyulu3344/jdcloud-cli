package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
)

func NewRds() *cli.Command {
	rds := &cli.Command{
		Name: "rds",
		Short: i18n.T("", "rds cli 子命令，目前RDS OpenAPI支持云数据库SQL Server，可以通过OpenAPI实现数据库管理，账号管理，备份管理，单库上云等功能，后续将支持云数据库MySQL。"),
		Usage: "usage: jdc rds [-h] [--output {json}] [-v] [subcommand]",
		Run: func(ctx *cli.Context, args []string) error {

			if len(args) > 0 {
				return cli.NewInvalidCommandError(args[0], ctx)
			}
			return doRds(ctx, args)
		},
		Hidden: false,
	}

	rds.AddSubCommand(NewDescribeAccounts())
	rds.AddSubCommand(NewDescribeAudit())
	rds.AddSubCommand(NewDescribeAuditDownloadUrl())
	rds.AddSubCommand(NewDescribeAuditFiles())
	rds.AddSubCommand(NewDescribeAuditOptions())
	rds.AddSubCommand(NewDescribeBackupDownloadUrl())
	rds.AddSubCommand(NewDescribeBackupPolicy())
	rds.AddSubCommand(NewDescribeBackups())
	rds.AddSubCommand(NewDescribeBinlogDownloadUrl())
	rds.AddSubCommand(NewDescribeBinlogs())
	rds.AddSubCommand(NewDescribeDatabases())
	rds.AddSubCommand(NewDescribeErrorLogs())
	rds.AddSubCommand(NewDescribeImportFiles())
	rds.AddSubCommand(NewDescribeIndexPerformance())
	rds.AddSubCommand(NewDescribeInstanceAttributes())
	rds.AddSubCommand(NewDescribeInstances())
	rds.AddSubCommand(NewDescribeQueryPerformance())
	rds.AddSubCommand(NewDescribeSlowLogAttributes())
	rds.AddSubCommand(NewDescribeSlowLogs())
	rds.AddSubCommand(NewDescribeWhiteList())
	return rds
}

func doRds(ctx *cli.Context, args []string) error {

	if len(args) == 0 {
		NewRds().ExecuteHelp(ctx, args)
	}
	return nil
}