package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
	"strconv"
)

func NewDescribeBackups() *cli.Command {
	describeBackups := &cli.Command{
		Name: "describe-backups",
		Short: i18n.T("", "查看该RDS实例下所有备份的详细信息，返回的备份列表按照备份开始时间（backupStartTime）降序排列。。"),
		Usage: `describe-backups [-h] [--region-id REGIONID] --instance-id INSTANCEID [--auto AUTO] [--backup-type-filter BACKUPTYPEFILTER] [--db-name-filter DBNAMEFILTER]
                                [--backup-time-range-start-filter BACKUPTIMERANGESTARTFILTER] [--backup-time-range-end-filter BACKUPTIMERANGEENDFILTER] --page-number PAGENUMBER --page-size PAGESIZE
                                [--input-json INPUT_JSON] [--headers HEADERS]`,
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeBackups(ctx)
			return nil
		},
	}

	describeBackups.Flags().Add(regionIdFlag)
	describeBackups.Flags().Add(instanceIdFlag)
	describeBackups.Flags().Add(autoFlag)
	describeBackups.Flags().Add(backupTypeFilterFlag)
	describeBackups.Flags().Add(dbNameFilterFlag)
	describeBackups.Flags().Add(backupTimeRangeStartFilterFlag)
	describeBackups.Flags().Add(backupTimeRangeEndFilterFlag)
	describeBackups.Flags().Add(pageNumberFlag)
	describeBackups.Flags().Add(pageSizeFlag)
	describeBackups.Flags().Add(inputJsonFlag)
	describeBackups.Flags().Add(headersFlag)
	return describeBackups
}

func doDescribeBackups(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()

	autoStr, _ := autoFlag.GetValue()

	auto, err := strconv.Atoi(autoStr)
	if autoStr != "" && err != nil {
		cli.Errorf(ctx.Writer(), "autoStr strconv.Atoi err: %s", err)
	}

	backupTypeFilter, _ := backupTypeFilterFlag.GetValue()
	dbNameFilter, _ := dbNameFilterFlag.GetValue()
	backupTimeRangeStartFilter, _ := backupTimeRangeStartFilterFlag.GetValue()
	backupTimeRangeEndFilter, _ := backupTimeRangeEndFilterFlag.GetValue()

	pageNumberStr, _ := pageNumberFlag.GetValue()
	pageNumber, err := strconv.Atoi(pageNumberStr)
	if pageNumberStr != "" && err != nil {
		cli.Errorf(ctx.Writer(), "pageNumberStr strconv.Atoi err: %s", err)
	}

	pageSizeStr, _ := pageSizeFlag.GetValue()
	pageSize, err := strconv.Atoi(pageSizeStr)
	if pageSizeStr != "" && err != nil {
		cli.Errorf(ctx.Writer(), "pageSizeStr strconv.Atoi err: %s", err)
	}


	req := apis.NewDescribeBackupsRequestWithAllParams(regionId, instanceId, &auto, &backupTypeFilter, &dbNameFilter, &backupTimeRangeStartFilter, &backupTimeRangeEndFilter, pageNumber, pageSize)

	resp, err := rdsClient.DescribeBackups(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
