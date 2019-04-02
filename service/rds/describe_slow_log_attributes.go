package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
	"strconv"
)

func NewDescribeSlowLogAttributes() *cli.Command {
	describeSlowLogAttributes := &cli.Command{
		Name: "describe-slow-log-attributes",
		Short: i18n.T("", "查询MySQL实例的慢日志的详细信息。<br>- 仅支持SQL Server。"),
		Usage: "describe-slow-log-attributes [-h] [--region-id REGIONID] --instance-id INSTANCEID --start-time STARTTIME --end-time ENDTIME [--db-name DBNAME] [--page-number PAGENUMBER] [--page-size PAGESIZE]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeSlowLogAttributes(ctx)
			return nil
		},
	}

	describeSlowLogAttributes.Flags().Add(regionIdFlag)
	describeSlowLogAttributes.Flags().Add(instanceIdFlag)
	describeSlowLogAttributes.Flags().Add(slowStartTimeFlag)
	describeSlowLogAttributes.Flags().Add(slowEndTimeFlag)
	describeSlowLogAttributes.Flags().Add(slowDbNameFlag)
	describeSlowLogAttributes.Flags().Add(pageNumberFlag)
	describeSlowLogAttributes.Flags().Add(pageSizeFlag)
	describeSlowLogAttributes.Flags().Add(inputJsonFlag)
	describeSlowLogAttributes.Flags().Add(headersFlag)
	return describeSlowLogAttributes
}

func doDescribeSlowLogAttributes(ctx *cli.Context) {

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


	req := apis.NewDescribeSlowLogAttributesRequestWithAllParams(regionId, instanceId, slowStartTime, slowEndTime, &slowDbName, &pageNumber, &pageSize)

	resp, err := rdsClient.DescribeSlowLogAttributes(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
