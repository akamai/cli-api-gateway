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

var flagsListEndpoints *api.ListEndpointOptions = &api.ListEndpointOptions{}

var commandListEndpoints cli.Command = cli.Command{
	Name:        "list-endpoints",
	ArgsUsage:   "",
	Description: "Retrieve a list of APIs running through the Gateway; may be used as input to other commands",
	HideHelp:    true,
	Action:      callListEndpoints,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "json",
			Usage: "Output JSON format",
		},
		cli.StringFlag{
			Name:        "contract",
			Usage:       "Filters endpoints to a specific contract. You need to specify this along with a groupId.",
			Destination: &flagsListEndpoints.ContractId,
		},
		cli.IntFlag{
			Name:        "group",
			Usage:       "Filters endpoints to a specific group. You need to specify this along with a contractId.",
			Destination: &flagsListEndpoints.GroupId,
		},
		cli.StringFlag{
			Name:        "show",
			Usage:       "The type of endpoints to return based on their visibility status. By default the API returns ALL endpoints. You can instead decide to return ONLY_VISIBLE endpoints, or ONLY_HIDDEN endpoints.",
			Destination: &flagsListEndpoints.Show,
		},
		cli.StringFlag{
			Name:        "version-preference",
			Usage:       "The preference for picking the endpoint version to return. By default the API returns the LAST_UPDATED version. If you set the preference to ACTIVATED_FIRST, the API first attempts to return the version currently active on the production network. If such version doesn’t exist, the API attempts to return the version currently active on the staging network. If both of these checks fail, the API returns the last updated version.",
			Destination: &flagsListEndpoints.VersionPreference,
		},
		cli.StringFlag{
			Name:        "sortBy",
			Usage:       "The field to sort endpoints by, either the API name (corresponding to the apiEndPointName member) or updateDate.",
			Destination: &flagsListEndpoints.SortBy,
		},
		cli.StringFlag{
			Name:        "sortOrder",
			Usage:       "The sort order, either desc for descending or the default asc for ascending.",
			Destination: &flagsListEndpoints.SortOrder,
		},
		cli.StringFlag{
			Name:        "contains",
			Usage:       "The search query substring criteria matching the endpoint’s name, description, basePath, apiCategoryName, and resourcePath.",
			Destination: &flagsListEndpoints.Contains,
		},
		cli.StringFlag{
			Name:        "category",
			Usage:       "Filters endpoints by the specified apiCategoryName, including the __UNCATEGORIZED__ keyword.",
			Destination: &flagsListEndpoints.Category,
		},
		cli.IntFlag{
			Name:        "page-size",
			Usage:       "The number of endpoints on each page of results, 25 by default.",
			Destination: &flagsListEndpoints.PageSize,
		},
		cli.IntFlag{
			Name:        "page",
			Usage:       "The page number index, starting at the default value of 1.",
			Destination: &flagsListEndpoints.Page,
		},
	},
}

func callListEndpoints(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Fetching endpoints list...",
		fmt.Sprintf("Fetching endpoints list...... [%s]", color.GreenString("OK")),
	)

	list := &api.EndpointList{}
	err = list.ListEndpoints(flagsListEndpoints)

	return output(c, list.APIEndPoints, err)
}
