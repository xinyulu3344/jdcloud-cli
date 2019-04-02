package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"strconv"
	"encoding/json"
	"fmt"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
)

func NewDescribeInstances() *cli.Command {

	describeInstances := &cli.Command{
		Name: "describe-instances",
		Short: i18n.T("", "批量查询云主机的详细信息; 此接口支持分页查询，默认每页20条。"),
		Usage: "vm describe-instances [-h] [--region-id REGIONID] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--filters FILTERS] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeInstances(ctx)
			return nil
		},
	}
	describeInstances.Flags().Add(regionIdFlag)
	describeInstances.Flags().Add(pageNumberFlag)
	describeInstances.Flags().Add(filtersFlag)
	describeInstances.Flags().Add(pageSizeFlag)
	describeInstances.Flags().Add(inputJsonFlag)
	describeInstances.Flags().Add(headersFlag)
	return describeInstances
}


func doDescribeInstances(ctx *cli.Context) {
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

	req := apis.NewDescribeInstancesRequestWithAllParams(regionId, &pageNumber, &pageSize, filters)
	resp, err := vmClient.DescribeInstances(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))

}
