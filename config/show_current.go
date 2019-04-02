package config

import (
	"jdcloud-cli/i18n"
	"jdcloud-cli/cli"
	"fmt"
)

const configureShowCurrentHelpEn = ``
const configureShowCurrentHelpZh = ``

func NewConfigureShowCurrent() *cli.Command {
	showCurrentCmd := &cli.Command{
		Name: "show-current",
		Usage: "show-current",
		Short: i18n.T("query current profile", "查看当前配置"),
		Long: i18n.T(
			configureShowCurrentHelpEn,
			configureShowCurrentHelpZh,
			),
		Run: func(ctx *cli.Context, args []string) error {
			doConfigureShowCurrent(ctx, args)
			return nil
		},
	}
	return showCurrentCmd
}

func doConfigureShowCurrent(ctx *cli.Context, args []string) {
	config, err := LoadConfiguration()
	if err != nil {
		cli.Errorf(ctx.Writer(), "load configuration failed %s", err)
	}

	profile := config.GetCurrentProfile(ctx)

	fmt.Printf("================= %s =================\n", profile.Name)
	fmt.Println("access_key:  	", profile.AccessKey)
	fmt.Println("secret_key:  	", profile.SecretKey)
	fmt.Println("region_id:  	", profile.RegionId)
	fmt.Println("endpoint:  	", profile.Endpoint)
	fmt.Println("scheme:  	", profile.Scheme)
	fmt.Println("timeout:  	", profile.Timeout)

}

