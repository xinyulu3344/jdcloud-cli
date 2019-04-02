package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeAuditFiles() *cli.Command {
	describeAuditFiles := &cli.Command{
		Name: "describe-audit-files",
		Short: i18n.T("", "获取当前实例下的所有审计结果文件的列表<br>- 仅支持SQL Server。"),
		Usage: "describe-audit-files [-h] [--region-id REGIONID] --instance-id INSTANCEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeAuditFiles(ctx)
			return nil
		},
	}

	describeAuditFiles.Flags().Add(regionIdFlag)
	describeAuditFiles.Flags().Add(instanceIdFlag)
	describeAuditFiles.Flags().Add(inputJsonFlag)
	describeAuditFiles.Flags().Add(headersFlag)
	return describeAuditFiles
}

func doDescribeAuditFiles(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()


	req := apis.NewDescribeAuditFilesRequestWithAllParams(regionId, instanceId)

	resp, err := rdsClient.DescribeAuditFiles(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
