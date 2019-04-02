package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	"encoding/json"
	"fmt"
)

func NewDescribeImageConstraints() *cli.Command {
	describeImageConstraints := &cli.Command{
		Name: "describe-image-constraints",
		Short: i18n.T("", "查询镜像的实例规格限制; 通过此接口可以查看镜像不支持的实例规格。只有官方镜像、第三方镜像有实例规格的限制，个人的私有镜像没有此限制。"),
		Usage: "describe-image-constraints [-h] [--region-id REGIONID] --image-id IMAGEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {
			doDescribeImageConstraints(ctx)
			return nil
		},
	}
	describeImageConstraints.Flags().Add(imageIdFlag)
	describeImageConstraints.Flags().Add(headersFlag)
	describeImageConstraints.Flags().Add(regionIdFlag)
	describeImageConstraints.Flags().Add(inputJsonFlag)
	return describeImageConstraints
}

func doDescribeImageConstraints(ctx *cli.Context) {
	profile := GetProfile(ctx)
	vmClient := GetVmClient(ctx)
	regionIdFlag.DefaultValue = profile.RegionId
	regionIdFlag.Required = true
	regionId, _ := regionIdFlag.GetValue()
	imageId, _ := imageIdFlag.GetValue()
	req := apis.NewDescribeImageConstraintsRequestWithAllParams(regionId, imageId)

	resp, err := vmClient.DescribeImageConstraints(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}