package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	"encoding/json"
	"fmt"
)


func NewDescribeImage() *cli.Command {
	describeImage := &cli.Command{
		Name: "describe-image",
		Short: i18n.T("", "查询镜像详情"),
		Usage: "describe-image [-h] [--region-id REGIONID] --image-id IMAGEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {
			doDescribeImage(ctx)
			return nil
		},
	}
	describeImage.Flags().Add(imageIdFlag)
	describeImage.Flags().Add(headersFlag)
	describeImage.Flags().Add(regionIdFlag)
	describeImage.Flags().Add(inputJsonFlag)
	return describeImage
}


func doDescribeImage(ctx *cli.Context) {
	profile := GetProfile(ctx)
	vmClient := GetVmClient(ctx)
	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId
	regionId, _ := regionIdFlag.GetValue()
	imageId, _ := imageIdFlag.GetValue()
	req := apis.NewDescribeImageRequestWithAllParams(regionId, imageId)

	resp, err := vmClient.DescribeImage(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}