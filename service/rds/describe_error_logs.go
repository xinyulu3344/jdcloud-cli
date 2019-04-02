package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeErrorLogs() *cli.Command {
	describeErrorLogs := &cli.Command{
		Name: "describe-error-logs",
		Short: i18n.T("", "获取SQL Server 错误日志及下载信息<br>- 仅支持SQL Server"),
		Usage: "describe-error-logs [-h] [--region-id REGIONID] --instance-id INSTANCEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeErrorLogs(ctx)
			return nil
		},
	}

	describeErrorLogs.Flags().Add(regionIdFlag)
	describeErrorLogs.Flags().Add(instanceIdFlag)
	describeErrorLogs.Flags().Add(inputJsonFlag)
	describeErrorLogs.Flags().Add(headersFlag)
	return describeErrorLogs
}

func doDescribeErrorLogs(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()

	req := apis.NewDescribeErrorLogsRequestWithAllParams(regionId, instanceId)

	resp, err := rdsClient.DescribeErrorLogs(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
