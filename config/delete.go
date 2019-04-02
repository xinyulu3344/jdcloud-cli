package config

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"io"
)

func NewConfigureDelet() *cli.Command {
	deleteCmd := &cli.Command{
		Name: "delete",
		Usage: "delete --profile <profileName>",
		Short: i18n.T("delete profile","删除指定配置集"),
		Run: func(ctx *cli.Context, args []string) error {
			profileName, ok := ProfileFlag(ctx.Flags()).GetValue()
			if !ok {
				cli.Errorf(ctx.Writer(),  "missing --profile <profileName>\n")
				cli.Notice(ctx.Writer(),  "\nusage:\n  jdc configure delete --profile <profileName>\n")
				return nil
			}
			doConfigureDelete(ctx.Writer(), profileName)
			return nil
		},
	}
	deleteCmd.Flags().Add(NewProfileFlag())
	return deleteCmd
}


func doConfigureDelete(w io.Writer, profileName string) {
	conf, err := LoadConfiguration()
	if err != nil {
		cli.Errorf(w, "ERROR: load configure failed: %v\n", err)
	}
	deleted := false
	r := make([]Profile, 0)
	for _, p := range conf.Profiles {
		if p.Name != profileName {
			r = append(r, p)
		} else {
			deleted = true
		}
	}

	if !deleted {
		cli.Errorf(w, "Error: configuration profile `%s` not found\n", profileName)
	}

	conf.Profiles = r
	if conf.CurrentProfile == profileName {
		if len(conf.Profiles) > 0 {
			conf.CurrentProfile = conf.Profiles[0].Name
		} else {
			conf.CurrentProfile = DefaultConfigProfileName
		}
	}
	err = SaveConfiguration(conf)
	if err != nil {
		cli.Errorf(w, "Error: save configuration failed %s\n", err)
	}
}