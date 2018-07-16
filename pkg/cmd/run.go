package cmd

import (
	"fmt"
	"os"
	"syscall"

	"github.com/RyanTKing/wombats/pkg/builder"
	"github.com/RyanTKing/wombats/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the current project",
	Long: `Compile the project if necessary then run it if successfully built.
	
All arguments passed to the run command will be passed to patscc.`,
	Run: runRun,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, args []string) {
	config, err := config.Read()
	if err != nil {
		log.Debugf("error reading config: %s", err)
		log.Fatalf(
			"could not find '%s' in this directory or any parent directory",
			"Wombats.toml",
		)
	}

	executable := fmt.Sprintf("./BUILD/%s", config.Package.Name)
	if config.Package.Small {
		executable = fmt.Sprintf("./%s", config.Package.Name)
	}

	b := builder.New(config.Package.Name, config.Package.Small, patscc)
	if _, err := os.Stat(executable); os.IsNotExist(err) {
		if err := b.Build(); err != nil {
			log.Debugf("build error: %s", err)
			log.Fatalf("could not build '%s' project", config.Package.Name)
		}

		log.Infof("compiled '%s' project", config.Package.Name)
	}

	env := os.Environ()
	if err := syscall.Exec(b.ExecFile, args, env); err != nil {
		log.Fatalf("encountered error while running '%s' project", b.ExecFile)
	}
	log.Infof("successfully ran '%s' project", b.ExecFile)
}
