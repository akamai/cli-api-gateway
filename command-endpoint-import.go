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
	"strconv"

	api "github.com/akamai/AkamaiOPEN-edgegrid-golang/api-endpoints-v2"
	akamai "github.com/akamai/cli-common-golang"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var flagsCreate *api.CreateEndpointFromFileOptions = &api.CreateEndpointFromFileOptions{}
var flagsUpdate *api.UpdateEndpointFromFileOptions = &api.UpdateEndpointFromFileOptions{}

var commandImportEndpoint cli.Command = cli.Command{
	Name:        "import",
	ArgsUsage:   "",
	Description: "This operation creates or updates an endpoint by importing an API definition file, in Swagger 2.0 or RAML 0.8 format.",
	HideHelp:    true,
	Action:      callImportEndpoint,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "format",
			Usage: "Format of the input file, either 'raml', or 'swagger'",
		},
		cli.StringFlag{
			Name:  "file",
			Usage: "Absolute path to the file containing the API definition.",
		},
		cli.StringFlag{
			Name:        "endpoint",
			Usage:       "The unique identifier for the endpoint.",
			Destination: &flagsUpdate.EndpointId,
		},
		cli.StringFlag{
			Name:        "version",
			Usage:       "The endpoint version number.",
			Destination: &flagsUpdate.Version,
		},
		cli.StringFlag{
			Name:        "contract",
			Usage:       "The unique identifier for the contract under which to provision the endpoint.",
			Destination: &flagsCreate.ContractId,
		},
		cli.StringFlag{
			Name:        "group",
			Usage:       "The unique identifier for the group under which to provision the endpoint.",
			Destination: &flagsCreate.GroupId,
		},
	},
}

func callImportEndpoint(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Importing API endpoint...",
		fmt.Sprintf("Importing API endpoint...... [%s]", color.GreenString("OK")),
	)

	flagsCreate.File = c.String("file")
	flagsUpdate.File = c.String("file")
	flagsCreate.Format = c.String("format")
	flagsUpdate.Format = c.String("format")

	if c.String("file") == "" && hasSTDIN() == true {
		// TODO: windows support?
		flagsCreate.File = "/dev/stdin"
		flagsUpdate.File = "/dev/stdin"
	}

	if c.String("version") == "" {
		flagsUpdate.Version = "latest"
	}

	var endpoint *api.Endpoint

	if flagsUpdate.EndpointId != "" {
		endpoint, err = api.GetVersion(&api.GetVersionOptions{
			flagsUpdate.EndpointId,
			flagsUpdate.Version,
		})

		if err != nil {
			return output(c, endpoint, err)
		}

		if api.IsActive(endpoint, "production") || api.IsActive(endpoint, "staging") {
			endpoint, err = api.CloneVersion(&api.CloneVersionOptions{
				flagsUpdate.EndpointId,
				flagsUpdate.Version,
			})

			if err != nil {
				return output(c, endpoint, err)
			}

			flagsUpdate.EndpointId = strconv.Itoa(endpoint.APIEndPointID)
			flagsUpdate.Version = strconv.Itoa(endpoint.VersionNumber)
		}

		endpoint, err = api.UpdateEndpointFromFile(flagsUpdate)
	} else {
		endpoint, err = api.CreateEndpointFromFile(flagsCreate)
	}

	return output(c, endpoint, err)
}
