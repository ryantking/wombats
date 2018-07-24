package cmd

import (
	"github.com/RyanTKing/wombats/pkg/ats"
	"github.com/RyanTKing/wombats/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the current project",
	Long: `If the current project is an executable then it is compiled in the
BUILD directory.

All arguments passed to the command will be added as arguments to the patscc
compiler command.`,
	Run: runBuild,
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func runBuild(cmd *cobra.Command, args []string) {
	config, err := config.Read()
	if err != nil {
		log.Debugf("error reading config: %s", err)
		log.Fatalf(
			"could not find '%s' in this directory or any parent directory",
			"Wombats.toml",
		)
	}

	projName, err := getProjName()
	if err != nil {
		log.Fatalf("could not build '%s' project", config.Package.Name)
	}

	ats.Build(projName, config.Package.EntryPoint, config.Package.Clibs)
}
