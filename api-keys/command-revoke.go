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

var commandRevokeKey cli.Command = cli.Command{
	Name:        "revoke",
	ArgsUsage:   "",
	Description: "This operation revokes a key.",
	HideHelp:    true,
	Action:      callRevokeKey,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "key",
			Usage: "The key to revoke.",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "Output JSON format",
		},
	},
}

func callRevokeKey(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Revoking key...",
		fmt.Sprintf("Revoking key...... [%s]", color.GreenString("OK")),
	)

	key, err := api.RevokeKey(c.Int("key"))

	return output(c, key, err)
}
