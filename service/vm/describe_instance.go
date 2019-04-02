package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	"encoding/json"
	"fmt"
)


func NewDescribeInstance() *cli.Command {
	describeInstance := &cli.Command{
		Name: "describe-instance",
		Short: i18n.T("", "查询一台云主机的详细信息"),
		Usage: "usage: jdc vm describe-instance [-h] [--region-id REGIONID] --instance-id INSTANCEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeInstance(ctx)
			return nil
		},
	}
	describeInstance.Flags().Add(regionIdFlag)
	describeInstance.Flags().Add(instanceIdFlag)
	describeInstance.Flags().Add(inputJsonFlag)
	describeInstance.Flags().Add(headersFlag)
	return describeInstance
}


func doDescribeInstance(ctx *cli.Context) {

	profile := GetProfile(ctx)
	vmClient := GetVmClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId

	regionId, _ := regionIdFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()

	req := apis.NewDescribeInstanceRequestWithAllParams(regionId, instanceId)
	resp, err := vmClient.DescribeInstance(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}