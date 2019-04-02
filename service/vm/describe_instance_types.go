package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"encoding/json"
	"fmt"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
)

func NewDescribeInstanceType() *cli.Command {
	describeInstanceType := &cli.Command{
		Name: "describe-instance-type",
		Short: i18n.T("", "查询实例规格信息列表"),
		Usage: "describe-instance-types [-h] [--region-id REGIONID] [--filters FILTERS] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeInstanceType(ctx)
			return nil
		},
	}
	describeInstanceType.Flags().Add(regionIdFlag)
	describeInstanceType.Flags().Add(filtersFlag)
	describeInstanceType.Flags().Add(inputJsonFlag)
	describeInstanceType.Flags().Add(headersFlag)
	return describeInstanceType
}

func doDescribeInstanceType(ctx *cli.Context){
	profile := GetProfile(ctx)
	vmClient := GetVmClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()

	//inputJson := inputJsonFlag.GetValue()
	//headers := headersFlag.GetValue()

	filtersStr, _ := filtersFlag.GetValue()
	filters := Str2Filters(filtersStr)

	req := apis.NewDescribeInstanceTypesRequestWithAllParams(regionId, filters)
	resp, err := vmClient.DescribeInstanceTypes(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}