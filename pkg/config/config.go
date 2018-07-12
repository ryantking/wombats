package config

import (
	log "github.com/Sirupsen/logrus"
	gitconfig "github.com/tcnksm/go-gitconfig"
)

// New creaates an initial config for a new project.
func New(name string) *Config {
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
	}

	return &Config{Package: pkgCfg}
}
