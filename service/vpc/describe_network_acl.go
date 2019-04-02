package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeNetworkAcl() *cli.Command {
	describeNetworkAcl := &cli.Command{
		Name: "describe-network-acl",
		Short: i18n.T("", "查询networkAcl资源详情。"),
		Usage: "describe-network-acl [-h] [--region-id REGIONID] --network-acl-id NETWORKACLID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeNetworkAcl(ctx)
			return nil
		},
	}

	describeNetworkAcl.Flags().Add(regionIdFlag)
	describeNetworkAcl.Flags().Add(inputJsonFlag)
	describeNetworkAcl.Flags().Add(headersFlag)
	describeNetworkAcl.Flags().Add(networkAclIdFlag)
	return describeNetworkAcl
}

func doDescribeNetworkAcl(ctx *cli.Context) {
	profile := GetProfile(ctx)
	vpcClient := GetVpcClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	networkAclId, _ := networkAclIdFlag.GetValue()

	req := apis.NewDescribeNetworkAclRequest(regionId, networkAclId)

	resp, err := vpcClient.DescribeNetworkAcl(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}