package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeAuditOptions() *cli.Command {
	describeAuditOptions := &cli.Command{
		Name: "describe-audit-options",
		Short: i18n.T("", "获取当前系统所支持的各种数据库版本的审计选项及相应的推荐选项<br>- 仅支持SQL Server。"),
		Usage: "describe-audit-options [-h] [--region-id REGIONID] --instance-id INSTANCEID --name NAME [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeAuditOptions(ctx)
			return nil
		},
	}

	describeAuditOptions.Flags().Add(regionIdFlag)
	describeAuditOptions.Flags().Add(instanceIdFlag)
	describeAuditOptions.Flags().Add(inputJsonFlag)
	describeAuditOptions.Flags().Add(headersFlag)
	describeAuditOptions.Flags().Add(nameFlag)
	return describeAuditOptions
}

func doDescribeAuditOptions(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()
	name, _ := nameFlag.GetValue()

	req := apis.NewDescribeAuditOptionsRequestWithAllParams(regionId, instanceId, name)

	resp, err := rdsClient.DescribeAuditOptions(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
