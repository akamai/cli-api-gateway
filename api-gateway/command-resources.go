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
	"fmt"

	api "github.com/akamai/AkamaiOPEN-edgegrid-golang/api-endpoints-v2"
	akamai "github.com/akamai/cli-common-golang"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var commandListResources cli.Command = cli.Command{
	Name:        "list-resources",
	ArgsUsage:   "",
	Description: "Retrieve a list resources for a given endpoint.",
	HideHelp:    true,
	Action:      callListResources,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "endpoint",
			Usage: "The unique identifier for the endpoint.",
		},
		cli.StringFlag{
			Name:  "version",
			Usage: "The endpoint version number.",
		},
	},
}

func callListResources(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Fetching endpoint resources list...",
		fmt.Sprintf("Fetching endpoint resources list...... [%s]", color.GreenString("OK")),
	)

	version := c.Int("version")
	if version == 0 {
		version, err = api.GetLatestVersionNumber(c.Int("endpoint"))
		if err != nil {
			return output(c, nil, err)
		}
	}

	resources, err := api.GetResources(c.Int("endpoint"), version)
	return output(c, resources, err)
}
