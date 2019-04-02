package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeWhiteList() *cli.Command {
	describeWhiteList := &cli.Command{
		Name: "describe-white-list",
		Short: i18n.T("", "查看RDS实例当前白名单。白名单是允许访问当前实例的IP/IP段列表，缺省情况下，白名单对本VPC开放。如果用户开启了外网访问的功能，还需要对外网的IP配置白名单。。"),
		Usage: "describe-white-list [-h] [--region-id REGIONID] --instance-id INSTANCEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeWhiteList(ctx)
			return nil
		},
	}

	describeWhiteList.Flags().Add(regionIdFlag)
	describeWhiteList.Flags().Add(instanceIdFlag)
	describeWhiteList.Flags().Add(inputJsonFlag)
	describeWhiteList.Flags().Add(headersFlag)
	return describeWhiteList
}

func doDescribeWhiteList(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()


	req := apis.NewDescribeWhiteListRequestWithAllParams(regionId, instanceId)

	resp, err := rdsClient.DescribeWhiteList(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
