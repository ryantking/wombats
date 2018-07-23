package cmd

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"

	"github.com/RyanTKing/wombats/pkg/ats"
	log "github.com/sirupsen/logrus"
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
		return fmt.Errorf("found unnexpected argument: %s", args[0])
	}

	atsVersion, err := getATSVersion()
	if err != nil {
		log.Debugf("error getting ATS version: %s", err)
		return fmt.Errorf("could not get ATS version")
	}

	gccVersion, err := getGCCVersion()
	if err != nil {
		log.Debugf("error getting gcc version: %s", err)
		return fmt.Errorf("could not get gcc version")
	}

	fmt.Printf("wombats %s\n", womVersion)
	fmt.Printf("ATS %s\n", atsVersion)
	fmt.Printf("gcc %s\n", gccVersion)

	return nil
}

func getATSVersion() (string, error) {
	out, err := ats.ExecPatsccOutput("-vats")
	if err != nil {
		return "", err
	}
	r, err := regexp.Compile("version [\\d+.?]+")
	if err != nil {
		return "", err
	}
	version := r.FindString(out)
	if version == "" {
		return "", fmt.Errorf("could not find ATS version")
	}

	return strings.Split(version, " ")[1], nil
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
