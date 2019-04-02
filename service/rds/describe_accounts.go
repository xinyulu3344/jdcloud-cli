package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeAccounts() *cli.Command {
	describeAccounts := &cli.Command{
		Name: "describe-accounts",
		Short: i18n.T("", "查看某个RDS实例下所有账号信息，包括账号名称、对各个数据库的访问权限信息等。"),
		Usage: "describe-accounts [-h] [--region-id REGIONID] --instance-id INSTANCEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeAccounts(ctx)
			return nil
		},
	}

	describeAccounts.Flags().Add(regionIdFlag)
	describeAccounts.Flags().Add(inputJsonFlag)
	describeAccounts.Flags().Add(headersFlag)
	describeAccounts.Flags().Add(instanceIdFlag)
	return describeAccounts
}

func doDescribeAccounts(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()


	req := apis.NewDescribeAccountsRequestWithAllParams(regionId, instanceId)

	resp, err := rdsClient.DescribeAccounts(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
