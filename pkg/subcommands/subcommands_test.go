package subcommands

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestATSBinary(t *testing.T) {
	assert.NotEmpty(t, patshome)
	info, err := os.Stat(patscc)
	assert.Nil(t, err)
	assert.True(t, info.Mode().IsRegular())
}
