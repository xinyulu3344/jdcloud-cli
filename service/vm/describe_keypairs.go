package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"strconv"
	"encoding/json"
	"fmt"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
)

func NewDescribeKeypairs() *cli.Command {

	describeKeypairs := &cli.Command{
		Name: "describe-keypairs",
		Short: i18n.T("", "批量查询密钥对; 此接口支持分页查询，默认每页20条。"),
		Usage: "describe-keypairs [-h] [--region-id REGIONID] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--filters FILTERS] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeKeypairs(ctx)
			return nil
		},
	}
	describeKeypairs.Flags().Add(regionIdFlag)
	describeKeypairs.Flags().Add(pageNumberFlag)
	describeKeypairs.Flags().Add(filtersFlag)
	describeKeypairs.Flags().Add(pageSizeFlag)
	describeKeypairs.Flags().Add(inputJsonFlag)
	describeKeypairs.Flags().Add(headersFlag)
	return describeKeypairs
}


func doDescribeKeypairs(ctx *cli.Context) {

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

	req := apis.NewDescribeKeypairsRequestWithAllParams(regionId, &pageNumber, &pageSize, filters)
	resp, err := vmClient.DescribeKeypairs(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}