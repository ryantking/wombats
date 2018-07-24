package cmd

import (
	"strings"

	"github.com/RyanTKing/wombats/pkg/ats"
	"github.com/RyanTKing/wombats/pkg/config"
	"github.com/RyanTKing/wombats/pkg/logging"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ATSError holds info about a line error
type ATSError struct {
	fname                      string
	startL, endL, startO, endO int
	errors                     []string
}

// checkCmd represents the build command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Typecheck the current project",
	Long: `Use patscc to typecheck the project, but not compile it and report
any errors.

All arguments passed to the command will be added as arguments to the patscc
compiler command.`,
	Run: runCheck,
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func runCheck(cmd *cobra.Command, args []string) {
	config, err := config.Read()
	if err != nil {
		log.Debugf("error reading config: %s", err)
		log.Fatalf(
			"could not find '%s' in this directory or any parent directory",
			"Wombats.toml",
		)
	}

	log.Infof("Typechecking '%s' project", config.Package.Name)
	output, err := ats.ExecPatsccOutput("-tcats", "./**/*.dats", "./**/*.sats")
	if err == nil {
		log.Info("Found no typechecking errors")
		return
	}

	logging.CheckErrors(strings.TrimSpace(output))
}
