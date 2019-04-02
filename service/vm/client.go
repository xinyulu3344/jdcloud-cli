package vm

import (
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/client"
	"jdcloud-cli/config"
	"jdcloud-cli/cli"
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)


func GetProfile(ctx *cli.Context) config.Profile{
	conf, err := config.LoadConfiguration()
	if err != nil {
		cli.Errorf(ctx.Writer(), "load configuration failed %s", err)
	}
	profile := conf.GetCurrentProfile(ctx)
	return profile
}

func GetVmClient(ctx *cli.Context) *client.VmClient{
	profile := GetProfile(ctx)
	credentials := core.NewCredentials(profile.AccessKey, profile.SecretKey)
	vmClient := client.NewVmClient(credentials)
	return vmClient
}

func Str2Filters(filtersStr string) (filters []models.Filter) {

	return
}