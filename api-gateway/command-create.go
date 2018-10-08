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

var flagsCreateEndpoint *api.CreateEndpointOptions = &api.CreateEndpointOptions{}

var commandCreateEndpoint cli.Command = cli.Command{
	Name:        "create",
	ArgsUsage:   "",
	Description: "This operation creates an empty API endpoint.",
	HideHelp:    true,
	Action:      callCreateEndpoint,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "name",
			Usage:       "The name of the endpoint. Must be unique in the account.",
			Destination: &flagsCreateEndpoint.Name,
		},
		cli.StringFlag{
			Name:        "base-path",
			Usage:       "The URL path that serves as a root prefix for all resources' resourcePath values for the endpoint. This is / if empty. Do not append a / character to the path.",
			Destination: &flagsCreateEndpoint.BasePath,
		},
		cli.StringSliceFlag{
			Name:  "hostname",
			Usage: "The hostname that may receive traffic for the endpoint. At least one hostname is required, multiple hostnames can be added.",
		},
		cli.StringFlag{
			Name:        "contract",
			Usage:       "The unique identifier for the contract under which to provision the endpoint.",
			Destination: &flagsCreateEndpoint.ContractId,
		},
		cli.IntFlag{
			Name:        "group",
			Usage:       "The unique identifier for the group under which to provision the endpoint.",
			Destination: &flagsCreateEndpoint.GroupId,
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "Output JSON format",
		},
	},
}

func callCreateEndpoint(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Creating new API endpoint...",
		fmt.Sprintf("Creating new API endpoint...... [%s]", color.GreenString("OK")),
	)

	flagsCreateEndpoint.Hostnames = c.StringSlice("hostname")

	endpoint, err := api.CreateEndpoint(flagsCreateEndpoint)

	return output(c, endpoint, err)
}
