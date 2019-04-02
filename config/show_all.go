package config

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/i18n"
	"io"
	"text/tabwriter"
	"fmt"
)

func NewConfigureShowAll() *cli.Command {
	return &cli.Command{
		Name:  "show-all",
		Usage: "jdc configure show-all",
		Short: i18n.T("show all config profile", "列出所有配置集"),
		Run: func(c *cli.Context, args []string) error {
			doConfigureList(c.Writer())
			return nil
		},
	}
}


func doConfigureList(w io.Writer) {
	conf, err := LoadConfiguration()
	if err != nil {
		cli.Errorf(w, "ERROR: load configure failed: %v\n", err)
	}
	tw := tabwriter.NewWriter(w, 8, 0, 1, ' ', 0)
	fmt.Fprint(tw, "Profile\t| Credential \t| Valid\t| Region\t\n")
	fmt.Fprint(tw, "---------\t| ------------------\t| -------\t| ----------------\t\n")
	for _, pf := range conf.Profiles {
		name := pf.Name
		if name == conf.CurrentProfile {
			name = name + " *"
		}
		err := pf.Validate()
		valid := "Valid"
		if err != nil {
			valid = "Invalid"
		}

		//cred := "AK:" + "***" + GetLastChars(pf.AccessKey, 3)
		cred := "AK:" + "***" + GetLastChars(pf.AccessKey, 3)
		fmt.Fprintf(tw, "%s\t| %s\t| %s\t| %s\t\n", name, cred, valid, pf.RegionId)
	}
	tw.Flush()
}