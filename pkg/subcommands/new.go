package subcommands

import (
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
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

func runNew(c *cli.Context) error {
	if len(c.Args()) > 1 {
		return &ErrUnknownArguments{c.Args()[1:]}
	}

	// If a directory is provided, make it and change to it
	if len(c.Args()) > 0 {
		name = c.Args()[0]
		os.Mkdir(name, os.ModePerm)
		os.Chdir(name)
	}

	// Assume the name is the current directory if not set
	if name == "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		name = path.Base(wd)
	}

	// Get the initial config and write it to a file.
	config := config.New(name)
	f, err := os.Create("Wombats.toml")
	if os.IsExist(err) {
		log.Fatalln("Wombats.toml already exists")
		return err
	} else if err != nil {
		log.Fatalf("Error creating Wombats.toml: %s", err)
		return err
	}
	defer f.Close()
	err = toml.NewEncoder(f).Encode(config)
	if err != nil {
		return err
	}

	return nil
}
