package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"encoding/json"
	"fmt"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
)

func NewDescribeVpc() *cli.Command {

	describeVpc := &cli.Command{
		Name: "describe-vpc",
		Short: i18n.T("", "查询Vpc信息详情。"),
		Usage: "describe-vpc [-h] [--region-id REGIONID] --vpc-id VPCID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeVpc(ctx)
			return nil
		},
	}

	describeVpc.Flags().Add(regionIdFlag)
	describeVpc.Flags().Add(vpcIdFlag)
	describeVpc.Flags().Add(inputJsonFlag)
	describeVpc.Flags().Add(headersFlag)
	return describeVpc
}


func doDescribeVpc(ctx *cli.Context) {

	profile := GetProfile(ctx)
	vpcClient := GetVpcClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	vpcId, _ := vpcIdFlag.GetValue()

	req := apis.NewDescribeVpcRequestWithAllParams(regionId, vpcId)

	resp, err := vpcClient.DescribeVpc(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
