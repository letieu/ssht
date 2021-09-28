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
	"os"
	"os/exec"

	"github.com/letieu/ssht/util"
	"github.com/spf13/cobra"
)

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Go to new SSH",
	Long:  `Go to ssh`,

	Run: func(cmd *cobra.Command, args []string) {
        id := args[0]
		var targets []util.Target = util.ListTargets()
		for _, target := range targets {
			if target.Name == id {
                ssh(target)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
}

func ssh(target util.Target) {
    app := "ssh"

    if target.Port == "" {
        target.Port = "80"
    }
    arg0 := "-p"
    arg1 := target.Port
    arg2 := target.Username + "@" + target.Host

    cmd := exec.Command(app, arg0, arg1, arg2)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout

    err := cmd.Run()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
}
