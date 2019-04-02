package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"strconv"
	"fmt"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
)

func NewDescribeSlowLogs() *cli.Command {
	describeSlowLogs := &cli.Command{
		Name: "describe-slow-logs",
		Short: i18n.T("", "查询MySQL实例的慢日志的概要信息。<br>- 仅支持SQL Server。"),
		Usage: "describe-slow-logs [-h] [--region-id REGIONID] --instance-id INSTANCEID --start-time STARTTIME --end-time ENDTIME [--db-name DBNAME] [--page-number PAGENUMBER] [--page-size PAGESIZE]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeSlowLogs(ctx)
			return nil
		},
	}

	describeSlowLogs.Flags().Add(regionIdFlag)
	describeSlowLogs.Flags().Add(instanceIdFlag)
	describeSlowLogs.Flags().Add(slowStartTimeFlag)
	describeSlowLogs.Flags().Add(slowEndTimeFlag)
	describeSlowLogs.Flags().Add(slowDbNameFlag)
	describeSlowLogs.Flags().Add(pageNumberFlag)
	describeSlowLogs.Flags().Add(pageSizeFlag)
	describeSlowLogs.Flags().Add(inputJsonFlag)
	describeSlowLogs.Flags().Add(headersFlag)
	return describeSlowLogs
}

func doDescribeSlowLogs(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()
	slowStartTime, _ := slowStartTimeFlag.GetValue()
	slowEndTime, _ := slowEndTimeFlag.GetValue()
	slowDbName, _ := slowDbNameFlag.GetValue()

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

	fmt.Println(regionId, instanceId, slowStartTime, slowEndTime, slowDbName, pageNumber, pageSize)

	req := apis.NewDescribeSlowLogsRequestWithAllParams(regionId, instanceId, slowStartTime, slowEndTime, &slowDbName, &pageNumber, &pageSize)

	resp, err := rdsClient.DescribeSlowLogs(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
