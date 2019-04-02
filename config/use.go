package config

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
)

const configureUseHelpEn = ``
const configureUseHelpZh = ``

func NewConfigureUse() *cli.Command {
	useCmd := &cli.Command{
		Name: "use",
		Short: i18n.T("switch profile", "切换profile"),
		Long: i18n.T(configureUseHelpEn, configureUseHelpZh),
		Usage: "use --profile [Profile]",
		Run: func(ctx *cli.Context, args []string) error {
			profileName, ok := ProfileFlag(ctx.Flags()).GetValue()
			if !ok {
				cli.Errorf(ctx.Writer(),  "missing --profile <profileName>\n")
				cli.Notice(ctx.Writer(),  "\nusage:\n  jdc configure use --profile <profileName>\n")
				return nil
			}
			doConfigureUse(ctx, profileName)
			return nil
		},
	}
	useCmd.Flags().Add(NewProfileFlag())
	return useCmd
}

func doConfigureUse(ctx *cli.Context, profileName string) {
	w := ctx.Writer()
	used := false
	config, err := LoadConfiguration()
	if err != nil {
		cli.Errorf(ctx.Writer(), "load configuration failed %s", err)
	}

	for _, p := range config.Profiles {
		if p.Name == profileName {
			config.CurrentProfile = profileName
			SaveConfiguration(config)
			used = true
			cli.Printf(w, "configure success\n")
			break
		}
	}
	if !used {
		cli.Errorf(w, "Error: configuration profile `%s` not found\n", profileName)
	}
}