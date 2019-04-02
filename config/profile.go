package config

import (
	"fmt"
	"regexp"
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vm/client"
	"jdcloud-cli/cli"
	"os"
)

type Profile struct {
	Name         string         `json:"name"`
	AccessKey    string         `json:"access_key"`
	SecretKey    string         `json:"secret_key"`
	RegionId     string         `json:"region_id"`
	Endpoint     string         `json:"endpoint"`
	Scheme       string         `json:"scheme"`
	Timeout      int            `json:"timeout"`
	OutputFormat string         `json:"output_format"`
	parent       *Configuration `json:"-"`
}

func NewProfile(name string) Profile {
	return Profile{
		Name: name,
		RegionId: "cn-north-1",
		Endpoint: "www.jdcloud-api.com",
		Scheme: "https",
		Timeout: 20,
		OutputFormat: "json",
	}
}

func (cp *Profile) Validate() error {
	if cp.RegionId == "" {
		return fmt.Errorf("region can't be empty")
	}

	if !IsRegion(cp.RegionId) {
		return fmt.Errorf("invalid region %s", cp.RegionId)
	}

	return nil
}

func (cp *Profile) GetParent() *Configuration {
	return cp.parent
}

func (cp *Profile) OverwriteWithFlags(ctx *cli.Context) {
	cp.AccessKey = AccessKeyFlag(ctx.Flags()).GetStringOrDefault(cp.AccessKey)
	cp.SecretKey = SecretKeyFlag(ctx.Flags()).GetStringOrDefault(cp.SecretKey)
	cp.RegionId = RegionFlag(ctx.Flags()).GetStringOrDefault(cp.RegionId)
	cp.Endpoint = EndpointFlag(ctx.Flags()).GetStringOrDefault(cp.Endpoint)
	cp.Scheme = SchemeFlag(ctx.Flags()).GetStringOrDefault(cp.Scheme)
	cp.Timeout = TimeoutFlag(ctx.Flags()).GetIntegerOrDefault(cp.Timeout)

	if cp.AccessKey == "" {
		cp.AccessKey = os.Getenv("ACCESS_KEY")
	}

	if cp.SecretKey == "" {
		cp.SecretKey = os.Getenv("SECRET_KEY")
	}

	if cp.RegionId == "" {
		cp.RegionId = os.Getenv("REGION")
	}

}

func (cp *Profile) ValidateKey() error {
	if len(cp.AccessKey) == 0 {
		return fmt.Errorf("invalid access_key %s", cp.AccessKey)
	}
	if len(cp.SecretKey) == 0 {
		return fmt.Errorf("invalid secret_key %s", cp.SecretKey)
	}
	return nil
}

func (cp *Profile) GetVmClient() *client.VmClient {
	credentials := core.NewCredentials(cp.AccessKey, cp.SecretKey)
	return client.NewVmClient(credentials)
}

func IsRegion(region string) bool {
	if match, _ := regexp.MatchString("^[a-zA-Z0-9-]*$", region); !match {
		return false
	}
	return true
}
