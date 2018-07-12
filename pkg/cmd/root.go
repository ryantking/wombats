package cmd

import (
	"fmt"
	"os"

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
	patshome = os.Getenv("PATSHOME")
	patscc = patshome + "/bin/patscc"
}
