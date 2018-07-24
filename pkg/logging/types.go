package logging

import (
	log "github.com/sirupsen/logrus"
)

// ATSErrorType represents the type of an ATS error
type ATSErrorType int

const (
	// LineError holds information about an error occuring on a specific line
	LineError ATSErrorType = iota
	// ErrorCount holds the number of errors
	ErrorCount
	// TypeError holds info about a type error
	TypeError
)

// LogrusHook is a hook to print out logrus info to stdout
type LogrusHook struct {
	levels []log.Level
}

// ATSErrorLine holds information about various ATS errors
type ATSErrorLine interface {
	Type() ATSErrorType
}

// ATSLineError holds information about an error reported by patscc
type ATSLineError struct {
	Fname                      string
	StartL, EndL, StartO, EndO int
	Errors                     []string
}

// ATSErrorCount holds the number of errors that occured
type ATSErrorCount struct {
	Count int
}

// ATSTypeError holds info about an ATS type error
type ATSTypeError struct {
	Status  string
	ATSType string
}
