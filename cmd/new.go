/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/letieu/ssht/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var username string
var host string
var name string
var port string

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new SSH",
	Long:  `Create new SSH by password, user and host`,

	Run: func(cmd *cobra.Command, args []string) {
		var targets []util.Target = util.ListTargets()
		newTarget, err := util.CreateTarget(name, host, username, port)
		if err != nil {
			return
		}


		targets = append(targets, newTarget)

		jsonString, _ := json.Marshal(targets)
        home := os.Getenv("HOME")
		ioutil.WriteFile(home + "/.ssht/targets.json", jsonString, 0644)
		println("Added")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// username flag
	newCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	viper.BindPFlag("username", newCmd.Flags().Lookup("username"))
	viper.SetDefault("username", "root")

	// host flag
	newCmd.Flags().StringVarP(&host, "host", "a", "", "Host")
	viper.BindPFlag("host", newCmd.Flags().Lookup("host"))

	// host flag
	newCmd.Flags().StringVarP(&port, "port", "p", "22", "Port")
	viper.BindPFlag("p", newCmd.Flags().Lookup("port"))
	viper.SetDefault("p", "22")

	// name flag
	newCmd.Flags().StringVarP(&name, "name", "n", "", "Name")
	viper.BindPFlag("name", newCmd.Flags().Lookup("name"))
}
