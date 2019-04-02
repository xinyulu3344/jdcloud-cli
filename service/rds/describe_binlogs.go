package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
	"strconv"
)

func NewDescribeBinlogs() *cli.Command {
	describeBinlogs := &cli.Command{
		Name: "describe-binlogs",
		Short: i18n.T("", "获取MySQL实例中binlog的详细信息<br>- 仅支持MySQL。"),
		Usage: "describe-binlogs [-h] [--region-id REGIONID] --instance-id INSTANCEID [--page-number PAGENUMBER] [--page-size PAGESIZE] [--start-time STARTTIME] [--end-time ENDTIME] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeBinlogs(ctx)
			return nil
		},
	}

	describeBinlogs.Flags().Add(regionIdFlag)
	describeBinlogs.Flags().Add(instanceIdFlag)
	describeBinlogs.Flags().Add(pageNumberFlag)
	describeBinlogs.Flags().Add(pageSizeFlag)
	describeBinlogs.Flags().Add(inputJsonFlag)
	describeBinlogs.Flags().Add(headersFlag)
	describeBinlogs.Flags().Add(startTimeFlag)
	describeBinlogs.Flags().Add(endTimeFlag)
	return describeBinlogs
}

func doDescribeBinlogs(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()

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

	startTime, _ := startTimeFlag.GetValue()
	endTime, _ := endTimeFlag.GetValue()

	req := apis.NewDescribeBinlogsRequestWithAllParams(regionId, instanceId, &pageNumber, &pageSize, &startTime, &endTime)

	resp, err := rdsClient.DescribeBinlogs(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
