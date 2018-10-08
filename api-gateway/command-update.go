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

var flagsUpdateEndpoint *api.Endpoint = &api.Endpoint{}

var commandUpdateEndpoint cli.Command = cli.Command{
	Name:        "update",
	ArgsUsage:   "",
	Description: "This operation updates an API endpoint version.",
	HideHelp:    true,
	Action:      callUpdateEndpoint,
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
		cli.StringFlag{
			Name:        "name",
			Usage:       "The name of the endpoint. Must be unique in the account.",
			Destination: &flagsUpdateEndpoint.APIEndPointName,
		},
		cli.StringFlag{
			Name:        "description",
			Usage:       "A description of the endpoint.",
			Destination: &flagsUpdateEndpoint.Description,
		},
		cli.StringFlag{
			Name:        "base-path",
			Usage:       "The URL path that serves as a root prefix for all resources' resourcePath values for the endpoint. This is / if empty. Do not append a / character to the path.",
			Destination: &flagsUpdateEndpoint.BasePath,
		},
		cli.StringSliceFlag{
			Name:  "hostname",
			Usage: "The hostname that may receive traffic for the endpoint. At least one hostname is required, multiple hostnames can be added.",
		},
		cli.StringFlag{
			Name:        "scheme",
			Usage:       "The URL scheme to which the endpoint may respond, either http, https, or http/https for both.",
			Destination: &flagsUpdateEndpoint.APIEndPointScheme,
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "Output JSON format",
		},
	},
}

func callUpdateEndpoint(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Updating API endpoint...",
		fmt.Sprintf("Updating API endpoint...... [%s]", color.GreenString("OK")),
	)

	if flagsUpdateEndpoint.VersionNumber == 0 {
		flagsUpdateEndpoint.VersionNumber, err = api.GetLatestVersionNumber(flagsUpdateEndpoint.APIEndPointID)
		if err != nil {
			return output(c, nil, err)
		}
	}

	endpoint, err := api.GetVersion(flagsUpdateEndpoint.APIEndPointID, flagsUpdateEndpoint.VersionNumber)
	if err != nil {
		return output(c, endpoint, err)
	}

	if api.IsActive(endpoint, "production") || api.IsActive(endpoint, "staging") {
		endpoint, err = api.CloneVersion(flagsUpdateEndpoint.APIEndPointID, flagsUpdateEndpoint.VersionNumber)
		if err != nil {
			return output(c, endpoint, err)
		}

		flagsUpdateEndpoint.APIEndPointID = endpoint.APIEndPointID
		flagsUpdateEndpoint.VersionNumber = endpoint.VersionNumber
	}

	flagsUpdateEndpoint.APIEndPointHosts = c.StringSlice("hostname")

	endpoint, err = api.ModifyVersion(flagsUpdateEndpoint)

	return output(c, endpoint, err)
}
