package rds

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/config"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/client"
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
	common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

func GetProfile(ctx *cli.Context) config.Profile {
	conf, err := config.LoadConfiguration()
	if err != nil {
		cli.Errorf(ctx.Writer(), "load configuration failed %s", err)
	}
	profile := conf.GetCurrentProfile(ctx)
	return profile
}

func GetRdsClient(ctx *cli.Context) *client.RdsClient {
	profile := GetProfile(ctx)
	credentials := core.NewCredentials(profile.AccessKey, profile.SecretKey)
	rdsClient := client.NewRdsClient(credentials)
	return rdsClient
}

func Str2Filters(filtersStr string) (filters []common.Filter) {

	return
}
