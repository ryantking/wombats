package builder

import (
	"fmt"
	"os"
	"os/exec"
)

// New creates a new builder
func New(name string, small bool, patscc string) *Builder {
	datsDir := "./DATS/*.dats"
	execFile := fmt.Sprintf("./BUILD/%s", name)
	if small {
		datsDir = "./*.dats"
		execFile = fmt.Sprintf("./%s", name)
	}

	return &Builder{
		ProjName: name,
		DATSDir:  datsDir,
		ExecFile: execFile,
		Patscc:   patscc,
	}
}

func (b *Builder) cleanup() {
	os.RemoveAll("./**/*.{d,s}ats_c")
}

// Build compiles the project
func (b *Builder) Build() error {
	defer b.cleanup()
	cmd := exec.Command(b.Patscc, "-o", b.ExecFile, b.DATSDir)
	return cmd.Run()
}
