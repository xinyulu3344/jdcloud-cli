package config

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"strings"
)


func NewConfig() *cli.Command {
	configure :=  &cli.Command{
		Name: "configure",
		Short: i18n.T(
			"",
			"使用该子命令配置CLI"),
		Long: i18n.T(
			"",
			"配置 CLI 子命令，可以使用该子命令配置鉴权和区域等信息"),
		Usage: "configure --profile <profileName>",
		Run: func(ctx *cli.Context, args []string) error{
			if len(args) > 0 {
				return cli.NewInvalidCommandError(args[0], ctx)
			}
			//profileName, _ := ProfileFlag(ctx.Flags()).GetValue()
			return doConfigure(ctx, args)
		},
		Hidden: false,
	}
	configure.AddSubCommand(NewConfigureAdd())
	configure.AddSubCommand(NewConfigureDelet())
	configure.AddSubCommand(NewConfigureShowAll())
	configure.AddSubCommand(NewConfigureShowCurrent())
	configure.AddSubCommand(NewConfigureUse())
	return configure
}

func doConfigure(ctx *cli.Context, args []string) error{
	if len(args) == 0 {
		NewConfig().ExecuteHelp(ctx, args)
	}
	return nil
}


func GetLastChars(s string, lastChars int) string {
	r := len(s) - lastChars
	if r > 0 {
		return s[r:]
	} else {
		return strings.Repeat("*", len(s))
	}
}