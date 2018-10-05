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

var commandApiKeys cli.Command = cli.Command{
	Name:        "api-keys",
	ArgsUsage:   "",
	Description: "Enable/Disable API keys or change other settings.",
	HideHelp:    true,
	Action:      callApiKeys,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "endpoint",
			Usage: "The unique identifier for the endpoint.",
		},
		cli.StringFlag{
			Name:  "version",
			Usage: "The endpoint version number.",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "The name of the header, query parameter, or cookie where you located the API key.",
		},
		cli.StringFlag{
			Name:  "location",
			Usage: "The location of the API key in incoming requests, either cookie, header, or query parameter.",
		},
		cli.BoolFlag{
			Name:  "enable",
			Usage: "Enable api key security",
		},
		cli.BoolFlag{
			Name:  "disable",
			Usage: "Disable api key security",
		},
	},
}

func callApiKeys(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Updating API Keys...",
		fmt.Sprintf("Updating API Keys...... [%s]", color.GreenString("OK")),
	)

	endpoint, err := api.GetVersion(&api.GetVersionOptions{
		EndpointId: c.Int("endpoint"),
		Version:    c.Int("version"),
	})

	if err != nil {
		return output(c, endpoint, err)
	}

	if c.Bool("enable") {
		ss := &api.SecurityScheme{
			SecuritySchemeType: "apikey",
			SecuritySchemeDetail: &api.SecuritySchemeDetail{
				APIKeyLocation: c.String("location"),
				APIKeyName:     c.String("name"),
			},
		}
		endpoint.SecurityScheme = ss
		endpoint, err = api.ModifyVersion(endpoint)
		return output(c, endpoint, err)
	}

	if c.Bool("disable") {
		ss := &api.SecurityScheme{
			SecuritySchemeType:   "",
			SecuritySchemeDetail: &api.SecuritySchemeDetail{},
		}
		endpoint.SecurityScheme = ss
		endpoint, err = api.ModifyVersion(endpoint)
		return output(c, endpoint, err)
	}

	return output(c, endpoint, err)
}
