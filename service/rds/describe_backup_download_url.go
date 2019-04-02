package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeBackupDownloadUrl() *cli.Command {
	describeBackupDownloadUrl := &cli.Command{
		Name: "describe-backup-download-url",
		Short: i18n.T("", "获取整个备份或备份中单个文件的下载链接。<br>- 当输入参数中有文件名时，获取该文件的下载链接。<br>- 输入参数中无文件名时，获取整个备份的下载链接。<br>由于备份机制的差异，使用该接口下载备份时，SQL Server必须输入文件名，每个文件逐一下载，不支持下载整个备份。SQL Server备份中的文件名（不包括后缀）即为备份的数据库名。例如文件名为my_test_db.bak，表示该文件是my_test_db数据库的备份。<br>MySQL可下载整个备份集，但不支持单个文件的下载。。"),
		Usage: "describe-backup-download-url [-h] [--region-id REGIONID] --backup-id BACKUPID [--file-name FILENAME] [--url-expiration-second URLEXPIRATIONSECOND] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeBackupDownloadUrl(ctx)
			return nil
		},
	}

	describeBackupDownloadUrl.Flags().Add(regionIdFlag)
	describeBackupDownloadUrl.Flags().Add(backupIdFlag)
	describeBackupDownloadUrl.Flags().Add(fileNameFlag)
	describeBackupDownloadUrl.Flags().Add(urlExpirationSecondFlag)
	describeBackupDownloadUrl.Flags().Add(inputJsonFlag)
	describeBackupDownloadUrl.Flags().Add(headersFlag)
	return describeBackupDownloadUrl
}

func doDescribeBackupDownloadUrl(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	backupId, _ := backupIdFlag.GetValue()
	fileName, _ := fileNameFlag.GetValue()
	urlExpirationSecond, _ := urlExpirationSecondFlag.GetValue()


	req := apis.NewDescribeBackupDownloadURLRequestWithAllParams(regionId, backupId, &fileName, &urlExpirationSecond)

	resp, err := rdsClient.DescribeBackupDownloadURL(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
