package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"strconv"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeNetworkAcls() *cli.Command {
	describeNetworkAcls := &cli.Command{
		Name: "describe-network-acls",
		Short: i18n.T("", "查询Acl列表。"),
		Usage: "describe-network-acls [-h] [--region-id REGIONID] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--filters FILTERS] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeNetworkAcls(ctx)
			return nil
		},
	}

	describeNetworkAcls.Flags().Add(regionIdFlag)
	describeNetworkAcls.Flags().Add(inputJsonFlag)
	describeNetworkAcls.Flags().Add(headersFlag)
	describeNetworkAcls.Flags().Add(filtersFlag)
	describeNetworkAcls.Flags().Add(pageNumberFlag)
	describeNetworkAcls.Flags().Add(pageSizeFlag)
	return describeNetworkAcls
}


func doDescribeNetworkAcls(ctx *cli.Context) {

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

	req := apis.NewDescribeNetworkAclsRequestWithAllParams(regionId, &pageNumber, &pageSize, filters)

	resp, err := vpcClient.DescribeNetworkAcls(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}