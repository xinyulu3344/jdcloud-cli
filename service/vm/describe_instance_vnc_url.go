package vm

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"encoding/json"
	"fmt"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
)

func NewDescribeInstanceVncUrl() *cli.Command {
	describeInstanceVncUrl := &cli.Command{
		Name: "describe-instance-vnc-url",
		Short: i18n.T("", "获取云主机vnc，用于连接管理云主机; vnc地址的有效期为1个小时，调用接口获取vnc地址后如果1个小时内没有使用，vnc地址自动失效，再次使用需要重新获取。"),
		Usage: "describe-instance-vnc-url [-h] [--region-id REGIONID] --instance-id INSTANCEID [--input-json INPUT_JSON] [--headers HEADERS]",
		Run: func(ctx *cli.Context, args []string) error {

			doDescribeInstanceVncUrl(ctx)
			return nil
		},
	}

	describeInstanceVncUrl.Flags().Add(regionIdFlag)
	describeInstanceVncUrl.Flags().Add(instanceIdFlag)
	describeInstanceVncUrl.Flags().Add(inputJsonFlag)
	describeInstanceVncUrl.Flags().Add(headersFlag)

	return describeInstanceVncUrl
}

func doDescribeInstanceVncUrl(ctx *cli.Context) {

	profile := GetProfile(ctx)
	vmClient := GetVmClient(ctx)

	regionIdFlag.Required = true
	regionIdFlag.DefaultValue = profile.RegionId

	regionId, _ := regionIdFlag.GetValue()
	instanceId, _ := instanceIdFlag.GetValue()
	//inputJson, _ := inputJsonFlag.GetValue()
	//headers, _ := headersFlag.GetValue()

	req := apis.NewDescribeInstanceVncUrlRequestWithAllParams(regionId, instanceId)
	resp, err := vmClient.DescribeInstanceVncUrl(req)
	if err != nil {
		return
	}
	resultJsonByte, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(resultJsonByte))
}
