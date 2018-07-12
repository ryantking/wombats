package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Download all project dependencies",
	Long: `All project dependencies are downloaded if they do not already
exist. Version resolution also takes place if the downloaded version does not
match the version specified in the package config.

Specific dependencies can be fetched by giving a name or list of names.
For example:

	$ wom fetch [dependency]`,
	Run: runFetch,
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}

func runFetch(cmd *cobra.Command, args []string) {
	fmt.Println("fetch called")
}
