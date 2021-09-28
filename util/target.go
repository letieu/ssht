package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type Target struct {
	Name     string `json:""`
	Username string `json:""`
	Host     string `json:""`
	Port     string `json:""`
}

func CreateTarget(name string, host string, username string, port string) (Target, error) {
	if port == "" {
		port = "22"
	}
	err := pair(port, username, host)

	return Target{
		Username: username,
		Host:     host,
		Port:     port,
		Name:     name,
	}, err
}

func ListTargets() []Target {
	var targets []Target = make([]Target, 0)
    home := os.Getenv("HOME")

	jsonFile, err := os.Open(home + "/.ssht/targets.json")
	if err != nil {
		return targets
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &targets)

	return targets
}

func pair(port string, username string, host string) error {
	app := "ssh-copy-id"

	arg0 := "-p"
	arg1 := port
	arg2 := username + "@" + host

	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(stdout))
	return nil
}
