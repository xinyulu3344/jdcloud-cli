package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"encoding/json"
	"fmt"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	"strconv"
)

func NewDescribeInstancePrivateIpAddress() *cli.Command{
	describeInstancePrivateIpAddress := &cli.Command{
		Name: "describe-instance-private-ip-address",
		Short: i18n.T("", "批量查询云主机内网IP地址，查询的是主网卡内网主IP地址"),
		Usage: "usage: jdc vm describe-instance-private-ip-address [-h] [--region-id REGIONID] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--filters FILTERS] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeInstancePrivateIpAddress(ctx)
			return nil
		},
	}
	describeInstancePrivateIpAddress.Flags().Add(regionIdFlag)
	describeInstancePrivateIpAddress.Flags().Add(pageNumberFlag)
	describeInstancePrivateIpAddress.Flags().Add(filtersFlag)
	describeInstancePrivateIpAddress.Flags().Add(pageSizeFlag)
	describeInstancePrivateIpAddress.Flags().Add(inputJsonFlag)
	describeInstancePrivateIpAddress.Flags().Add(headersFlag)
	return describeInstancePrivateIpAddress
}

func doDescribeInstancePrivateIpAddress(ctx *cli.Context) {
	profile := GetProfile(ctx)
	vmClient := GetVmClient(ctx)

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

	req := apis.NewDescribeInstancePrivateIpAddressRequestWithAllParams(regionId, &pageNumber, &pageSize, filters)
	resp, err := vmClient.DescribeInstancePrivateIpAddress(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}

