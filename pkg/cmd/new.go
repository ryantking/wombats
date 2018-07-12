package cmd

import (
	"errors"
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
		Run: runNew,
	}

	// ErrProjectExists is thrown when a project's directory already exists
	ErrProjectExists = errors.New("Project already exists")

	name string
)

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&name, "name", "n", "[project name]",
		"The name of the project (default the directory name)")
}

func runNew(cmd *cobra.Command, args []string) {
	if len(args) > 1 {
		log.Fatalf("Unknown Arguments %v", args[1:])
	}

	// If a directory is provided, make it and change to it
	if len(args) > 0 {
		if err := makeProjectDir(args[0]); err != nil {
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
}

func makeProjectDir(name string) error {
	err := os.Mkdir(name, os.ModePerm)
	if os.IsExist(err) {
		return ErrProjectExists
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
