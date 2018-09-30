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

var flagsImportEndpoint *api.ImportEndpointOptions = &api.ImportEndpointOptions{}

var commandImportEndpoint cli.Command = cli.Command{
	Name:        "import",
	ArgsUsage:   "",
	Description: "This operation imports an API definition file and creates a new endpoint based on the file contents. You either upload or specify a URL to a Swagger 2.0 or RAML 0.8 file to import details about your API.",
	HideHelp:    true,
	Action:      callImportEndpoint,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "format",
			Usage:       "Format of the input file, either 'raml', or 'swagger'",
			Destination: &flagsImportEndpoint.Format,
		},
		cli.StringFlag{
			Name:        "file",
			Usage:       "Absolute path to the file containing the API definition.",
			Destination: &flagsImportEndpoint.File,
		},
		cli.StringFlag{
			Name:        "endpoint",
			Usage:       "The unique identifier for the endpoint.",
			Destination: &flagsImportEndpoint.EndpointId,
		},
		cli.StringFlag{
			Name:        "version",
			Usage:       "The endpoint version number.",
			Destination: &flagsImportEndpoint.Version,
		},
		cli.StringFlag{
			Name:        "contract",
			Usage:       "The unique identifier for the contract under which to provision the endpoint.",
			Destination: &flagsImportEndpoint.ContractId,
		},
		cli.StringFlag{
			Name:        "group",
			Usage:       "The unique identifier for the group under which to provision the endpoint.",
			Destination: &flagsImportEndpoint.GroupId,
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

	if flagsImportEndpoint.File == "" && hasSTDIN() == true {
		// TODO: windows support?
		flagsImportEndpoint.File = "/dev/stdin"
	}

	endpoint, err := api.ImportEndpoint(flagsImportEndpoint)
	return output(c, endpoint, err)
}
