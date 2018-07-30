package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/RyanTKing/wombats/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize the current directory as a Wombats project",
		Long: `Create a Wombats project in an existing ATS directory.
It attempts to to look at the files to figure out default settings.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runInit(args...); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&name, "name", "n", "",
		"The name of the project (default the directory name)")
	initCmd.Flags().BoolVar(&git, "git", false,
		"Initialize a new git repository.")
}

func runInit(args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("found unexpected argument '%s'", args[0])
	}

	projName, err := getProjName()
	if err != nil {
		log.Debugf("get current dir error: %s", err)
		return fmt.Errorf("could not get current directory")
	}

	if name == "" {
		name = projName
	}

	// Get the initial config and write it to a file
	config := config.New(name, projName, small)
	config.Package.EntryPoint = findEntryPoint(config)
	if err := config.Write(); err != nil {
		log.Debugf("config write error: %s", err)
		return fmt.Errorf("could not create 'Wombats.toml' file")
	}

	if err := initGitRepo(); err != nil {
		return err
	}

	log.Infof("Initialized '%s' project in current directory", name)
	return nil
}

func findEntryPoint(config *config.Config) string {
	entryPoint := config.Package.EntryPoint
	reader := bufio.NewReader(os.Stdin)
	_, err := os.Stat(entryPoint)
	for err != nil {
		log.Warnf("file: '%s' does not exist\n", entryPoint)
		fmt.Print("entry point: ")
		entryPoint, err = reader.ReadString('\n')
		entryPoint = strings.TrimSpace(entryPoint)
		if err != nil {
			log.Fatal("could not read input")
		}
		_, err = os.Stat(entryPoint)
	}

	return entryPoint
}
