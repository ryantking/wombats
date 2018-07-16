package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
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
			if err := runNew(args...); err != nil {
				log.Fatal(err)
			}
		},
	}

	// ErrProjectExists is thrown when a project's directory already exists
	ErrProjectExists = errors.New("Project already exists")

	// Flags
	name  string
	git   bool
	lib   bool
	cats  bool
	small bool
)

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&name, "name", "n", "",
		"The name of the project (default the directory name)")
	newCmd.Flags().BoolVar(&git, "git", false,
		"Initialize a new git repository.")
	newCmd.Flags().BoolVar(&lib, "lib", false, "Dont create a build directory")
	newCmd.Flags().BoolVar(&cats, "cats", false, "Create a CATS directory")
	newCmd.Flags().BoolVar(&small, "small", false,
		"Use a small project template (no DATS/SATS/BUILD dirs")
}

func runNew(args ...string) error {
	// Validate flags and arguments
	if len(args) > 1 {
		return fmt.Errorf("Unknown Arguments: %v", args[1:])
	}
	if cats && small {
		return fmt.Errorf("Cannot have CATS directory in small prject")
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
	config := config.New(name, small)
	if err := config.Write(); err != nil {
		return fmt.Errorf("Error creating Wombats.toml: %s", err)
	}

	// Initialize a git repo if specified
	if git {
		cmd := exec.Command("git", "init")
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("Error initializing git repo: %s", err)
		}
	}

	// If the small option is not specified, create the directories and
	// staloadall.hats file
	if !small {
		if err := createDirs(); err != nil {
			return fmt.Errorf("Error creating directories: %s", err)
		}

		if _, err := os.Create("staloadall.hats"); err != nil {
			return fmt.Errorf("Error creating staloadall.hats: %s", err)
		}
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

func createDirs() error {
	dirs := []string{"SATS", "DATS"}
	if !lib {
		dirs = append(dirs, "BUILD")
	}
	if cats {
		dirs = append(dirs, "CATS")
	}
	for _, dir := range dirs {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
