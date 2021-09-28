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
	"fmt"
	"github.com/letieu/ssht/util"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List SSH",
	Long:  `List SSH`,

	Run: func(cmd *cobra.Command, args []string) {
		var targets []util.Target = util.ListTargets()
		for _, target := range targets {
			if target.Port == "" {
				target.Port = "22"
			}
			fmt.Printf("%s: %s@%s:%s \n", target.Name, target.Username, target.Host, target.Port)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
