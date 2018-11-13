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

var commandConstraints cli.Command = cli.Command{
	Name:        "constraints",
	ArgsUsage:   "",
	Description: "Enable/disable KSD API Security and change settings.",
	HideHelp:    true,
	Action:      callConstraints,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "endpoint",
			Usage: "The unique identifier for the endpoint.",
		},
		cli.StringFlag{
			Name:  "version",
			Usage: "The endpoint version number.",
		},
		cli.IntFlag{
			Name:  "max-body-size",
			Usage: "The maximum allowed size of a request body.",
		},
		cli.IntFlag{
			Name:  "max-depth",
			Usage: "The maximum depth of nested data elements allowed in a request body.",
		},
		cli.IntFlag{
			Name:  "max-element-name",
			Usage: "The maximum length of an XML element name or JSON object key name allowed in a request body.",
		},
		cli.IntFlag{
			Name:  "max-integer",
			Usage: "The maximum integer value allowed in a request body.",
		},
		cli.IntFlag{
			Name:  "max-values",
			Usage: "The maximum number of XML elements, JSON object keys, or array items allowed in a request body.",
		},
		cli.IntFlag{
			Name:  "max-string",
			Usage: "The maximum length of any string value in a request body.",
		},
		cli.BoolFlag{
			Name:  "enable",
			Usage: "Enable api key security",
		},
		cli.BoolFlag{
			Name:  "disable",
			Usage: "Disable api key security",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "Output JSON format",
		},
	},
}

func callConstraints(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Updating KSD API Security Settings...",
		fmt.Sprintf("Updating KSD API Security Settings...... [%s]", color.GreenString("OK")),
	)

	version := c.Int("version")
	if version == 0 {
		version, err = api.GetLatestVersionNumber(c.Int("endpoint"), true)
		if err != nil {
			return output(c, nil, err)
		}
	}

	endpoint, err := api.GetVersion(c.Int("endpoint"), version)
	if err != nil {
		return output(c, endpoint, err)
	}

	if c.Bool("enable") {
		endpoint.AkamaiSecurityRestrictions.MaxJsonxmlElement = c.Int("max-element")
		endpoint.AkamaiSecurityRestrictions.MaxElementNameLength = c.Int("max-element-name")
		endpoint.AkamaiSecurityRestrictions.MaxDocDepth = c.Int("max-depth")
		endpoint.AkamaiSecurityRestrictions.MaxStringLength = c.Int("max-string")
		endpoint.AkamaiSecurityRestrictions.MaxBodySize = c.Int("max-body-size")
		endpoint.AkamaiSecurityRestrictions.MaxIntegerValue = c.Int("max-integer")
	}

	if c.Bool("disable") {
		endpoint.AkamaiSecurityRestrictions.MaxJsonxmlElement = 0
		endpoint.AkamaiSecurityRestrictions.MaxElementNameLength = 0
		endpoint.AkamaiSecurityRestrictions.MaxDocDepth = 0
		endpoint.AkamaiSecurityRestrictions.MaxStringLength = 0
		endpoint.AkamaiSecurityRestrictions.MaxBodySize = 0
		endpoint.AkamaiSecurityRestrictions.MaxIntegerValue = 0
	}
	endpoint, err = api.ModifyVersion(endpoint)
	endpointSecurity := &api.EndpointSecurity{
		APIEndPointID:              endpoint.APIEndPointID,
		APIEndPointName:            endpoint.APIEndPointName,
		SecurityScheme:             endpoint.SecurityScheme,
		AkamaiSecurityRestrictions: endpoint.AkamaiSecurityRestrictions,
	}

	return output(c, endpointSecurity, err)
}
