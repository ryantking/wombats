package cmd

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/ngaut/log"
	"github.com/spf13/cobra"
)

const womVersion = "v0.1-beta"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information about the various ATS binaries",
	Long:  `Show version information for Wombats, ATS, and gcc.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runVersion(args...); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("Unknown Arguments: %v", args)
	}

	atsVersion, err := getATSVersion()
	if err != nil {
		return fmt.Errorf("Error getting ATS version: %s", err)
	}

	gccVersion, err := getGCCVersion()
	if err != nil {
		return fmt.Errorf("Error getting gcc version: %s", err)
	}

	fmt.Printf("wombats %s\n", womVersion)
	fmt.Printf("ATS %s\n", atsVersion)
	fmt.Printf("gcc %s\n", gccVersion)

	return nil
}

func getATSVersion() (string, error) {
	versionPath := fmt.Sprintf("%s/VERSION", patshome)
	version, err := ioutil.ReadFile(versionPath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(version)), nil
}

func getGCCVersion() (string, error) {
	cmd := exec.Command("gcc", "-dumpversion")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	if err := cmd.Start(); err != nil {
		return "", err
	}
	version, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(version)), nil
}
