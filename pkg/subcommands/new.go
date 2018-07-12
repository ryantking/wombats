package subcommands

import (
	"log"
	"os"
	"path"

	"github.com/RyanTKing/wombats/pkg/config"
	"github.com/urfave/cli"
)

var (
	name string
)

// GetNewCommand returns the cli.Command struct for the "wom new" command.
func GetNewCommand() cli.Command {
	return cli.Command{
		Name:   "new",
		Usage:  "Create a new ATS project.",
		Action: runNew,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "name, n",
				Usage:       "Name of the project",
				Destination: &name,
			},
		},
	}
}

func makeProjectDir(name string) error {
	err := os.Mkdir(name, os.ModePerm)
	if os.IsExist(err) {
		return &ErrProjectExists{name}
	} else if err != nil {
		return err
	}

	return os.Chdir(name)
}

func getProjName() (string, error) {
	if name != "" {
		return name, nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Base(wd), nil
}

func runNew(c *cli.Context) error {
	if len(c.Args()) > 1 {
		return &ErrUnknownArguments{c.Args()[1:]}
	}

	// If a directory is provided, make it and change to it
	if len(c.Args()) > 0 {
		if err := makeProjectDir(c.Args()[0]); err != nil {
			log.Fatalf("Error creating project: %s", err)
		}
	}

	// Assume the name is the current directory if not set
	projName, err := getProjName()
	if err != nil {
		log.Fatalf("Error getting project name: %s", err)
	}

	// Get the initial config and write it to a file.
	config := config.New(projName)
	if err := config.Write(); err != nil {
		log.Fatalf("Error creating Wombats.toml: %s", err)
	}

	return nil
}
