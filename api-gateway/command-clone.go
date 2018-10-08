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

var commandClone cli.Command = cli.Command{
	Name:        "clone",
	ArgsUsage:   "",
	Description: "This operation clones an API endpoint version.",
	HideHelp:    true,
	Action:      callClone,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:        "endpoint",
			Usage:       "The unique identifier for the endpoint.",
			Destination: &flagsUpdateEndpoint.APIEndPointID,
		},
		cli.IntFlag{
			Name:        "version",
			Usage:       "The endpoint version number.",
			Destination: &flagsUpdateEndpoint.VersionNumber,
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "Output JSON format",
		},
	},
}

func callClone(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Cloning API endpoint...",
		fmt.Sprintf("Cloning API endpoint...... [%s]", color.GreenString("OK")),
	)

	version := c.Int("version")
	if version == 0 {
		version, err = api.GetLatestVersionNumber(c.Int("endpoint"), false)
		if err != nil {
			return output(c, nil, err)
		}
	}

	rversion, err := api.CloneVersion(c.Int("endpoint"), version)
	return output(c, rversion, err)
}
