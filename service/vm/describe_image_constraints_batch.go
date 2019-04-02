package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func NewDescribeImageConstraintsBatch() *cli.Command {
	describeImageConstraintsBatch := &cli.Command{
		Name: "describe-image-constraints-batch",
		Short: i18n.T("", "批量查询镜像的实例规格限制; 通过此接口可以查看镜像不支持的实例规格。只有官方镜像、第三方镜像有实例规格的限制，个人的私有镜像没有此限制。"),
		Usage: "describe-image-constraints-batch [-h] [--region-id REGIONID] [--ids IDS] [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeImageConstraintsBatch(ctx)
			return nil
		},
	}
	describeImageConstraintsBatch.Flags().Add(idsFlag)
	describeImageConstraintsBatch.Flags().Add(headersFlag)
	describeImageConstraintsBatch.Flags().Add(regionIdFlag)
	describeImageConstraintsBatch.Flags().Add(inputJsonFlag)
	return describeImageConstraintsBatch
}

func doDescribeImageConstraintsBatch(ctx *cli.Context) {
	profile := GetProfile(ctx)
	vmClient := GetVmClient(ctx)
	regionIdFlag.DefaultValue = profile.RegionId
	regionIdFlag.Required = true
	regionId, _ := regionIdFlag.GetValue()
	ids, _ := idsFlag.GetValue()
	//
	idsArr := Str2Arr(ids)
	req := apis.NewDescribeImageConstraintsBatchRequestWithAllParams(regionId, idsArr)

	resp, err := vmClient.DescribeImageConstraintsBatch(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}

// 字符串转换成字符串数组
func Str2Arr(s string) (strArr []string){
	// 判断字符串是否被[]包裹
	// 去掉首尾字符
	// 逗号分隔成字符串数组
	// 返回数组
	// [aaa, bbb,... ]
	if ok, _ := regexp.MatchString("^[[].*[]]$", s); ok {
		s = strings.TrimLeft(s, "[")
		s = strings.TrimRight(s, "]")
	}else {
		fmt.Println("Parameter ids is invalid!")
	}
	strArr = strings.Split(s, ",")
	return
}

