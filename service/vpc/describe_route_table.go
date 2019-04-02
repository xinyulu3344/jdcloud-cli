package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeRouteTable() *cli.Command {
	describeRouteTable := &cli.Command{
		Name: "describe-route-table",
		Short: i18n.T("", "查询路由表信息详情。"),
		Usage: "describe-route-table [-h] [--region-id REGIONID] --route-table-id ROUTETABLEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeRouteTable(ctx)
			return nil
		},
	}

	describeRouteTable.Flags().Add(regionIdFlag)
	describeRouteTable.Flags().Add(inputJsonFlag)
	describeRouteTable.Flags().Add(headersFlag)
	describeRouteTable.Flags().Add(routeTableIdFlag)
	return describeRouteTable
}

func doDescribeRouteTable(ctx *cli.Context) {

	profile := GetProfile(ctx)
	vpcClient := GetVpcClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	routeTableId, _ := routeTableIdFlag.GetValue()

	req := apis.NewDescribeRouteTableRequestWithAllParams(regionId, routeTableId)

	resp, err := vpcClient.DescribeRouteTable(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
