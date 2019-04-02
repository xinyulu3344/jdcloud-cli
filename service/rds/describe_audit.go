package rds

import (
"jdcloud-cli/cli"
"jdcloud-cli/i18n"
"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
"encoding/json"
"fmt"
)

func NewDescribeAudit() *cli.Command {
	describeAudit := &cli.Command{
		Name: "describe-audit",
		Short: i18n.T("", "查看当前实例已开启的审计选项。如当前实例未开启审计，则返回空<br>- 仅支持SQL Server"),
		Usage: "describe-audit [-h] [--region-id REGIONID] --instance-id INSTANCEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeAudit(ctx)
			return nil
		},
	}

	describeAudit.Flags().Add(regionIdFlag)
	describeAudit.Flags().Add(instanceIdFlag)
	describeAudit.Flags().Add(inputJsonFlag)
	describeAudit.Flags().Add(headersFlag)
	return describeAudit
}

func doDescribeAudit(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()


	req := apis.NewDescribeAuditRequestWithAllParams(regionId, instanceId)

	resp, err := rdsClient.DescribeAudit(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
