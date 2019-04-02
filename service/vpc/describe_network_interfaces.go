package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"strconv"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeNetworkInterfaces() *cli.Command {
	describeNetworkInterfaces := &cli.Command{
		Name: "describe-network-interfaces",
		Short: i18n.T("" ,"查询弹性网卡列表。"),
		Usage: "describe-network-interfaces [-h] [--region-id REGIONID] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--filters FILTERS] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeNetworkInterfaces(ctx)
			return nil
		},
	}

	describeNetworkInterfaces.Flags().Add(regionIdFlag)
	describeNetworkInterfaces.Flags().Add(pageNumberFlag)
	describeNetworkInterfaces.Flags().Add(pageSizeFlag)
	describeNetworkInterfaces.Flags().Add(filtersFlag)
	describeNetworkInterfaces.Flags().Add(inputJsonFlag)
	describeNetworkInterfaces.Flags().Add(headersFlag)
	return describeNetworkInterfaces
}

func doDescribeNetworkInterfaces(ctx *cli.Context) {

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