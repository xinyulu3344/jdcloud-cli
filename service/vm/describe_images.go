package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	"encoding/json"
	"fmt"
	"strconv"
)

func NewDescribeImages() *cli.Command {
	describeImages := &cli.Command{
		Name: "describe-images",
		Short: i18n.T("", "查询镜像信息列表; 通过此接口可以查询到京东云官方镜像、第三方镜像、私有镜像、或其他用户共享给您的镜像; 此接口支持分页查询，默认每页20条。"),
		Usage: "describe-images [-h] [--region-id REGIONID] [--image-source IMAGESOURCE] [--platform PLATFORM] [--ids IDS] [--root-device-type ROOTDEVICETYPE] [--status STATUS] [--page-number PAGENUMBER] [--page-size PAGESIZE] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeImages(ctx)
			return nil
		},
	}
	describeImages.Flags().Add(regionIdFlag)
	describeImages.Flags().Add(imageSourceFlag)
	describeImages.Flags().Add(platFormFlag)
	describeImages.Flags().Add(idsFlag)
	describeImages.Flags().Add(rootDeviceTypeFlag)
	describeImages.Flags().Add(statusFlag)
	describeImages.Flags().Add(pageNumberFlag)
	describeImages.Flags().Add(pageSizeFlag)
	return describeImages
}

func doDescribeImages(ctx *cli.Context) {
	profile := GetProfile(ctx)
	vmClient := GetVmClient(ctx)
	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId

	// 获取参数值
	regionId, _ := regionIdFlag.GetValue()
	imageSource, _ := imageIdFlag.GetValue()
	platForm, _ := platFormFlag.GetValue()
	ids, _ := idsFlag.GetValue()
	idsArr := Str2Arr(ids)
	rootDeviceType, _ := rootDeviceTypeFlag.GetValue()
	status, _ := statusFlag.GetValue()
	pageNumberStr, _ := pageNumberFlag.GetValue()
	pageNumber, err := strconv.Atoi(pageNumberStr)
	if pageNumberStr != "" && err != nil {
		cli.Errorf(ctx.Writer(), "pageNumberStr strconv.Atoi err: %s", err)
	}
	pageSizeStr, _ := pageSizeFlag.GetValue()
	pageSize, err := strconv.Atoi(pageSizeStr)
	if pageSizeStr != "" && err != nil {
		cli.Errorf(ctx.Writer(), "pageSizeStr strconv.Atoi err: %s", err)
	}

	req := apis.NewDescribeImagesRequestWithAllParams(regionId, &imageSource, &platForm, idsArr, &rootDeviceType, &status, &pageNumber, &pageSize)

	resp, err := vmClient.DescribeImages(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}