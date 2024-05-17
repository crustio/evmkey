package main

import (
	"github.com/urfave/cli/v2"
)

var (
	VersionFlag = &cli.BoolFlag{
		Name:               "v",
		Aliases:            []string{"version"},
		Usage:              "Output the version number",
		DisableDefaultText: true,
	}

	KeystorePathFlag = &cli.StringFlag{
		Name:  "keystore",
		Usage: "keystore path",
		Value: "./keystore",
	}

	PasswordFlag = &cli.StringFlag{
		Name:  "password",
		Usage: "set password",
		Value: "", // set later
	}
)

var rootFlags = []cli.Flag{
	VersionFlag,
	KeystorePathFlag,
	PasswordFlag,
}
