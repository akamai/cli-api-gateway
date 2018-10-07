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

var commandPrivacy cli.Command = cli.Command{
	Name:        "privacy-add",
	ArgsUsage:   "",
	Description: "Make an endpoint|resource|method public or private",
	HideHelp:    true,
	Action:      callPrivacy,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "endpoint",
			Usage: "The unique identifier for the endpoint.",
		},
		cli.StringFlag{
			Name:  "version",
			Usage: "The endpoint version number.",
		},
		cli.BoolFlag{
			Name:  "public",
			Usage: "Make this endpoint public.",
		},
		cli.BoolFlag{
			Name:  "private",
			Usage: "Make this endpoint private.",
		},
		cli.StringFlag{
			Name:  "resource",
			Usage: "The resource name to apply the settings to.",
		},
	},
}

func callPrivacy(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Updating privacy...",
		fmt.Sprintf("Updating privacy...... [%s]", color.GreenString("OK")),
	)

	settings, err := api.GetAPIPrivacySettings(c.Int("endpoint"), c.Int("version"))
	if err != nil {
		return output(c, nil, err)
	}

	if c.String("resource") != "" {
		resources, err := api.GetResourceMulti(c.Int("endpoint"), c.String("resource"), c.Int("version"))
		if err != nil {
			return output(c, nil, err)
		}

		for _, resource := range resources {
			newSettings := api.APIPrivacyResource{ResourceSettings: api.ResourceSettings{}}
			newSettings.Path = resource.ResourcePath

			for id, sresource := range settings.Resources {
				if resource.APIResourceID == id {
					newSettings = sresource
				}
			}

			if c.Bool("public") {
				newSettings.Methods = resource.APIResourceMethodsNameLists
			} else {
				newSettings.Methods = []string{}
			}

			newSettings.Public = c.Bool("public")
			settings.Resources[resource.APIResourceID] = newSettings
		}
	} else {
		settings.Public = c.Bool("public")
	}

	_, err = api.UpdateAPIPrivacySettings(c.Int("endpoint"), c.Int("version"), settings)

	return output(c, settings, err)
}
