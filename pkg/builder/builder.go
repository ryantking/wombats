package builder

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
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
	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	execCmd := exec.Command(b.Patscc, args)
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

	return nil
}
