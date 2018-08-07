package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/RyanTKing/wombats/pkg/ats"
	"github.com/RyanTKing/wombats/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the binary to your bin folder",
	Long: `Build the project if necessary and then install it your bin so the
binary can be accessed globally.

All arguments passed to this command are passed to patscc.`,
	Run: runInstall,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func runInstall(cmd *cobra.Command, args []string) {
	config, err := config.Read()
	if err != nil {
		log.Debug(err)
		log.Fatalf(
			"could not find '%s' in this directory or any parent directory",
			"Wombats.toml",
		)
	}

	patshome := os.Getenv("PATSHOME")
	if _, err := os.Stat(patshome); err != nil {
		log.Fatal("could not find PATSHOME location")
	}
	bin := fmt.Sprintf("%s/bin", patshome)

	projName, err := getProjName()
	if err != nil {
		log.Debug(err)
		log.Fatalf("could not install '%s' project", config.Package.Name)
	}

	execFile := ats.Build(projName, config.Package.EntryPoint,
		config.Package.Clibs)
	log.Infof("installing '%s' to '%s'", filepath.Base(execFile), bin)

	wd, err := os.Getwd()
	if err != nil {
		log.Debug(err)
		log.Fatal("could not get current directory")
	}

	absExecFile := fmt.Sprintf("%s/%s", wd, execFile[2:])
	globalExecFile := fmt.Sprintf("%s/%s", bin, execFile[2:])
	if err = os.Rename(absExecFile, globalExecFile); err != nil {
		log.Debug(err)
		log.Fatalf("could not install '%s' to '%s'", absExecFile, bin)
	}
	log.Infof("Installed '%s' to '%s'", execFile, globalExecFile)
}
