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
		Run: runInit,
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("name", "n", "",
		"The name of the project (default the directory name)")
	initCmd.Flags().Bool("git", false, "Initialize a new git repository.")
}

func runInit(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		log.Fatalf("found unexpected argument '%s'", args[0])
	}

	projName, err := getProjName()
	if err != nil {
		log.Debug(err)
		log.Fatal("could not get current directory")
	}

	// Get the project name
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		log.Debug(err)
		log.Fatal("could not check command flag")
	}
	if name == "" {
		name = projName
	}

	// Get the initial config and write it to a file
	small := isSmall()
	config := config.New(name, projName, small)
	config.Package.EntryPoint = findEntryPoint(config)
	if err := config.Write(); err != nil {
		log.Debug(err)
		log.Fatal("could not create 'Wombats.toml' file")
	}

	// Iniitlize the git repo is specified
	git, err := cmd.Flags().GetBool("git")
	if err != nil {
		log.Debug(err)
		log.Fatal("could not check command flag")
	}
	if git {
		if err := initGitRepo(); err != nil {
			log.Fatal(err)
		}
	}

	log.Infof("Initialized '%s' project in current directory", name)
}

func isSmall() bool {
	dirs := []string{"DATS", "SATS", "CATS", "BUILD"}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); err == nil {
			return false
		}
	}

	return true
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
			log.Debug(err)
			log.Fatal("could not read input")
		}
		_, err = os.Stat(entryPoint)
	}

	return entryPoint
}
