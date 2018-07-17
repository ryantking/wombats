package builder

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// New creates a new builder
func New(name, entryPoint string, small bool, patscc string) *Builder {
	execFile := fmt.Sprintf("./BUILD/%s", name)
	if small {
		execFile = fmt.Sprintf("./%s", name)
	}

	return &Builder{
		ProjName:   name,
		ExecFile:   execFile,
		EntryPoint: entryPoint,
		Patscc:     patscc,
	}
}

func (b *Builder) cleanup() {
	os.RemoveAll("./**/*.{d,s}ats_c")
}

// Build compiles the project
func (b *Builder) Build() error {
	defer b.cleanup()

	args := fmt.Sprintf("-o %s %s", b.ExecFile, b.EntryPoint)
	cmd := exec.Command(b.Patscc, args)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Debugf("error running build command: %s", err)
		return err
	}

	return nil
}