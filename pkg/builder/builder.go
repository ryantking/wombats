package builder

import (
	"fmt"
	"strings"

	"github.com/RyanTKing/wombats/pkg/ats"
)

// New creates a new builder
func New(name, entryPoint string) *Builder {
	execFile := fmt.Sprintf("./BUILD/%s", name)
	if !strings.Contains(entryPoint, "./DATS/") {
		execFile = fmt.Sprintf("./%s", name)
	}

	return &Builder{
		ProjName:   name,
		ExecFile:   execFile,
		EntryPoint: entryPoint,
	}
}

// Build compiles the project
func (b *Builder) Build() error {
	if err := ats.ExecPatscc("-o", b.ExecFile, b.EntryPoint); err != nil {
		return err
	}

	return nil
}
