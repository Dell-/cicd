package cmd

import (
	"cicd-lite/context"
	"github.com/urfave/cli"
)

var Run = cli.Command{
	Name:        "run",
	Usage:       "Running subtask of project.",
	Description: "Running subtask of project.",
	Action:      runRun,
}

func runRun(ctx *cli.Context) error {
	context.Initialize()

	return nil
}
