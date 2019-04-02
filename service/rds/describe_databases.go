package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeDatabases() *cli.Command {
	describeDatabases := &cli.Command{
		Name: "describe-databases",
		Short: i18n.T("", "获取当前实例的所有数据库详细信息的列表。"),
		Usage: "describe-databases [-h] [--region-id REGIONID] --instance-id INSTANCEID [--db-name DBNAME] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeDatabases(ctx)
			return nil
		},
	}

	describeDatabases.Flags().Add(regionIdFlag)
	describeDatabases.Flags().Add(instanceIdFlag)
	describeDatabases.Flags().Add(dbNameFlag)
	describeDatabases.Flags().Add(inputJsonFlag)
	describeDatabases.Flags().Add(headersFlag)
	return describeDatabases
}

func doDescribeDatabases(ctx *cli.Context) {

	profile := GetProfile(ctx)
	rdsClient := GetRdsClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()
	dbName, _ := dbNameFlag.GetValue()

	req := apis.NewDescribeDatabasesRequestWithAllParams(regionId, instanceId, &dbName)

	resp, err := rdsClient.DescribeDatabases(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
