package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
	"strconv"
)

func NewDescribeInstances() *cli.Command {
	describeInstances := &cli.Command{
		Name: "describe-instances",
		Short: i18n.T("", "获取当前账号下所有RDS实例及MySQL只读实例的概要信息，例如实例类型，版本，计费信息等。"),
		Usage: "describe-instances [-h] [--region-id REGIONID] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeInstances(ctx)
			return nil
		},
	}

	describeInstances.Flags().Add(regionIdFlag)
	describeInstances.Flags().Add(pageNumberFlag)
	describeInstances.Flags().Add(pageSizeFlag)
	describeInstances.Flags().Add(inputJsonFlag)
	describeInstances.Flags().Add(headersFlag)
	return describeInstances
}

func doDescribeInstances(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()

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

	req := apis.NewDescribeInstancesRequestWithAllParams(regionId, &pageNumber, &pageSize)

	resp, err := rdsClient.DescribeInstances(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
