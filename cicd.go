package main

import (
	"cicd-lite/cmd"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "CI/CD lite"
	app.Usage = "A painless self-hosted CI/CD service"
	app.Version = "0.0.0.0-alpha"
	app.Commands = []cli.Command{
		cmd.Config,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)

	_ = app.Run(os.Args)
}
