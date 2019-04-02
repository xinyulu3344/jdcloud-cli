package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeBackupPolicy() *cli.Command {
	describeBackupPolicy := &cli.Command{
		Name: "describe-backup-policy",
		Short: i18n.T("", "查看RDS实例备份策略。根据数据库类型的不同，支持的备份策略也略有差异，具体请看返回参数中的详细说明。"),
		Usage: "describe-backup-policy [-h] [--region-id REGIONID] --instance-id INSTANCEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeBackupPolicy(ctx)
			return nil
		},
	}

	describeBackupPolicy.Flags().Add(regionIdFlag)
	describeBackupPolicy.Flags().Add(instanceIdFlag)
	describeBackupPolicy.Flags().Add(inputJsonFlag)
	describeBackupPolicy.Flags().Add(headersFlag)
	return describeBackupPolicy
}

func doDescribeBackupPolicy(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()


	req := apis.NewDescribeBackupPolicyRequestWithAllParams(regionId, instanceId)

	resp, err := rdsClient.DescribeBackupPolicy(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
