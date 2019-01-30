package cmd

import (
	"cicd-lite/config"
	"fmt"
	"github.com/urfave/cli"
	"log"
)

var Config = cli.Command{
	Name:        "config",
	Usage:       "Show a configuration for your project.",
	Description: "Show current configuration you are working.",
	Action:      runConfig,
}

func runConfig(ctx *cli.Context) error {
	project, err := config.GetProject()

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	source, err := config.GetSource(&project)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("--- Current configuration:\n%s\n\n", string(source))

	return nil
}
