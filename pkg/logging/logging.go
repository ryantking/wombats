package logging

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

var (
	info = color.New(color.FgGreen).Add(color.Bold)
	warn = color.New(color.FgYellow).Add(color.Bold)
	err  = color.New(color.FgRed).Add(color.Bold)
)

// NewLogrusHook returns the hook for adding to logrus
func NewLogrusHook() LogrusHook {
	return LogrusHook{
		levels: []log.Level{
			log.InfoLevel,
			log.WarnLevel,
			log.ErrorLevel,
			log.FatalLevel,
			log.PanicLevel,
		},
	}
}

// Levels returns the logging levels
func (lh LogrusHook) Levels() []log.Level {
	return lh.levels
}

// Fire prints out logrus messages to stdout
func (lh LogrusHook) Fire(e *log.Entry) error {
	switch e.Level {
	case log.InfoLevel:
		words := strings.Split(e.Message, " ")
		info.Printf("%s ", words[0])
		fmt.Print(strings.Join(words[1:], " "))
	case log.WarnLevel:
		warn.Printf("warning: ")
		fmt.Print(e.Message)
	case log.ErrorLevel:
		err.Printf("error: ")
		fmt.Print(e.Message)
	case log.FatalLevel:
		err.Printf("error: ")
		fmt.Print(e.Message)
	case log.PanicLevel:
		err.Printf("error: ")
		fmt.Print(e.Message)
	}
	fmt.Println()
	return nil
}
