package cmd

import (
	"fmt"

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
	fmt.Println("run called")
}
