// Copyright © 2018 Alex Kolbasov
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

package cmd

import (
	"log"
	"strings"

	"github.com/akolb1/gometastore/hmsclient"
	"github.com/spf13/cobra"
)

// TODO support database description

// dbCmd represents the db command
var dbCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create database",
	Long: `Create database with given parameters.
Parameters are specified as a list of name=value pairs.
`,
	Run: createDB,
}

// argsToParams converts list of x=y strings into a parameters map
func argsToParams(args []string) map[string]string {
	if len(args) == 0 {
		return nil
	}
	params := make(map[string]string)
	for _, a := range args {
		parts := strings.Split(a, "=")
		// Ignore any arguments that do not look like parameters
		if len(parts) != 2 {
			continue
		}
		params[parts[0]] = parts[1]
	}
	return params
}

// createDB creates database
func createDB(cmd *cobra.Command, args []string) {
	dbName, _ := cmd.Flags().GetString(optDbName)
	owner := getOwner()
	params := argsToParams(args)
	client, err := getClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// Check whether database already exists
	if databases, err := client.GetAllDatabases(); err != nil {
		log.Fatal(err)
	} else {
		for _, name := range databases {
			if name == dbName {
				log.Fatalf("database %s already exists\n", dbName)
			}
		}
	}

	err = client.CreateDatabase(&hmsclient.Database{
		Name:       dbName,
		Owner:      owner,
		Parameters: params,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	dbCmd.AddCommand(dbCreateCmd)
}
