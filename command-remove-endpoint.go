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

var flagsDeactivateEndpoint *api.DeactivateEndpointOptions = &api.ActivateEndpointOptions{}
var flagsRemoveEndpoint *api.RemoveEndpointOptions = &api.RemoveEndpointOptions{}

var commandRemoveEndpoint cli.Command = cli.Command{
	Name:        "remove-endpoint",
	ArgsUsage:   "",
	Description: "Remove an API that has been onboarded to Akamai.",
	HideHelp:    true,
	Action:      callRemoveEndpoint,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:        "endpoint",
			Usage:       "The unique identifier for the endpoint.",
			Destination: &flagsActivateEndpoint.APIEndPointId,
		},
		cli.IntFlag{
			Name:        "version",
			Usage:       "The endpoint version number.",
			Destination: &flagsActivateEndpoint.VersionNumber,
		},
		cli.StringSliceFlag{
			Name:  "network",
			Usage: "[Staging and/or Production] Which network to remove the endpoint on, pass multiple flags if needed.",
		},
		cli.StringSliceFlag{
			Name:  "notificationRecipient",
			Usage: "Email address(es) to notify when the removal is complete, pass multiple flags if needed.",
		},
		cli.StringFlag{
			Name:        "notes",
			Usage:       "Comments on the removal",
			Destination: &flagsActivation.Notes,
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

	flagsActivation.NotificationRecipients = c.StringSlice("notificationRecipient")
	flagsActivation.Networks = c.StringSlice("network")

	activation, err := api.ActivateEndpoint(flagsActivateEndpoint, flagsActivation)
	return output(c, activation, err)
}
