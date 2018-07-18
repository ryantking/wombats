package builder

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

// New creates a new builder
func New(name, entryPoint string, small bool) *Builder {
	execFile := fmt.Sprintf("./BUILD/%s", name)
	if small {
		execFile = fmt.Sprintf("./%s", name)
	}

	return &Builder{
		ProjName:   name,
		ExecFile:   execFile,
		EntryPoint: entryPoint,
	}
}

func hasPatscc() bool {
	for _, loc := range strings.Split(os.Getenv("PATH"), ":") {
		if _, err := os.Stat(loc + "/patscc"); err == nil {
			return true
		}
	}

	return false
}

// Build compiles the project
func (b *Builder) Build() error {
	if !hasPatscc() {
		log.Errorf("could not find patscc executable in PATH")
		return fmt.Errorf("could not find patscc executable in PATH")
	}

	cmd := exec.Command("patscc", "-o", b.ExecFile, b.EntryPoint)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Debugf("error running build command: %s", err)
		return err
	}

	return nil
}
