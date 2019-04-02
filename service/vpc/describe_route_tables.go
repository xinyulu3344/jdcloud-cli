package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"strconv"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeRouteTables() *cli.Command {

	describeRouteTables := &cli.Command{
		Name: "describe-route-tables",
		Short: i18n.T("", "查询路由表列表。"),
		Usage: "describe-route-tables [-h] [--region-id REGIONID] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--filters FILTERS] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeRouteTables(ctx)
			return nil
		},
	}

	describeRouteTables.Flags().Add(regionIdFlag)
	describeRouteTables.Flags().Add(pageNumberFlag)
	describeRouteTables.Flags().Add(pageSizeFlag)
	describeRouteTables.Flags().Add(filtersFlag)
	describeRouteTables.Flags().Add(inputJsonFlag)
	describeRouteTables.Flags().Add(headersFlag)
	return describeRouteTables
}

func doDescribeRouteTables(ctx *cli.Context) {

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

	req := apis.NewDescribeRouteTablesRequestWithAllParams(regionId, &pageNumber, &pageSize, filters)

	resp, err := vpcClient.DescribeRouteTables(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
