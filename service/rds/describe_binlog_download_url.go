package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeBinlogDownloadUrl() *cli.Command {
	describeBinlogDownloadUrl := &cli.Command{
		Name: "describe-binlog-download-url",
		Short: i18n.T("", "获取MySQL实例的binlog的下载链接<br>- 仅支持MySQL"),
		Usage: "describe-binlog-download-url [-h] [--region-id REGIONID] --instance-id INSTANCEID --binlog-backup-id BINLOGBACKUPID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeBinlogDownloadUrl(ctx)
			return nil
		},
	}

	describeBinlogDownloadUrl.Flags().Add(regionIdFlag)
	describeBinlogDownloadUrl.Flags().Add(instanceIdFlag)
	describeBinlogDownloadUrl.Flags().Add(binlogBackupIdFlag)
	describeBinlogDownloadUrl.Flags().Add(inputJsonFlag)
	describeBinlogDownloadUrl.Flags().Add(headersFlag)
	return describeBinlogDownloadUrl
}

func doDescribeBinlogDownloadUrl(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()
	binlogBackupId, _ := binlogBackupIdFlag.GetValue()


	req := apis.NewDescribeBinlogDownloadURLRequestWithAllParams(regionId, instanceId, binlogBackupId)

	resp, err := rdsClient.DescribeBinlogDownloadURL(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
