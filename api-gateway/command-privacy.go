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
	"strings"

	api "github.com/akamai/AkamaiOPEN-edgegrid-golang/api-endpoints-v2"
	akamai "github.com/akamai/cli-common-golang"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var commandPrivacy cli.Command = cli.Command{
	Name:        "privacy-add",
	ArgsUsage:   "",
	Description: "Make an endpoint|resource|method public or private",
	HideHelp:    true,
	Action:      callPrivacy,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "endpoint",
			Usage: "The unique identifier for the endpoint.",
		},
		cli.StringFlag{
			Name:  "version",
			Usage: "The endpoint version number.",
		},
		cli.BoolFlag{
			Name:  "public",
			Usage: "Make this endpoint public.",
		},
		cli.BoolFlag{
			Name:  "private",
			Usage: "Make this endpoint private.",
		},
		cli.StringFlag{
			Name:  "resource",
			Usage: "The resource name to apply the settings to.",
		},

		cli.BoolFlag{
			Name:  "get",
			Usage: "Apply to GET requests.",
		},
		cli.BoolFlag{
			Name:  "post",
			Usage: "Apply to POST requests.",
		},
		cli.BoolFlag{
			Name:  "put",
			Usage: "Apply to PUT requests.",
		},
		cli.BoolFlag{
			Name:  "delete",
			Usage: "Apply to DELETE requests.",
		},
		cli.BoolFlag{
			Name:  "patch",
			Usage: "Apply to PATCH requests.",
		},
		cli.BoolFlag{
			Name:  "head",
			Usage: "Apply to HEAD requests.",
		},
		cli.BoolFlag{
			Name:  "options",
			Usage: "Apply to OPTIONS requests.",
		},
	},
}

func callPrivacy(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Updating privacy...",
		fmt.Sprintf("Updating privacy...... [%s]", color.GreenString("OK")),
	)

	version := c.Int("version")
	if version == 0 {
		versions, err := api.ListVersions(&api.ListVersionsOptions{EndpointId: c.Int("endpoint")})
		if err != nil {
			return err
		}

		loc := len(versions.APIVersions) - 1
		v := versions.APIVersions[loc]
		version = v.VersionNumber
	}

	settings, err := api.GetAPIPrivacySettings(c.Int("endpoint"), version)
	if err != nil {
		return output(c, nil, err)
	}

	if c.String("resource") != "" {

		for i, resource := range settings.Resources {
			if c.String("resource") == resource.Path {
				methods := getMethodsPassed(c)
				if c.Bool("public") {
					methods = merge(methods, resource.Methods)
				} else {
					methods = remove(methods, resource.Methods)
				}
				resource.Methods = methods
				resource.Public = c.Bool("public")
				settings.Resources[i] = resource
			}
		}
	} else {
		settings.Public = c.Bool("public")
	}

	_, err = api.UpdateAPIPrivacySettings(c.Int("endpoint"), version, settings)

	return output(c, settings, err)
}

func getMethodsPassed(c *cli.Context) []string {
	allMethods := []string{
		"HEAD",
		"DELETE",
		"POST",
		"GET",
		"OPTIONS",
		"PUT",
		"PATCH",
	}
	methods := []string{}

	for _, m := range allMethods {
		if c.Bool(strings.ToLower(m)) {
			methods = append(methods, m)
		}
	}

	return methods
}

func merge(s1, s2 []string) []string {
	input := append(s1, s2...)
	merged := make([]string, 0, len(input))
	seen := make(map[string]bool)

	for _, val := range input {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			merged = append(merged, val)
		}
	}

	return merged
}

func remove(s1, s2 []string) []string {
	final := []string{}
	for i, v2 := range s2 {
		for _, v1 := range s1 {
			if v1 == v2 {
				final = append(s1[:i], s2[i+1:]...)
			}
		}
	}
	return final
}
