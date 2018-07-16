package logging

import (
	log "github.com/sirupsen/logrus"
)

// LogrusHook is a hook to print out logrus info to stdout
type LogrusHook struct {
	levels []log.Level
}
