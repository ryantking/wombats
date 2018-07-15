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

var dirName = "test_dir"

type newTestSuite struct {
	suite.Suite
}

func (s *newTestSuite) TearDownTest() {
	os.RemoveAll(dirName)
	os.RemoveAll("Wombats.toml")
}

func TestNewTestSuite(t *testing.T) {
	tests := new(newTestSuite)
	suite.Run(t, tests)
}

func (s *newTestSuite) TestNewUnknownArgs() {
	err := runNew("name", "arg1", "arg2")
	assert.NotNil(s.T(), err)
}

func (s *newTestSuite) TestNewMakeDir() {
	err := runNew("name", dirName)
	require.Nil(s.T(), err)
	err = os.Chdir("../")
	require.Nil(s.T(), err)
	_, err = os.Stat(dirName)
	assert.Nil(s.T(), err)
}

func (s *newTestSuite) TestNewMakeDirAlreadyExists() {
	err := makeProjectDir(dirName)
	require.Nil(s.T(), err)
	err = os.Chdir("../")
	require.Nil(s.T(), err)

	err = runNew("name", dirName)
	require.NotNil(s.T(), err)
}

func (s *newTestSuite) TestGetProjName() {
	name, err := getProjName()
	require.Nil(s.T(), err)
	wd, err := os.Getwd()
	require.Nil(s.T(), err)
	assert.Equal(s.T(), path.Base(wd), name)
}

func (s *newTestSuite) TestConfigNewDir() {
	err := runNew("", dirName)
	require.Nil(s.T(), err)

	var config config.Config
	_, err = toml.DecodeFile("Wombats.toml", &config)
	require.Nil(s.T(), err)
	assert.Equal(s.T(), dirName, config.Package.Name)
	err = os.Chdir("../")
	require.Nil(s.T(), err)
}

func (s *newTestSuite) TestConfigCustomName() {
	err := runNew("test_name", dirName)
	require.Nil(s.T(), err)

	var config config.Config
	_, err = toml.DecodeFile("Wombats.toml", &config)
	require.Nil(s.T(), err)
	assert.Equal(s.T(), "test_name", config.Package.Name)
	err = os.Chdir("../")
	require.Nil(s.T(), err)
}
