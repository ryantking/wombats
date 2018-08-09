package logging

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

var (
	infofmt  = color.New(color.FgGreen).Add(color.Bold)
	debugfmt = color.New(color.FgBlue).Add(color.Bold)
	warnfmt  = color.New(color.FgYellow).Add(color.Bold)
	errfmt   = color.New(color.FgRed).Add(color.Bold)
	boldfmt  = color.New(color.Bold)
	faintfmt = color.New(color.FgBlue).Add(color.Faint)
	spaces   = 4
)

// NewLogrusHook returns the hook for adding to logrus
func NewLogrusHook() LogrusHook {
	return LogrusHook{
		levels: []log.Level{
			log.InfoLevel,
			log.DebugLevel,
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

func printSpaces(inc bool) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}

	if inc {
		spaces++
	}
}

// Fire prints out logrus messages to stdout
func (lh LogrusHook) Fire(e *log.Entry) error {
	printSpaces(e.Level != log.DebugLevel)
	switch e.Level {
	case log.InfoLevel:
		words := strings.Split(e.Message, " ")
		infofmt.Printf("%s ", words[0])
		fmt.Print(strings.Join(words[1:], " "))
	case log.DebugLevel:
		debugfmt.Print("debug: ")
		fmt.Print(e.Message)
	case log.WarnLevel:
		warnfmt.Printf("warning: ")
		fmt.Print(e.Message)
	case log.ErrorLevel:
		errfmt.Printf("error: ")
		fmt.Print(e.Message)
	case log.FatalLevel:
		errfmt.Printf("error: ")
		fmt.Print(e.Message)
	case log.PanicLevel:
		errfmt.Printf("error: ")
		fmt.Print(e.Message)
	}
	fmt.Println()
	return nil
}
