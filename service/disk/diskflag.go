package disk

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
)

var regionIdFlag = &cli.Flag{
	Name: "region-id",
	Short: i18n.T("", "(string) 地域ID"),
}

var diskIdFlag = &cli.Flag{
	Name: "disk-id",
	Short: i18n.T("", "(string) 云硬盘ID"),
}

var inputJsonFlag = &cli.Flag{
	Name: "input-json",
	Short: i18n.T("", `
                (json) 以json字符串或文件绝对路径形式作为输入参数。
                字符串方式举例：--input-json '{"field":"value"}';
                文件格式举例：--input-json file:///xxxx.json`),
}

var headersFlag = &cli.Flag{
	Name: "headers",
	Short: i18n.T("", `(json) 用户自定义Header，举例：'{"x-jdcloud-security-token":"abc","test":"123"}'`),
}

var pageNumberFlag = &cli.Flag{
	Name: "page-number",
	Short: i18n.T("", "页码；默认为1"),
	DefaultValue: "1",
	Required: true,
}

var pageSizeFlag = &cli.Flag{
	Name: "page-size",
	Short: i18n.T("", "分页大小；默认为20；取值范围[10, 100]"),
	DefaultValue: "20",
	Required: true,
}

var filtersFlag = &cli.Flag{
	Name: "filters",
	Short: i18n.T("", `instanceId - 云主机ID，精确匹配，支持多个; privateIpAddress - 主网卡内网主IP地址，模糊匹配，支持多个; az - 可用区，精确匹配，支持多个; vpcId - 私有网络ID，精确匹配，支持多个; status - 云主机状态，精确匹配，支持多个，<a href="http://docs.jdcloud.com/virtual-machines/api/vm_status">参考云主机状态</a>; name - 云主机名称，模糊匹配，支持单个; imageId - 镜像ID，精确匹配，支持多个; networkInterfaceId - 弹性网卡ID，精确匹配，支持多个; subnetId - 子网ID，精确匹配，支持多个; agId - 使用可用组id，支持单个; faultDomain - 错误域，支持多个`),
}

var tagsFlag = &cli.Flag{
	Name: "tags",
	Short: i18n.T("", "(array: tagFilter) Tag筛选条件"),
}

var snapshotIdFlag = &cli.Flag{
	Name: "snapshot-id",
	Short: i18n.T("", "(string) 快照ID"),
}
