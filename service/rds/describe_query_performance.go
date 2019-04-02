package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
	"strconv"
)

func NewDescribeQueryPerformance() *cli.Command {
	describeQueryPerformance := &cli.Command{
		Name: "describe-query-performance",
		Short: i18n.T("", "根据用户定义的查询条件，获取SQL执行的性能统计信息，例如慢SQL等。用户可以根据这些信息查找与SQL执行相关的性能瓶颈，并进行优化。<br>- 仅支持SQL Server。"),
		Usage: `describe-query-performance [-h] [--region-id REGIONID] --instance-id INSTANCEID --query-type QUERYTYPE [--threshold THRESHOLD] [--page-number PAGENUMBER] [--page-size PAGESIZE]
                                          [--input-json INPUT_JSON] [--headers HEADERS]`,
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeQueryPerformance(ctx)
			return nil
		},
	}

	describeQueryPerformance.Flags().Add(regionIdFlag)
	describeQueryPerformance.Flags().Add(instanceIdFlag)
	describeQueryPerformance.Flags().Add(queryTypeFlag)
	describeQueryPerformance.Flags().Add(thresholdFlag)
	describeQueryPerformance.Flags().Add(pageNumberFlag)
	describeQueryPerformance.Flags().Add(pageSizeFlag)
	describeQueryPerformance.Flags().Add(inputJsonFlag)
	describeQueryPerformance.Flags().Add(headersFlag)
	return describeQueryPerformance
}

func doDescribeQueryPerformance(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()
	queryType, _ := queryTypeFlag.GetValue()

	thresholdStr, _ := thresholdFlag.GetValue()
	threshold, err := strconv.Atoi(thresholdStr)
	if thresholdStr != "" && err != nil {
		cli.Errorf(ctx.Writer(), "thresholdStr strconv.Atoi err: %s", err)
	}

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


	req := apis.NewDescribeQueryPerformanceRequestWithAllParams(regionId, instanceId, queryType, &threshold, &pageNumber, &pageSize)

	resp, err := rdsClient.DescribeQueryPerformance(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
