package disk

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"strconv"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/disk/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeSnapshots() *cli.Command {
	describeSnapshots := &cli.Command{
		Name: "describe-snapshots",
		Short: i18n.T("", ""),
		Usage: "describe-snapshots [-h] [--region-id REGIONID] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--filters FILTERS] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeSnapshots(ctx)
			return nil
		},
	}

	describeSnapshots.Flags().Add(regionIdFlag)
	describeSnapshots.Flags().Add(pageNumberFlag)
	describeSnapshots.Flags().Add(pageSizeFlag)
	describeSnapshots.Flags().Add(filtersFlag)
	describeSnapshots.Flags().Add(inputJsonFlag)
	describeSnapshots.Flags().Add(headersFlag)
	return describeSnapshots
}


func doDescribeSnapshots(ctx *cli.Context) {

	profile := GetProfile(ctx)
	diskClient := GetDiskClient(ctx)

	regionIdFlag.DefaultValue = profile.RegionId
	regionIdFlag.Required = true
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

	req := apis.NewDescribeSnapshotsRequestWithAllParams(regionId, &pageNumber, &pageSize, filters)

	resp, err := diskClient.DescribeSnapshots(req)

	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
