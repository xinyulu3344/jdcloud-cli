package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
	"strconv"
)

func NewDescribeIndexPerformance() *cli.Command {
	describeIndexPerformance := &cli.Command{
		Name: "describe-index-performance",
		Short: i18n.T("", "根据用户定义的查询条件，获取索引性能的统计信息，并提供缺失索引及索引创建建议。用户可以根据这些信息查找与索引相关的性能瓶颈，并进行优化。<br>- 仅支持SQL Server。"),
		Usage: "describe-index-performance [-h] [--region-id REGIONID] --instance-id INSTANCEID --query-type QUERYTYPE [--db DB] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--input-json INPUT_JSON]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeIndexPerformance(ctx)
			return nil
		},
	}

	describeIndexPerformance.Flags().Add(regionIdFlag)
	describeIndexPerformance.Flags().Add(instanceIdFlag)
	describeIndexPerformance.Flags().Add(queryIndexTypeFlag)
	describeIndexPerformance.Flags().Add(dbFlag)
	describeIndexPerformance.Flags().Add(pageNumberFlag)
	describeIndexPerformance.Flags().Add(pageSizeFlag)
	describeIndexPerformance.Flags().Add(inputJsonFlag)
	describeIndexPerformance.Flags().Add(headersFlag)
	return describeIndexPerformance
}

func doDescribeIndexPerformance(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()
	queryIndexType, _ := queryIndexTypeFlag.GetValue()
	db, _ := dbFlag.GetValue()

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


	req := apis.NewDescribeIndexPerformanceRequestWithAllParams(regionId, instanceId, queryIndexType, &db, &pageNumber, &pageSize)

	resp, err := rdsClient.DescribeIndexPerformance(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
