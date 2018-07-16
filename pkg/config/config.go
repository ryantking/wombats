package config

import (
	"os"

	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	gitconfig "github.com/tcnksm/go-gitconfig"
)

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

// Write writes the config to the Wombats.toml file
func (c *Config) Write() error {
	f, err := os.Create("Wombats.toml")
	if err != nil && !os.IsExist(err) {
		return err
	}
	defer f.Close()

	return toml.NewEncoder(f).Encode(c)
}
