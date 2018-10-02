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

var commandAddCollectionKey cli.Command = cli.Command{
	Name:        "add-collection-key",
	ArgsUsage:   "",
	Description: "This operation adds a new key to a collection.",
	HideHelp:    true,
	Action:      callAddCollectionKey,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "collection",
			Usage: "The collection ID to modify.",
		},
		cli.StringFlag{
			Name:  "key",
			Usage: "The name of the key.",
		},
		cli.StringFlag{
			Name:  "value",
			Usage: "The value of the key.",
		},
	},
}

func callAddCollectionKey(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Adding key to collection...",
		fmt.Sprintf("Adding key to collection...... [%s]", color.GreenString("OK")),
	)

	collection, err := api.CollectionAddKey(
		c.Int("collection"),
		c.String("name"),
		c.String("value"),
	)

	return output(c, collection, err)
}
