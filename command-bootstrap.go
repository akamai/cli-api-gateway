// Copyright 2018. Akamai Technologies, Inc
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
	api "github.com/johannac/AkamaiOPEN-edgegrid-golang/api-endpoints-v2"
	// api "github.com/akamai/AkamaiOPEN-edgegrid-golang/configdns-v1"
	akamai "github.com/akamai/cli-common-golang"

	"github.com/urfave/cli"
)

var commandLocator akamai.CommandLocator = func() ([]cli.Command, error) {
	commands := []cli.Command{
		cli.Command{
			Name:        "list",
			Description: "List commands",
			Action:      akamai.CmdList,
		},
		cli.Command{
			Name:         "help",
			Description:  "Displays help information",
			ArgsUsage:    "[command] [sub-command]",
			Action:       akamai.CmdHelp,
			BashComplete: akamai.DefaultAutoComplete,
		},
		commandListEndpoints,
	}

	return commands, nil
}

func initConfig(c *cli.Context) error {
	config, err := akamai.GetEdgegridConfig(c)
	if err != nil {
		return err
	}
	api.Init(config)
	return nil
}
