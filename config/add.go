package config

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"io"
	"fmt"
	"strings"
)

const configureSetHelpEn = ``
const configureSetHelpZh = ``

func NewConfigureAdd() *cli.Command {
	addCmd := &cli.Command{
		Name: "add",
		Short: i18n.T(
			"set config in non interactive mode",
			"使用非交互式方式进行配置"),
		Long: i18n.T(
			configureSetHelpEn,
			configureSetHelpZh,
			),
		Usage:"add [--profile <profileName>] ...",
		Run: func(ctx *cli.Context, args []string) error {
			doConfigureAdd(ctx, ctx.Flags(), args)
			return nil
		},
	}
	AddFlags(addCmd.Flags())
	return addCmd
}

func doConfigureAdd(ctx *cli.Context, flags *cli.FlagSet, args []string) {
	w := ctx.Writer()

	config, err := LoadConfiguration()
	if err != nil {
		cli.Errorf(w, "load configuration failed %s", err)
		return
	}

	profileName, ok := ProfileFlag(flags).GetValue()
	if !ok {
		profileName = config.CurrentProfile
	}

	profile, ok := config.GetProfile(profileName)
	if !ok {
		profile = NewProfile(profileName)
	}
	profile.AccessKey = AccessKeyFlag(flags).GetStringOrDefault(profile.AccessKey)
	profile.SecretKey = SecretKeyFlag(flags).GetStringOrDefault(profile.SecretKey)
	profile.RegionId = RegionFlag(flags).GetStringOrDefault(profile.RegionId)
	profile.OutputFormat = "json"
	profile.Endpoint = EndpointFlag(flags).GetStringOrDefault(profile.Endpoint)
	profile.Scheme = SchemeFlag(flags).GetStringOrDefault(profile.Scheme)
	profile.Timeout = TimeoutFlag(flags).GetIntegerOrDefault(profile.Timeout)
	err = profile.Validate()
	if err != nil {
		cli.Errorf(w, "fail to set configuration: %s", err.Error())
		return
	}

	config.PutProfile(profile)
	config.CurrentProfile = profile.Name
	err = SaveConfiguration(config)
	if err != nil {
		cli.Error(w,"save configuration failed %s", err)
	}

	// 如果add命令的参数为空，则通过交互式命令来完成添加
}


func configureAK(w io.Writer, cp *Profile) error {
	cli.Printf(w,"Access Key [%s]: ", MosaicString(cp.AccessKey, 3))
	cp.AccessKey = ReadInput(cp.AccessKey)
	cli.Printf(w,"Secret Key [%s]: ", MosaicString(cp.SecretKey, 3))
	cp.SecretKey = ReadInput(cp.SecretKey)
	return nil
}


func ReadInput(defaultValue string) string {
	var s string
	fmt.Scanf("%s\n", &s)
	if s == "" {
		return defaultValue
	}
	return s
}

func MosaicString(s string, lastChars int) string {
	r := len(s) - lastChars
	if r > 0 {
		return strings.Repeat("*", r) + s[r:]
	} else {
		return strings.Repeat("*", len(s))
	}
}
