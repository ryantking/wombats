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
