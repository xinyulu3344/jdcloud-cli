package config

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
)

const (
	ProfileFlagName   = "profile"
	AccessKeyFlagName = "access-key"
	SecretKeyFlagName = "secret-key"
	RegionFlagName    = "region"
	EndpointFlagName  = "endpoint"
	SchemeFlagName    = "scheme"
	TimeoutFlagName   = "timeout"
)

func AddFlags(fs *cli.FlagSet) {
	fs.Add(NewProfileFlag())
	fs.Add(NewAccessKeyFlag())
	fs.Add(NewSecretKeyFlag())
	fs.Add(NewRegionFlag())
	fs.Add(NewEndpointFlag())
	fs.Add(NewSchemeFlag())
	fs.Add(NewTimeoutFlag())
}

func ProfileFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(ProfileFlagName)
}

func AccessKeyFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(AccessKeyFlagName)
}

func SecretKeyFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(SecretKeyFlagName)
}

func RegionFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(RegionFlagName)
}

func EndpointFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(EndpointFlagName)
}

func SchemeFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(SchemeFlagName)
}

func TimeoutFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(TimeoutFlagName)
}

//var OutputFlag = &cli.Flag{Category: "config",
//	Name: "output", AssignedMode: cli.AssignedOnce, Hidden: true,
//	Usage: i18n.T(
//		"* assign output format, only support json",
//		"* 指定输出格式, 目前仅支持Json")}

//varfs.Add(cli.Flag{Name: "site", AssignedMode: cli.AssignedOnce,
//Usage: i18n.T("assign site, support china/international/japan", "")})

func NewProfileFlag() *cli.Flag {
	return &cli.Flag{Category: "config",
		Name:         ProfileFlagName,
		Shorthand:    'p',
		DefaultValue: "default", Persistent: true,
		Short: i18n.T(
			"use `--profile <profileName>` to select profile",
			"使用 `--profile <profileName>` 指定操作的配置集")}
}

func NewAccessKeyFlag() *cli.Flag {
	return &cli.Flag{Category: "config",
		Name: AccessKeyFlagName, AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--access-key-id <AccessKeyId>` to assign AccessKeyId, required in AK/StsToken/RamRoleArn mode",
			"使用 `--access-key-id <AccessKeyId>` 指定AccessKeyId")}
}

func NewSecretKeyFlag() *cli.Flag {
	return &cli.Flag{Category: "config",
		Name: SecretKeyFlagName, AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--access-key-secret <AccessKeySecret>` to assign AccessKeySecret",
			"使用 `--access-key-secret <AccessKeySecret>` 指定AccessKeySecret")}
}

func NewRegionFlag() *cli.Flag {
	return &cli.Flag{Category: "config",
		Name: RegionFlagName, AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--region <regionId>` to assign region",
			"使用 `--region <regionId>` 来指定访问大区")}
}

func NewEndpointFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "config",
		Name:         EndpointFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--endpoint <ENDPOINT> to assign endpoint`",
			"使用 --endpoint <ENDPOINT> 指定endpoint",
		),
	}
}

func NewSchemeFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "config",
		Name:         SchemeFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--scheme <SCHEME>` to assign scheme",
			"使用 `--scheme <SCHEME` 指定scheme",
		),
	}
}

func NewTimeoutFlag() *cli.Flag {
	return &cli.Flag{Category: "caller",
		Name: TimeoutFlagName, AssignedMode: cli.AssignedOnce, Hidden: true,
		Short: i18n.T(
			"use `--timeout <seconds>` to set retry timeout(seconds)",
			"使用 `--timeout <seconds>` 指定请求超时时间(秒)"),
	}
}
