package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestATSBinary(t *testing.T) {
	require.NotEmpty(t, patshome)
	info, err := os.Stat(patscc)
	require.Nil(t, err)
	assert.True(t, info.Mode().IsRegular())
}
