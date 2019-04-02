package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeVpcPeering() *cli.Command {
	describeVpcPeering := &cli.Command{
		Name: "describe-vpc-peering",
		Short: i18n.T("", "查询VpcPeering资源详情。"),
		Usage: "describe-vpc-peering [-h] [--region-id REGIONID] --vpc-peering-id VPCPEERINGID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeVpcPeering(ctx)
			return nil
		},
	}

	describeVpcPeering.Flags().Add(regionIdFlag)
	describeVpcPeering.Flags().Add(inputJsonFlag)
	describeVpcPeering.Flags().Add(headersFlag)
	describeVpcPeering.Flags().Add(vpcPeeringIdFlag)
	return describeVpcPeering
}

func doDescribeVpcPeering(ctx *cli.Context) {

	profile := GetProfile(ctx)
	vpcClient := GetVpcClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	vpcPeeringId, _ := vpcPeeringIdFlag.GetValue()

	req := apis.NewDescribeVpcPeeringRequestWithAllParams(regionId, vpcPeeringId)

	resp, err := vpcClient.DescribeVpcPeering(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
