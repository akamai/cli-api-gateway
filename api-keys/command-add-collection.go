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

	api "github.com/akamai/AkamaiOPEN-edgegrid-golang/apikey-manager-v1"
	akamai "github.com/akamai/cli-common-golang"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var flagsCreateCollection *api.CreateCollectionOptions = &api.CreateCollectionOptions{}

var commandCreateCollection cli.Command = cli.Command{
	Name:        "add-collection",
	ArgsUsage:   "",
	Description: "This operation creates a new collection.",
	HideHelp:    true,
	Action:      callCreateCollection,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "name",
			Usage:       "The name of the collection.",
			Destination: &flagsCreateCollection.Name,
		},
		cli.StringFlag{
			Name:        "description",
			Usage:       "The description of the collection.",
			Destination: &flagsCreateCollection.Description,
		},
		cli.StringFlag{
			Name:        "contract",
			Usage:       "The unique identifier for the contract under which to provision the endpoint.",
			Destination: &flagsCreateCollection.ContractId,
		},
		cli.IntFlag{
			Name:        "group",
			Usage:       "The unique identifier for the group under which to provision the endpoint.",
			Destination: &flagsCreateCollection.GroupId,
		},
	},
}

func callCreateCollection(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Creating new key collection...",
		fmt.Sprintf("Creating new key collection...... [%s]", color.GreenString("OK")),
	)

	collection, err := api.CreateCollection(flagsCreateCollection)
	return output(c, collection, err)
}
