package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var versionCmd = &cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "Output the version number",
	Action: func(cCtx *cli.Context) error {
		fmt.Println("evmkey - v1.0.4")
		return nil
	},
}
