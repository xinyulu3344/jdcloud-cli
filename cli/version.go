package cli

import (
	"strings"
	"jdcloud-cli/i18n"
)


// This variable is replaced in compile time
var (
	Version = "0.0.1"
)

func GetVersion() string {
	return strings.Replace(Version, " ", "-", -1)
}

func NewVersionCommand() *Command {
	return &Command{
		Name:   "version",
		Short:  i18n.T("print current version", "打印当前版本号"),
		Hidden: true,
		Run: func(ctx *Context, args []string) error {
			Printf(ctx.Writer(), "%s\n", Version)
			return nil
		},
	}
}
