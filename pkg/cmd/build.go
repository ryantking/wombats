package cmd

import (
	"fmt"

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
	fmt.Println("build called")
}
