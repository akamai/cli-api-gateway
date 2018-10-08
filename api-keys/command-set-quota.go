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

var commandSetKeyQuota cli.Command = cli.Command{
	Name:        "set-quota",
	ArgsUsage:   "",
	Description: "This operation sets the quota on a key.",
	HideHelp:    true,
	Action:      callSetKeyQuota,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "collection",
			Usage: "The collection ID to modify.",
		},
		cli.IntFlag{
			Name:  "limit",
			Usage: "The quota limit.",
		},
		cli.StringFlag{
			Name:  "interval",
			Usage: "The interval at which to reset the quota limit. 1hr | 6hr | 12hr | day | week | month",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "Output JSON format",
		},
	},
}

func callSetKeyQuota(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Setting quota...",
		fmt.Sprintf("Setting quota...... [%s]", color.GreenString("OK")),
	)

	interval := "HOUR_1"
	switch c.String("interval") {
	case "1h":
		interval = "HOUR_1"
	case "1hour":
		interval = "HOUR_1"
	case "6h":
		interval = "HOUR_6"
	case "6hour":
		interval = "HOUR_6"
	case "12h":
		interval = "HOUR_12"
	case "12hour":
		interval = "HOUR_12"
	case "d":
		interval = "DAY"
	case "day":
		interval = "DAY"
	case "w":
		interval = "WEEK"
	case "week":
		interval = "WEEK"
	case "m":
		interval = "MONTH"
	case "month":
		interval = "MONTH"
	}

	collection, err := api.CollectionSetQuota(c.Int("collection"), c.Int("limit"), interval)

	return output(c, collection, err)
}
