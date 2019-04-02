package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"encoding/json"
	"fmt"
	"strconv"
)

func NewDescribeSubnets() *cli.Command {
	describeSubnets := &cli.Command{
		Name: "describe-subnets",
		Short: i18n.T("", "查询子网列表。"),
		Usage: "describe-subnets [-h] [--region-id REGIONID] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--filters FILTERS] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeSubnets(ctx)
			return nil
		},
	}

	describeSubnets.Flags().Add(regionIdFlag)
	describeSubnets.Flags().Add(subnetIdFlag)
	describeSubnets.Flags().Add(inputJsonFlag)
	describeSubnets.Flags().Add(headersFlag)
	describeSubnets.Flags().Add(pageNumberFlag)
	describeSubnets.Flags().Add(pageSizeFlag)
	return describeSubnets
}

func doDescribeSubnets(ctx *cli.Context) {


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

	req := apis.NewDescribeSubnetsRequestWithAllParams(regionId, &pageNumber, &pageSize, filters)

	resp, err := vpcClient.DescribeSubnets(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}