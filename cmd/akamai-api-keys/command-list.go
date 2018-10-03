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

var commandListKeyCollections cli.Command = cli.Command{
	Name:        "list-collections",
	ArgsUsage:   "",
	Description: "Retrieve a list of all key collections configured at Akamai.",
	HideHelp:    true,
	Action:      callListKeyCollections,
}

func callListKeyCollections(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Fetching key collections list...",
		fmt.Sprintf("Fetching key collections list...... [%s]", color.GreenString("OK")),
	)

	collections, err := api.ListCollections()

	return output(c, collections, err)
}
