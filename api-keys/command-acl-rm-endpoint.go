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
	"errors"
	"fmt"

	api "github.com/akamai/AkamaiOPEN-edgegrid-golang/apikey-manager-v1"
	akamai "github.com/akamai/cli-common-golang"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var commandAclRmEndpoint cli.Command = cli.Command{
	Name:        "acl-rm-endpoint",
	ArgsUsage:   "",
	Description: "This operation removes an endpoint(s) on the key collection ACL",
	HideHelp:    true,
	Action:      callAclRmEndpoint,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "collection",
			Usage: "The collection name or ID to modify.",
		},
		cli.IntSliceFlag{
			Name:  "endpoint",
			Usage: "The endpoint ID(s) to add to the ACL, multiples allowed",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "Output JSON format",
		},
	},
}

func callAclRmEndpoint(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Updating key collection ACL...",
		fmt.Sprintf("Updating key collection ACL...... [%s]", color.GreenString("OK")),
	)

	if c.String("collection") == "" {
		return output(c, nil, errors.New("Collection ID or name missing"))
	}

	collections, err := api.GetCollectionMulti(c.String("collection"))
	if err != nil {
		return output(c, nil, err)
	}

	out := api.Collections{}
	for _, collection := range *collections {
		var acl []string

		for _, e := range c.IntSlice("endpoint") {
			acl = append(acl, fmt.Sprintf("ENDPOINT-%d", e))
		}

		col, err := api.CollectionAclDeny(collection.Id, acl)
		if err != nil {
			return output(c, nil, err)
		}
		out = append(out, *col)
	}

	return output(c, out, err)
}
