package cli

import (
	"os"

	"github.com/urfave/cli"
)

func Run() error {
	app := cli.App{
		Name:      "manager-bot",
		Usage:     "A bot for interaction between the manager and the parcel management system",
		UsageText: "manager-bot [command] serve (s) [command option] debug (d)",
		Version:   "0.0.1",
		Commands: []*cli.Command{
			serveCMD,
		},
	}

	return app.Run(os.Args)
}
