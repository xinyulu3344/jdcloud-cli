package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"strconv"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeNetworkSecurityGroups() *cli.Command{

	describeNetworkSecurityGroups := &cli.Command{
		Name: "describe-network-security-groups",
		Short: i18n.T("", "查询安全组列表。"),
		Usage: "describe-network-security-groups [-h] [--region-id REGIONID] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--filters FILTERS] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeNetworkSecurityGroups(ctx)
			return nil
		},
	}

	describeNetworkSecurityGroups.Flags().Add(regionIdFlag)
	describeNetworkSecurityGroups.Flags().Add(pageNumberFlag)
	describeNetworkSecurityGroups.Flags().Add(pageSizeFlag)
	describeNetworkSecurityGroups.Flags().Add(filtersFlag)
	describeNetworkSecurityGroups.Flags().Add(inputJsonFlag)
	describeNetworkSecurityGroups.Flags().Add(headersFlag)
	return describeNetworkSecurityGroups
}


func doDescribeNetworkSecurityGroups(ctx *cli.Context) {

	profile := GetProfile(ctx)
	vpcClient := GetVpcClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()

	filtersStr, _ := filtersFlag.GetValue()
	filters := Str2Filters(filtersStr)

	pageNumberStr, _ := pageNumberFlag.GetValue()
	pageNumber, err := strconv.Atoi(pageNumberStr)
	if pageNumberStr != "" && err != nil {
		cli.Errorf(ctx.Writer(), "pageNumberStr strconv.Atoi err: %s", err)
	}

	pageSizeStr, _ := pageSizeFlag.GetValue()
	pageSize, err := strconv.Atoi(pageSizeStr)
	if pageSizeStr != "" && err != nil {
		cli.Errorf(ctx.Writer(), "pageSizeStr strconv.Atoi err: %s", err)
	}

	req := apis.NewDescribeNetworkInterfacesRequestWithAllParams(regionId, &pageNumber, &pageSize, filters)

	resp, err := vpcClient.DescribeNetworkInterfaces(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
