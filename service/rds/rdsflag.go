package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
)

var regionIdFlag = &cli.Flag{
	Name: "region-id",
	Short: i18n.T("", "(string) 地域ID"),
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

var instanceIdFlag = &cli.Flag{
	Name: "instance-id",
	Short: i18n.T("", "(string) RDS 实例ID，唯一标识一个RDS实例"),
}

var fileNameFlag = &cli.Flag{
	Name: "file-name",
	Short: i18n.T("", "(string) 审计文件名"),
}

var nameFlag = &cli.Flag{
	Name: "name",
	Short: i18n.T("", `(string) 审计选项类别，**大小写敏感**，目前支持两种类型：<br>（1）AuditOptions开头：在disalbed参数中返回SQL Server各个版本支持的所有选项，支持的名称为<br>AuditOptions2008R2<br>AuditOptions2012<br>AuditOptions2014<br>AuditOptions2016<br>例如输入参数为"AuditOptions2016"，则在disabled字段中返回SQL Server 2016 版本所支持的所有的审计选项<br>（2）AuditDefault开头：京东云建议的默认选项,在enabled参数中返回建议开启的选项，在disabled参数中返回不开启的选项，支持的名称为：<br>AuditDefault2008R2<br>AuditDefault2012<br>AuditDefault2014<br>AuditDefault2016<br>例如输入参数为"AuditDefault2016"，则在enabled字段返回SQL Server 2016 版本中京东云建议开启的审计选项，在disabled字段中返回建议不开启的选项`),
}

var backupIdFlag = &cli.Flag{
	Name: "backup-id",
	Short: i18n.T("", "获取整个备份或备份中单个文件的下载链接。<br>- 当输入参数中有文件名时，获取该文件的下载链接。<br>- 输入参数中无文件名时，获取整个备份的下载链接。<br>由于备份机制的差异，使用该接口下载备份时，SQL Server必须输入文件名，每个文件逐一下载，不支持下载整个备份。SQL Server备份中的文件名（不包括后缀）即为备份的数据库名。例如文件名为my_test_db.bak，表示该文件是my_test_db数据库的备份。<br>MySQL可下载整个备份集，但不支持单个文件的下载。。"),
}

var urlExpirationSecondFlag = &cli.Flag{
	Name: "url-expiration-second",
	Short: i18n.T("", "(string) 指定下载链接的过期时间，单位秒,缺省为86400秒，即24小时。<br>- MySQL：不支持该参数，只能是默认值<br>- SQL Server：支持"),
	Required: true,
	DefaultValue: "86400",
}

var autoFlag = &cli.Flag{
	Name: "auto",
	Short: i18n.T("", "(int) 查询备份类型，0为手动备份，1为自动备份，不传表示全部. <br>**- 测试参数，仅支持SQL Server，后续可能被其他参数取代**"),
}

var backupTypeFilterFlag = &cli.Flag{
	Name: "backup-type-filter",
	Short: i18n.T("", "(string) 返回backupType等于指定值的备份列表。full为全量备份，diff为增量备份<br>**- 测试参数，仅支持SQL Server，后续可能被其他参数取代**"),
}

var dbNameFilterFlag = &cli.Flag{
	Name: "db-name-filter",
	Short: i18n.T("", "(string) 返回dbName等于指定值的备份列表，不传或为空返回全部<br>**- 测试参数，仅支持SQL Server，后续可能被其他参数取代**"),
}

var backupTimeRangeStartFilterFlag = &cli.Flag{
	Name: "backup-time-range-start-filter",
	Short: i18n.T("", "(string) 返回备份开始时间大于该时间的备份列表<br>**- 测试参数，仅支持SQL Server，后续可能被其他参数取代**"),
}

var backupTimeRangeEndFilterFlag = &cli.Flag{
	Name: "backup-time-range-end-filter",
	Short: i18n.T("", "(string) 返回备份开始时间小于等于该时间的备份列表<br>**- 测试参数，仅支持SQL Server，后续可能被其他参数取代**"),
}

var binlogBackupIdFlag = &cli.Flag{
	Name: "binlog-backup-id",
	Short: i18n.T("", "(string) binlog的备份ID，可以通过describeBinlogs获得"),
}

var startTimeFlag = &cli.Flag{
	Name: "start-time",
	Short: i18n.T("", "(string) 查询开始时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到结束时间不超过三天"),
}

var endTimeFlag = &cli.Flag{
	Name: "end-time",
	Short: i18n.T("", "(string) 查询结束时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到结束时间不超过三天"),
}

var dbNameFlag = &cli.Flag{
	Name: "db-name",
	Short: i18n.T("", "(string) 数据库名。如果不指定数据库名，则返回所有数据库列表<br>- **MySQL：不支持该字段**<br>- **SQL Server：支持该字段**"),
}

var queryIndexTypeFlag = &cli.Flag{
	Name: "query-type",
	Short: i18n.T("", "(string) 查询类型，不同的查询类型按照相应的字段从高到低返回结果。<br>支持如下类型：<br>Missing：缺失索引<br>Size：索引大小，单位KB<br>Updates：索引更新次数<br>Scans：表扫描次数<br>Used：最少使用<br>"),
}

var dbFlag = &cli.Flag{
	Name: "db",
	Short: i18n.T("", "(string) 需要查询的数据库名，多个数据库名之间用英文逗号分隔，默认所有数据库"),
}

var queryTypeFlag = &cli.Flag{
	Name: "query-type",
	Short: i18n.T("", "(string) 查询类型，不同的查询类型按照相应的字段从高到低返回结果。<br>支持如下类型：<br>ExecutionCount：执行次数<br>LastRows：上次返回行数<br>ElapsedTime：平均执行时间<br>CPUTime：平均CPU时间<br>LogicalReads：平均逻辑读<br>LogicalWrites：平均逻辑写<br>PhysicalReads：平均物理读<br>"),
}

var thresholdFlag = &cli.Flag{
	Name: "threshold",
	Short: i18n.T("", "(int) 只返回查询条件大于等于threshold的记录，默认为0"),
	Required: true,
	DefaultValue: "0",
}

var slowStartTimeFlag = &cli.Flag{
	Name: "start-time",
	Short: i18n.T("", "(string) 慢日志开始时间,格式为：YYYY-MM-DD HH:mm:ss,开始时间到结束时间不能大于30天,结束时间不能大于当前时间 "),
}

var slowEndTimeFlag = &cli.Flag{
	Name: "end-time",
	Short: i18n.T("", "(string) 慢日志结束时间,格式为：YYYY-MM-DD HH:mm:ss,开始时间到结束时间不能大于30天,结束时间不能大于当前时间"),
}

var slowDbNameFlag = &cli.Flag{
	Name: "db-name",
	Short: i18n.T("", "(string) 查询哪个数据库的慢日志，不填表示返回所有数据库的慢日志"),
}