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
	// IncludeError represents an include error
	IncludeError
)

// LogrusHook is a hook to print out logrus info to stdout
type LogrusHook struct {
	levels []log.Level
}

// ATSErrorLine holds information about various ATS errors
type ATSErrorLine interface {
	Type() ATSErrorType
	Print()
}

// ATSLineError holds information about an error reported by patscc
type ATSLineError struct {
	Fname                      string
	StartL, EndL, StartO, EndO int
	Error                      string
}

// Type ...
func (e *ATSLineError) Type() ATSErrorType {
	return LineError
}

// ATSCountError holds the number of errors that occured
type ATSCountError struct {
	Count int
}

// Type ...
func (e *ATSCountError) Type() ATSErrorType {
	return ErrorCount
}

// ATSTypeError holds info about an ATS type error
type ATSTypeError struct {
	Status  string
	ATSType string
}

// Type ...
func (e *ATSTypeError) Type() ATSErrorType {
	return TypeError
}

// ATSIncludeError holds info about an ATS include error
type ATSIncludeError struct {
	Fname, IncludeFname        string
	StartL, EndL, StartO, EndO int
}

// Type ...
func (e *ATSIncludeError) Type() ATSErrorType {
	return IncludeError
}
