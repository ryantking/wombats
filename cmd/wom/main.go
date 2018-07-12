package main

import (
	"fmt"
	"os"

	"github.com/RyanTKing/wombats/pkg/subcommands"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:    "Wombats",
		Usage:   "Wombats is a tool for managing your ATS projects.",
		Version: "0.1",
		Commands: []cli.Command{
			subcommands.GetNewCommand(),
			{
				Name:  "run",
				Usage: "Run the compiled binary.",
				Action: func(c *cli.Context) error {
					fmt.Println("Run!")
					return nil
				},
			},
			{
				Name:   "build",
				Usage:  "Compile the project to a binary.",
				Action: subcommands.Build,
			},
			{
				Name:  "install",
				Usage: "Compile to a binary and install it.",
				Action: func(c *cli.Context) error {
					fmt.Println("Install!")
					return nil
				},
			},
			{
				Name:  "fetch",
				Usage: "Fetch all project dependencies.",
				Action: func(c *cli.Context) error {
					fmt.Println("Fetch!")
					return nil
				},
			},
			{
				Name:  "version",
				Usage: "Print out ATS version information.",
				Action: func(c *cli.Context) error {
					fmt.Println("version!")
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
