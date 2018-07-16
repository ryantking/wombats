package cmd

import (
	"fmt"

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
	rootCmd.AddCommand(buildCmd)
}

func runRun(cmd *cobra.Command, args []string) {
	fmt.Println("run halled")
}
