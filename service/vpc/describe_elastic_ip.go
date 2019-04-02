package vpc

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeElasticIp() *cli.Command {

	describeElasticIp := &cli.Command{
		Name: "describe-elastic-ip",
		Short: i18n.T("", "ElasticIp资源信息详情。"),
		Usage: "describe-elastic-ip [-h] [--region-id REGIONID] --elastic-ip-id ELASTICIPID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeElasticIp(ctx)
			return nil
		},
	}

	describeElasticIp.Flags().Add(regionIdFlag)
	describeElasticIp.Flags().Add(elasticIpIdFlag)
	describeElasticIp.Flags().Add(inputJsonFlag)
	describeElasticIp.Flags().Add(headersFlag)
	return describeElasticIp
}


func doDescribeElasticIp(ctx *cli.Context) {

	profile := GetProfile(ctx)
	vpcClient := GetVpcClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()
	elasticIPid, _ := elasticIpIdFlag.GetValue()

	req := apis.NewDescribeElasticIpRequestWithAllParams(regionId, elasticIPid)

	resp, err := vpcClient.DescribeElasticIp(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
