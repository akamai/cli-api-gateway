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

	api2 "github.com/akamai/AkamaiOPEN-edgegrid-golang/api-endpoints-v2"
	api "github.com/akamai/AkamaiOPEN-edgegrid-golang/apikey-manager-v1"
	akamai "github.com/akamai/cli-common-golang"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var commandAclResource cli.Command = cli.Command{
	Name:        "acl-resource",
	ArgsUsage:   "",
	Description: "This operation allows or denies a resource(s) on the key collection ACL",
	HideHelp:    true,
	Action:      callAclResource,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "collection",
			Usage: "The collection name or ID to modify.",
		},
		cli.IntFlag{
			Name:  "endpoint",
			Usage: "The endpoint ID the resources are associated with",
		},
		cli.IntFlag{
			Name:  "version",
			Usage: "The endpoint version.",
		},
		cli.StringSliceFlag{
			Name:  "resource",
			Usage: "The resource name or ID to allow/deny on the ACL",
		},
		cli.BoolFlag{
			Name:  "allow",
			Usage: "The resource(s) should be allowed in the ACL",
		},
		cli.BoolFlag{
			Name:  "deny",
			Usage: "The resource(s) should be denied in the ACL",
		},
	},
}

func callAclResource(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Updating key collection ACL...",
		fmt.Sprintf("Updating key collection ACL...... [%s]", color.GreenString("OK")),
	)

	version := c.Int("version")
	if version == 0 {
		version, err = api2.GetLatestVersionNumber(c.Int("endpoint"))
		if err != nil {
			return output(c, nil, err)
		}
	}

	collections, err := api.GetCollectionMulti(c.String("collection"))
	if err != nil {
		return output(c, nil, err)
	}

	for _, collection := range *collections {
		var acl []string

		for _, e := range c.StringSlice("resource") {
			resources, err := api2.GetResourceMulti(c.Int("endpoint"), e, version)
			if err != nil {
				return output(c, nil, err)
			}
			for _, r := range resources {
				acl = append(acl, fmt.Sprintf("RESOURCE-%d", r.APIResourceLogicID))
			}
		}

		if c.Bool("allow") {
			_, err := api.CollectionAclAllow(collection.Id, acl)
			if err != nil {
				return output(c, nil, err)
			}
		} else {
			_, err := api.CollectionAclDeny(collection.Id, acl)
			if err != nil {
				return output(c, nil, err)
			}
		}
	}

	return output(c, nil, err)
}
