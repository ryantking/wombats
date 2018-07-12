package subcommands

import "fmt"

// ErrUnknownArguments is to be thrown when unknown arguments are passed to a
// subcommand.
type ErrUnknownArguments struct {
	args []string
}

func (e *ErrUnknownArguments) Error() string {
	return fmt.Sprintf("%v", e.args)
}

// ErrProjectExists is to be thrown when the project folder already exists.
type ErrProjectExists struct {
	name string
}

func (e *ErrProjectExists) Error() string {
	return fmt.Sprintf("Project: %s already exists", e.name)
}
