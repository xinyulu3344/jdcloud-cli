package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeNetworkSecurityGroup() *cli.Command {
	describeNetworkSecurityGroup := &cli.Command{
		Name: "describe-network-security-group",
		Short: i18n.T("", "查询安全组信息详情。"),
		Usage: "describe-network-security-group [-h] [--region-id REGIONID] --network-security-group-id NETWORKSECURITYGROUPID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeNetworkSecurityGroup(ctx)
			return nil
		},
	}

	describeNetworkSecurityGroup.Flags().Add(regionIdFlag)
	describeNetworkSecurityGroup.Flags().Add(networkSecurityGroupIdFlag)
	describeNetworkSecurityGroup.Flags().Add(inputJsonFlag)
	describeNetworkSecurityGroup.Flags().Add(headersFlag)
	return describeNetworkSecurityGroup
}


func doDescribeNetworkSecurityGroup(ctx *cli.Context) {

	profile := GetProfile(ctx)
	vpcClient := GetVpcClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	networkSecurityGroupId, _ := networkSecurityGroupIdFlag.GetValue()

	req := apis.NewDescribeNetworkSecurityGroupRequestWithAllParams(regionId, networkSecurityGroupId)

	resp, err := vpcClient.DescribeNetworkSecurityGroup(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
