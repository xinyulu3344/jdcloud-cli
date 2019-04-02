package disk

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"strconv"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/disk/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeDisks() *cli.Command {
	describeDisks := &cli.Command{
		Name: "describe-disks",
		Short: i18n.T("", "查询云硬盘列表。"),
		Usage: "describe-disks [-h] [--region-id REGIONID] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--tags TAGS] [--filters FILTERS] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeDisks(ctx)
			return nil
		},
	}
	describeDisks.Flags().Add(regionIdFlag)
	describeDisks.Flags().Add(pageNumberFlag)
	describeDisks.Flags().Add(pageSizeFlag)
	describeDisks.Flags().Add(filtersFlag)
	describeDisks.Flags().Add(inputJsonFlag)
	describeDisks.Flags().Add(headersFlag)
	describeDisks.Flags().Add(tagsFlag)
	return describeDisks
}


func doDescribeDisks(ctx *cli.Context) {

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

	tagsStr, _ := tagsFlag.GetValue()
	tags := Str2Tags(tagsStr)


	req := apis.NewDescribeDisksRequestWithAllParams(regionId, &pageNumber, &pageSize, tags, filters)

	resp, err := diskClient.DescribeDisks(req)

	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))

}
