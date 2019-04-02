package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeInstanceAttributes() *cli.Command {
	describeInstanceAttributes := &cli.Command{
		Name: "describe-instance-attributes",
		Short: i18n.T("", "查询RDS实例（MySQL、SQL Server等）的详细信息以及MySQL只读实例详细信息。"),
		Usage: "describe-instance-attributes [-h] [--region-id REGIONID] --instance-id INSTANCEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeInstanceAttributes(ctx)
			return nil
		},
	}

	describeInstanceAttributes.Flags().Add(regionIdFlag)
	describeInstanceAttributes.Flags().Add(instanceIdFlag)
	describeInstanceAttributes.Flags().Add(inputJsonFlag)
	describeInstanceAttributes.Flags().Add(headersFlag)
	return describeInstanceAttributes
}

func doDescribeInstanceAttributes(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()


	req := apis.NewDescribeInstanceAttributesRequestWithAllParams(regionId, instanceId)

	resp, err := rdsClient.DescribeInstanceAttributes(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
