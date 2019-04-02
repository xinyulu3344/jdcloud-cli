package disk

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"encoding/json"
	"fmt"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/disk/apis"
)

func NewDescribeDisk() *cli.Command {
	describeDisk := &cli.Command{
		Name: "describe-disk",
		Short: i18n.T("", "查询云硬盘信息详情。"),
		Usage: "describe-disk [-h] [--region-id REGIONID] --disk-id DISKID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeDisk(ctx)
			return nil
		},
	}
	describeDisk.Flags().Add(regionIdFlag)
	describeDisk.Flags().Add(diskIdFlag)
	describeDisk.Flags().Add(inputJsonFlag)
	describeDisk.Flags().Add(headersFlag)
	return describeDisk
}

func doDescribeDisk(ctx *cli.Context) {
	profile := GetProfile(ctx)
	diskClient := GetDiskClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	diskId, _ := diskIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()


	req := apis.NewDescribeDiskRequestWithAllParams(regionId, diskId)

	resp, err := diskClient.DescribeDisk(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
