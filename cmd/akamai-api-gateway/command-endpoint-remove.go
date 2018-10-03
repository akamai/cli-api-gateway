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

var commandRemoveEndpoint cli.Command = cli.Command{
	Name:        "remove",
	ArgsUsage:   "",
	Description: "Remove an API endpoint that has been onboarded to Akamai.",
	HideHelp:    true,
	Action:      callRemoveEndpoint,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "endpoint",
			Usage: "The unique identifier for the endpoint.",
		},
	},
}

func callRemoveEndpoint(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Removing API endpoint...",
		fmt.Sprintf("Removing API endpoint...... [%s]", color.GreenString("OK")),
	)

	endpoint, err := api.RemoveEndpoint(c.Int("endpoint"))
	return output(c, endpoint, err)
}
