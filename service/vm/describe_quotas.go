package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeQuotas() *cli.Command {
	describeQuotas := &cli.Command{
		Name: "describe-quotas",
		Short: i18n.T("", "查询配额，支持：云主机、镜像、密钥、模板、镜像共享"),
		Usage: "describe-quotas [-h] [--region-id REGIONID] [--filters FILTERS] [--image-id IMAGEID] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeQuotas(ctx)
			return nil
		},
	}
	describeQuotas.Flags().Add(regionIdFlag)
	describeQuotas.Flags().Add(imageIdFlag)
	describeQuotas.Flags().Add(inputJsonFlag)
	describeQuotas.Flags().Add(headersFlag)
	describeQuotas.Flags().Add(filtersFlag)
	return describeQuotas
}

func doDescribeQuotas(ctx *cli.Context) {


	profile := GetProfile(ctx)
	vmClient := GetVmClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId

	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	filtersStr, _ := filtersFlag.GetValue()
	filters := Str2Filters(filtersStr)
	imageId, _ := imageIdFlag.GetValue()

	req := apis.NewDescribeQuotasRequestWithAllParams(regionId, filters, &imageId)

	resp, err := vmClient.DescribeQuotas(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
