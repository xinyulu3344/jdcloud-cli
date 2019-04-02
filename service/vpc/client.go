package vpc

import (
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/client"
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
	"jdcloud-cli/cli"
	common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
	"jdcloud-cli/config"
)

func GetProfile(ctx *cli.Context) config.Profile{
	conf, err := config.LoadConfiguration()
	if err != nil {
		cli.Errorf(ctx.Writer(), "load configuration failed %s", err)
	}
	profile := conf.GetCurrentProfile(ctx)
	return profile
}


func GetVpcClient(ctx *cli.Context) *client.VpcClient{
	profile := GetProfile(ctx)
	credentials := core.NewCredentials(profile.AccessKey, profile.SecretKey)
	diskClient := client.NewVpcClient(credentials)
	return diskClient
}

func Str2Filters(filtersStr string) (filters []common.Filter) {

	return
}

