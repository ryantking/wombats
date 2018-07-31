package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetVersions(t *testing.T) {
	womVersionStr := fmt.Sprintf("wombats %s", womVersion)
	atsVersion, err := getATSVersion()
	require.Nil(t, err)
	atsVersionStr := fmt.Sprintf("ATS %s", atsVersion)
	gccVersion, err := getGCCVersion()
	require.Nil(t, err)
	gccVersionStr := fmt.Sprintf("gcc %s", gccVersion)

	origStdout := os.Stdout
	r, w, err := os.Pipe()
	require.Nil(t, err)
	os.Stdout = w
	runVersion(nil, []string{})
	outC := make(chan string, 5)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	w.Close()
	os.Stdout = origStdout
	out := <-outC
	lines := strings.Split(strings.TrimSpace(out), "\n")
	require.Equal(t, 3, len(lines))
	assert.Equal(t, womVersionStr, lines[0])
	assert.Equal(t, atsVersionStr, lines[1])
	assert.Equal(t, gccVersionStr, lines[2])
}
