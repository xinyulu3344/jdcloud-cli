package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeQuota() *cli.Command {
	describeQuota := &cli.Command{
		Name: "describe-quota",
		Short: i18n.T("", "查询配额信息。"),
		Usage: "describe-quota [-h] [--region-id REGIONID] --type TYPE [--parent-resource-id PARENTRESOURCEID] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeQuota(ctx)
			return nil
		},
	}

	describeQuota.Flags().Add(regionIdFlag)
	describeQuota.Flags().Add(typeFlag)
	describeQuota.Flags().Add(parentResourceIdFlag)
	describeQuota.Flags().Add(inputJsonFlag)
	describeQuota.Flags().Add(headersFlag)
	return describeQuota
}

func doDescribeQuota(ctx *cli.Context) {

	profile := GetProfile(ctx)
	vpcClient := GetVpcClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	tp, _ :=  typeFlag.GetValue()
	parentResourceId, _ := parentResourceIdFlag.GetValue()

	req := apis.NewDescribeQuotaRequestWithAllParams(regionId, tp, &parentResourceId)

	resp, err := vpcClient.DescribeQuota(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
