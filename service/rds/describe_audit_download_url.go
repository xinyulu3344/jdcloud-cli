package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeAuditDownloadUrl() *cli.Command {
	describeAuditDownloadUrl := &cli.Command{
		Name: "describe-audit-download-url",
		Short: i18n.T("", "获取某个审计文件的下载链接，同时支持内链和外链，链接的有效时间为24小时<br>- 仅支持SQL Server。"),
		Usage: "describe-audit-download-url [-h] [--region-id REGIONID] --instance-id INSTANCEID --file-name FILENAME [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeAuditDownloadUrl(ctx)
			return nil
		},
	}

	describeAuditDownloadUrl.Flags().Add(regionIdFlag)
	describeAuditDownloadUrl.Flags().Add(instanceIdFlag)
	describeAuditDownloadUrl.Flags().Add(fileNameFlag)
	describeAuditDownloadUrl.Flags().Add(inputJsonFlag)
	describeAuditDownloadUrl.Flags().Add(headersFlag)
	return describeAuditDownloadUrl
}

func doDescribeAuditDownloadUrl(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()
	fileName, _ := fileNameFlag.GetValue()


	req := apis.NewDescribeAuditDownloadURLRequestWithAllParams(regionId, instanceId, fileName)

	resp, err := rdsClient.DescribeAuditDownloadURL(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}