package cmd

import (
	"os"
	"path"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/RyanTKing/wombats/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	dirName = "test_dir"

	// Starting directory
	origDir string
)

type newTestSuite struct {
	suite.Suite
}

func (s *newTestSuite) SetupSuite() {
	wd, err := os.Getwd()
	require.Nil(s.T(), err)
	origDir = wd
}

func (s *newTestSuite) TearDownTest() {
	os.RemoveAll("Wombats.toml")

	// Reset flags
	name = ""
	git = false
	lib = false
	cats = false
	small = false

	// Go to starting directory
	err := os.Chdir(origDir)
	require.Nil(s.T(), err)

	os.RemoveAll(dirName)
}

func TestNewTestSuite(t *testing.T) {
	tests := new(newTestSuite)
	suite.Run(t, tests)
}

func (s *newTestSuite) TestNewUnknownArgs() {
	err := runNew("arg1", "arg2")
	assert.NotNil(s.T(), err)
}

func (s *newTestSuite) TestNewMakeDir() {
	err := runNew(dirName)
	require.Nil(s.T(), err)
	wd, err := os.Getwd()
	require.Nil(s.T(), err)
	assert.Equal(s.T(), dirName, path.Base(wd))
}

func (s *newTestSuite) TestNewMakeDirAlreadyExists() {
	err := makeProjectDir(dirName)
	require.Nil(s.T(), err)
	err = os.Chdir("../")
	require.Nil(s.T(), err)
	err = runNew(dirName)
	assert.NotNil(s.T(), err)
}

func (s *newTestSuite) TestGetProjName() {
	name, err := getProjName()
	assert.Nil(s.T(), err)
	wd, err := os.Getwd()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), path.Base(wd), name)
}

func (s *newTestSuite) TestConfigNewDir() {
	err := runNew(dirName)
	require.Nil(s.T(), err)
	var config config.Config
	_, err = toml.DecodeFile("Wombats.toml", &config)
	require.Nil(s.T(), err)
	assert.Equal(s.T(), dirName, config.Package.Name)
}

func (s *newTestSuite) TestConfigCustomName() {
	name = "test_name"
	err := runNew(dirName)
	assert.Nil(s.T(), err)

	var config config.Config
	_, err = toml.DecodeFile("Wombats.toml", &config)
	require.Nil(s.T(), err)
	assert.Equal(s.T(), "test_name", config.Package.Name)
}

func (s *newTestSuite) TestNewWithGit() {
	git = true
	err := runNew(dirName)
	require.Nil(s.T(), err)
	_, err = os.Stat(".git")
	assert.Nil(s.T(), err)
}

func (s *newTestSuite) TestNewDirs() {
	cats = true
	err := runNew(dirName)
	require.Nil(s.T(), err)
	for _, dir := range []string{"SATS", "DATS", "CATS", "BUILD"} {
		_, err = os.Stat(dir)
		assert.Nil(s.T(), err)
	}
}

func (s *newTestSuite) TestNewNoDirs() {
	small = true
	err := runNew(dirName)
	require.Nil(s.T(), err)
	for _, dir := range []string{"SATS", "DATS", "CATS", "BUILD"} {
		_, err = os.Stat(dir)
		assert.True(s.T(), os.IsNotExist(err))
	}
}

func (s *newTestSuite) TestNewLib() {
	lib = true
	err := runNew(dirName)
	require.Nil(s.T(), err)
	for _, dir := range []string{"SATS", "DATS"} {
		_, err = os.Stat(dir)
		assert.Nil(s.T(), err)
	}
	_, err = os.Stat("BUILD")
	assert.True(s.T(), os.IsNotExist(err))
}
