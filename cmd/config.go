package cmd

import (
	"cicd-lite/config"
	"fmt"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Config = cli.Command{
	Name:        "config",
	Usage:       "Show a configuration for your project.",
	Description: "Show current configuration you are working.",
	Action:      runConfig,
}

func commandInitialize() {
	fmt.Println("Run 'config' command")
}

func runConfig(ctx *cli.Context) error {
	commandInitialize()

	data, err := ioutil.ReadFile("/home/dell/project/go/src/cicd-lite/etc/project.yml")

	if err != nil {
		panic(err)
	}

	project := config.Project{}

	err = yaml.Unmarshal([]byte(data), &project)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	d, err := yaml.Marshal(&project)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("--- Current configuration:\n%s\n\n", string(d))

	return nil
}
