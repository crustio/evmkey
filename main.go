package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Usage: "evm account manager",
		Flags: rootFlags,
		Action: func(c *cli.Context) error {
			if c.Bool("v") { // handle root flag `-v``
				return versionCmd.Action(c)
			}

			cli.ShowAppHelp(c)
			return nil
		},
		Commands: []*cli.Command{
			versionCmd,
			accountCmd,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
