package cmd

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dirName = "test_dir"

func tearDown() {
	os.Remove(dirName)
}

func TestMakeProjectDir(t *testing.T) {
	err := makeProjectDir(dirName)
	assert.Nil(t, err)
	defer tearDown()
	err = os.Chdir("../")
	assert.Nil(t, err)
}

func TestMakeProjectDirAlreadyExists(t *testing.T) {
	err := makeProjectDir(dirName)
	assert.Nil(t, err)
	defer tearDown()
	err = os.Chdir("../")
	assert.Nil(t, err)
	err = makeProjectDir(dirName)
	assert.Equal(t, ErrProjectExists, err)
}

func TestGetProjName(t *testing.T) {
	name, err := getProjName("test_name")
	assert.Nil(t, err)
	assert.Equal(t, "test_name", name)
}

func TestGetProjNameDir(t *testing.T) {
	wd, err := os.Getwd()
	assert.Nil(t, err)
	name, err := getProjName("")
	assert.Nil(t, err)
	assert.Equal(t, path.Base(wd), name)
}

func TestRunNew(t *testing.T) {
	err := runNew("")
	assert.Nil(t, err)
}

func TestRunNewCustomDir(t *testing.T) {
	defer tearDown()
	err := runNew("", dirName)
	assert.Nil(t, err)
	err = os.Chdir("../")
	assert.Nil(t, err)
}
