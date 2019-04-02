package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"strconv"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeInstanceStatus() *cli.Command {
	describeInstanceStatus := &cli.Command{
		Name:  "describe-instance-status",
		Short: i18n.T("", "批量查询云主机状态"),
		Usage: "describe-instance-status [-h] [--region-id REGIONID] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--filters FILTERS] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeInstanceStatus(ctx)
			return nil
		},
	}
	describeInstanceStatus.Flags().Add(regionIdFlag)
	describeInstanceStatus.Flags().Add(pageNumberFlag)
	describeInstanceStatus.Flags().Add(filtersFlag)
	describeInstanceStatus.Flags().Add(pageSizeFlag)
	describeInstanceStatus.Flags().Add(inputJsonFlag)
	describeInstanceStatus.Flags().Add(headersFlag)
	return describeInstanceStatus
}


func doDescribeInstanceStatus(ctx *cli.Context) {
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


	req := apis.NewDescribeInstanceStatusRequestWithAllParams(regionId, &pageNumber, &pageSize, filters)
	resp, err := vmClient.DescribeInstanceStatus(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}