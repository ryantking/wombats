package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/RyanTKing/wombats/pkg/builder"
	"github.com/RyanTKing/wombats/pkg/config"
	log "github.com/sirupsen/logrus"
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
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, args []string) {
	config, err := config.Read()
	if err != nil {
		log.Debugf("error reading config: %s", err)
		log.Fatalf(
			"could not find '%s' in this directory or any parent directory",
			"Wombats.toml",
		)
	}

	projName, err := getProjName()
	if err != nil {
		log.Fatalf("could not build '%s' project", config.Package.Name)
	}

	executable := fmt.Sprintf("./BUILD/%s", projName)
	if config.Package.Small {
		executable = fmt.Sprintf("./%s", projName)
	}

	b := builder.New(projName, config.Package.EntryPoint, config.Package.Small,
		patscc)
	if _, err := os.Stat(executable); os.IsNotExist(err) {
		if err := b.Build(); err != nil {
			log.Debugf("build error: %s", err)
			log.Fatalf("could not build '%s' project", config.Package.Name)
		}

		log.Infof("compiled '%s' project", config.Package.Name)
	}

	origStdout := os.Stdout
	r, w, err := os.Pipe()
	os.Stdout = w
	execCmd := exec.Command(b.ExecFile, strings.Join(args, " "))
	execCmd.Start()
	outC := make(chan string)
	go func() {
		for out := range outC {
			fmt.Print(out)
		}
	}()
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	execCmd.Wait()
	w.Close()
	os.Stdout = origStdout

	log.Infof("successfully ran '%s' project", b.ExecFile)
}
