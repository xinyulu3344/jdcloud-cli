package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeNetworkInterface() *cli.Command {
	describeNetworkInterface := &cli.Command{
		Name: "describe-network-interface",
		Short: i18n.T("", "查询弹性网卡信息详情。"),
		Usage: "describe-network-interface [-h] [--region-id REGIONID] --network-interface-id NETWORKINTERFACEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeNetworkInterface(ctx)
			return nil
		},
	}

	describeNetworkInterface.Flags().Add(regionIdFlag)
	describeNetworkInterface.Flags().Add(networkInterfaceIdFlag)
	describeNetworkInterface.Flags().Add(inputJsonFlag)
	describeNetworkInterface.Flags().Add(headersFlag)
	return describeNetworkInterface
}

func doDescribeNetworkInterface(ctx *cli.Context) {

	profile := GetProfile(ctx)
	vpcClient := GetVpcClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	networkInterfaceId, _ := networkInterfaceIdFlag.GetValue()

	req := apis.NewDescribeNetworkInterfaceRequestWithAllParams(regionId, networkInterfaceId)

	resp, err := vpcClient.DescribeNetworkInterface(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}