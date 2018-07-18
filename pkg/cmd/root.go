package cmd

import (
	"fmt"
	"os"

	"github.com/RyanTKing/wombats/pkg/logging"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
			viper.BindPFlag("verbose", cmd.Flags().Lookup("verbose"))

			if viper.GetBool("verbose") {
				log.SetLevel(log.DebugLevel)
			} else {
				log.SetLevel(log.InfoLevel)
			}
		},
	}

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
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	// Setup Logging
	logFile := fmt.Sprintf("%s/.wombats.log", os.Getenv("HOME"))
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening %s for logging: %s", logFile, err)
	} else {
		log.SetOutput(f)
	}

	log.SetFormatter(new(log.TextFormatter))
	log.AddHook(logging.NewLogrusHook())
}
