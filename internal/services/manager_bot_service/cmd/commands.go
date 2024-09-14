package cmd

import "github.com/urfave/cli"

var (
	serveCMD = &cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "Serving telegram bot",
	}
)
