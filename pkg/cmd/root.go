package cmd

import (
	"fmt"
	"os"

	"github.com/RyanTKing/wombats/pkg/logging"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "wom",
		Short: "Wombats is a tool for managing ATS projects",
		Long: `Wombats is a tool for building, running, and install ATS as well as
managing dependencies.
	
A project can be initialized and ran in the directory of an existing ATS
project as follows:
	$ wom new
	$ wom run`,
	}

	verbose  bool
	patshome string
	patscc   string
)

// Execute is the entry point for wombats
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false,
		"verbose output")

	// Setup Logging
	logFile := fmt.Sprintf("%s/.wombats.log", os.Getenv("HOME"))
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening %s for logging: %s", logFile, err)
	} else {
		log.SetOutput(f)
	}
	log.AddHook(logging.NewLogrusHook())

	patshome = os.Getenv("PATSHOME")
	patscc = patshome + "/bin/patscc"
}
