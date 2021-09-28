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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// jokeCmd represents the joke command
var jokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "Show dadjoke",
	Long:  `Show random joke`,
	Run: func(cmd *cobra.Command, args []string) {
        getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(jokeCmd)
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
    url := "https://icanhazdadjoke.com"
    responseBytes := getJokeData(url)
    joke := Joke{}

    if err := json.Unmarshal(responseBytes, &joke); err != nil {
        fmt.Printf("Could not Unmarshal responseBytes. %v", err)
    }

    fmt.Println(joke.Joke)
}

func getJokeData(baseApi string) []byte {
    request, err := http.NewRequest(
        http.MethodGet,
        baseApi,
        nil,
    )

    if err != nil {
        log.Printf("Could not request a dadjoke. %v", err)
    }

    request.Header.Add("Accept", "application/json")
    request.Header.Add("User-Agent", "SSHT CLI")

    response, err := http.DefaultClient.Do(request)
    if err != nil {
        log.Printf("Could not make request. %v", err)
    }

    responseBytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Printf("Could not read response body. %v", err)
    }

    return responseBytes
}
