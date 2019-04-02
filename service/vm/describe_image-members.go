package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	"encoding/json"
	"fmt"
)


func NewDescribeImageMembers() *cli.Command{
	describeImageMembers := &cli.Command{
		Name: "describe-image-members",
		Short: i18n.T("", "查询镜像共享帐户列表，只允许操作您的个人私有镜像。"),
		Usage: "describe-image-members [-h] [--region-id REGIONID] --image-id IMAGEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {
			doDescribeImageMembers(ctx)
			return nil
		},
	}

	describeImageMembers.Flags().Add(imageIdFlag)
	describeImageMembers.Flags().Add(headersFlag)
	describeImageMembers.Flags().Add(regionIdFlag)
	describeImageMembers.Flags().Add(inputJsonFlag)
	return describeImageMembers
}

func doDescribeImageMembers(ctx *cli.Context) {
	profile := GetProfile(ctx)
	vmClient := GetVmClient(ctx)
	regionIdFlag.DefaultValue = profile.RegionId
	regionIdFlag.Required = true

	regionId, _ := regionIdFlag.GetValue()
	imageId, _ := imageIdFlag.GetValue()

	req := apis.NewDescribeImageMembersRequestWithAllParams(regionId, imageId)

	resp, err := vmClient.DescribeImageMembers(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
