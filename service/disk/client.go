package disk

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/config"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/disk/client"
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
	common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
	disk "github.com/jdcloud-api/jdcloud-sdk-go/services/disk/models"
)

func GetProfile(ctx *cli.Context) config.Profile{
	conf, err := config.LoadConfiguration()
	if err != nil {
		cli.Errorf(ctx.Writer(), "load configuration failed %s", err)
	}
	profile := conf.GetCurrentProfile(ctx)
	return profile
}


func GetDiskClient(ctx *cli.Context) *client.DiskClient{
	profile := GetProfile(ctx)
	credentials := core.NewCredentials(profile.AccessKey, profile.SecretKey)
	diskClient := client.NewDiskClient(credentials)
	return diskClient
}

func Str2Filters(filtersStr string) (filters []common.Filter) {

	return
}

func Str2Tags(tagsStr string) (tags []disk.TagFilter) {
	return
}