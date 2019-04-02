package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeSubnet() *cli.Command {
	describeSubnet := &cli.Command{
		Name: "describe-subnet",
		Short: i18n.T("", "查询子网信息详情。"),
		Usage: "describe-subnet [-h] [--region-id REGIONID] --subnet-id SUBNETID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeSubnet(ctx)
			return nil
		},
	}

	describeSubnet.Flags().Add(regionIdFlag)
	describeSubnet.Flags().Add(subnetIdFlag)
	describeSubnet.Flags().Add(inputJsonFlag)
	describeSubnet.Flags().Add(headersFlag)
	return describeSubnet
}

func doDescribeSubnet(ctx *cli.Context) {

	profile := GetProfile(ctx)
	vpcClient := GetVpcClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	subnetId, _ := subnetIdFlag.GetValue()

	req := apis.NewDescribeSubnetRequestWithAllParams(regionId, subnetId)

	resp, err := vpcClient.DescribeSubnet(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}