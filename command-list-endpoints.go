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
	// "os"

	api "github.com/johannac/AkamaiOPEN-edgegrid-golang/api-endpoints-v2"
	// api "github.com/akamai/AkamaiOPEN-edgegrid-golang/configdns-v1"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var commandListEndpoints cli.Command = cli.Command{
	Name:        "list-endpoints",
	ArgsUsage:   "<record type> <hostname>",
	Description: "Retrieve a list of APIs running through the Gateway; may be used as input to other commands",
	HideHelp:    true,
	Action:      callListEndpoints,
	Subcommands: cli.Commands{},
}

func callListEndpoints(c *cli.Context) error {
	fmt.Println(color.YellowString("Hello World"))

	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	list := api.EndpointList{}
	err = list.GetEndpointsList(api.ListEndpointOptions{
		ContractId: "ctr_C-1FRYVV3",
		GroupId:    "grp_68817",
		Page:       5,
		PageSize:   50,
	})
	// zone, err := api.GetZone("akamaideveloper.net")
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}
	// fmt.Println(color.BlueString(fmt.Sprintf("%s", zone)))
	fmt.Println(color.BlueString("Goodbye World"))
	return nil
}
