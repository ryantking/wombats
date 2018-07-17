package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

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

	projName, err := getProjName()
	if err != nil {
		log.Fatalf("could not build '%s' project", config.Package.Name)
	}

	executable := fmt.Sprintf("./BUILD/%s", projName)
	if config.Package.Small {
		executable = fmt.Sprintf("./%s", projName)
	}

	b := builder.New(projName, config.Package.EntryPoint, config.Package.Small,
		patscc)
	if _, err := os.Stat(executable); os.IsNotExist(err) {
		log.Infof("Building '%s' project")
		start := time.Now()
		if err := b.Build(); err != nil {
			log.Debugf("build error: %s", err)
			log.Fatalf("could not build", config.Package.Name)
		}

		dur := time.Since(start)
		unit := "s"
		t := dur.Seconds()
		if t >= 60 && t < 3600 {
			unit = "m"
			t = dur.Minutes()
		} else if t >= 36000 {
			unit = "h"
			t = dur.Hours()
		}
		log.Infof("Finished building in %.2f%s", t, unit)
	}

	log.Infof("Running '%s'", b.ExecFile)
	execCmd := exec.Command(b.ExecFile, args...)
	execCmd.Env = os.Environ()
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	if err := execCmd.Run(); err != nil {
		log.Debugf("error running '%s': %s", config.Package.Name, err)
		log.Fatalf("could not run '%s' project", config.Package.Name)
	}

}
