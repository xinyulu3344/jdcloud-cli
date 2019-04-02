package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeImportFiles() *cli.Command {
	describeImportFiles := &cli.Command{
		Name: "describe-import-files",
		Short: i18n.T("", "获取用户通过单库上云工具上传到该实例上的文件列表<br>- 仅支持SQL Server。"),
		Usage: "describe-import-files [-h] [--region-id REGIONID] --instance-id INSTANCEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeImportFiles(ctx)
			return nil
		},
	}

	describeImportFiles.Flags().Add(regionIdFlag)
	describeImportFiles.Flags().Add(instanceIdFlag)
	describeImportFiles.Flags().Add(inputJsonFlag)
	describeImportFiles.Flags().Add(headersFlag)
	return describeImportFiles
}

func doDescribeImportFiles(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()


	req := apis.NewDescribeImportFilesRequestWithAllParams(regionId, instanceId)

	resp, err := rdsClient.DescribeImportFiles(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
