package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/RyanTKing/wombats/pkg/config"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var (
	newCmd = &cobra.Command{
		Use:   "new",
		Short: "Create a new Wombats project",
		Long: `Create a new Wombats project in the current directory or in a
specified directory if a name is provided. For example:
	$ wom new     # Initializes a project in the current directory
	$ wom new foo # Creates the directory foo and initializes a project in it`,
		Run: func(cmd *cobra.Command, args []string) {
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				log.Fatalf("Error getting flag value")
			}

			err = runNew(name, args...)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	// ErrProjectExists is thrown when a project's directory already exists
	ErrProjectExists = errors.New("Project already exists")
)

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringP("name", "n", "",
		"The name of the project (default the directory name)")
}

func runNew(name string, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("Unknown Arguments: %v", args[1:])
	}

	// If a directory is provided, make it and change to it
	if len(args) > 0 {
		if err := makeProjectDir(args[0]); err != nil {
			return fmt.Errorf("Error creating project: %s", err)
		}
	}

	// Assume the name is the current directory if not set
	if name == "" {
		var err error
		name, err = getProjName()
		if err != nil {
			return fmt.Errorf("Error getting project name: %s", err)
		}
	}

	// Get the initial config and write it to a file.
	config := config.New(name)
	if err := config.Write(); err != nil {
		return fmt.Errorf("Error creating Wombats.toml: %s", err)
	}

	return nil
}

func makeProjectDir(name string) error {
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		return ErrProjectExists
	}

	if err := os.Mkdir(name, os.ModePerm); err != nil {
		return err
	}

	return os.Chdir(name)
}

func getProjName() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Base(wd), nil
}
