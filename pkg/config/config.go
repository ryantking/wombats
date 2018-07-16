package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	gitconfig "github.com/tcnksm/go-gitconfig"
)

// ErrNoConfigFile is to be thrown when a config file cannot be located in the
// current directory or any parent directory
var ErrNoConfigFile = fmt.Errorf("Could not locate config file")

// New creaates an initial config for a new project.
func New(name string, small bool) *Config {
	authors := []string{}
	username, err := gitconfig.Username()
	if err != nil {
		log.Errorf("Could not retrieve username from gitconfig.")
	} else {
		authors = []string{username}
	}

	pkgCfg := PackageConfig{
		Name:    name,
		Authors: authors,
		Version: "v0.1",
		Small:   small,
	}

	return &Config{Package: pkgCfg}
}

// Read reads an existing config from a file
func Read() (*Config, error) {
	var config *Config
	for {
		if _, err := toml.DecodeFile("Wombats.toml", &config); err == nil {
			break
		}

		if wd, err := os.Getwd(); err != nil {
			return nil, err
		} else if wd == "/" {
			return nil, ErrNoConfigFile
		}

		if err := os.Chdir("../"); err != nil {
			return nil, err
		}
	}

	return config, nil
}

// Write writes the config to the Wombats.toml file
func (c *Config) Write() error {
	f, err := os.Create("Wombats.toml")
	if err != nil && !os.IsExist(err) {
		return err
	}
	defer f.Close()

	return toml.NewEncoder(f).Encode(c)
}
