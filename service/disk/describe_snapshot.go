package disk

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/disk/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeSnapshot() *cli.Command {
	describeSnapshot := &cli.Command{
		Name: "describe-snapshot",
		Short: i18n.T("", "查询云硬盘快照信息详情。"),
		Usage: "describe-snapshot [-h] [--region-id REGIONID] --snapshot-id SNAPSHOTID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeSnapshot(ctx)
			return nil
		},
	}

	describeSnapshot.Flags().Add(regionIdFlag)
	describeSnapshot.Flags().Add(inputJsonFlag)
	describeSnapshot.Flags().Add(headersFlag)
	describeSnapshot.Flags().Add(snapshotIdFlag)
	return describeSnapshot
}

func doDescribeSnapshot(ctx *cli.Context) {

	profile := GetProfile(ctx)
	diskClient := GetDiskClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()

	snapshotId, _ := snapshotIdFlag.GetValue()


	req := apis.NewDescribeSnapshotRequestWithAllParams(regionId, snapshotId)

	resp, err := diskClient.DescribeSnapshot(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
