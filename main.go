// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"jdcloud-cli/cli"
	"jdcloud-cli/config"
	"os"
	"fmt"
	"jdcloud-cli/i18n"
	"jdcloud-cli/service/vm"
	"jdcloud-cli/service/disk"
	"jdcloud-cli/service/vpc"
	"jdcloud-cli/service/rds"
)

func main() {
	cli.PlatformCompatible()
	writer := cli.DefaultWriter()

	_, err := config.LoadCurrentProfile()
	if err != nil {
		cli.Errorf(writer, "ERROR: load current configuration failed %s", err)
		return
	}

	// set user agent
	userAgentFromEnv := os.Getenv("JDCLOUD_USER_AGENT")
	if userAgentFromEnv != "" {
		defaultUserAgent := config.GetUserAgent()
		userAgent := fmt.Sprintf("%s (%s)", defaultUserAgent, userAgentFromEnv)
		config.SetUserAgent(userAgent)
	}

	rootCmd := &cli.Command{
		Name:   "jdc",
		Short:  i18n.T("JDCloud Command Line Interface Version "+cli.Version, "京东云CLI命令行工具 "+cli.Version),
		Usage:  "jdc <product> <operation> [--parameter1 value1 --parameter2 value2 ...]",
		Sample: "jdc vm describe-instance-status",
		EnableUnknownFlag: true,
	}
	config.AddFlags(rootCmd.Flags())


	ctx := cli.NewCommandContext(writer)
	ctx.EnterCommand(rootCmd)
	ctx.SetCompletion(cli.ParseCompletionForShell())

	rootCmd.AddSubCommand(config.NewConfig())
	rootCmd.AddSubCommand(vm.NewVm())
	rootCmd.AddSubCommand(disk.NewDisk())
	rootCmd.AddSubCommand(vpc.NewVpc())
	rootCmd.AddSubCommand(rds.NewRds())

	rootCmd.Execute(ctx, os.Args[1:])
}
