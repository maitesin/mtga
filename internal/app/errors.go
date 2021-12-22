package app

import "fmt"

const errMsgInvalidCommand = "invalid command %q received. Expected %q"

type InvalidCommandError struct {
	Received Command
	Expected Command
}

func (ice InvalidCommandError) Error() string {
	return fmt.Sprintf(errMsgInvalidCommand, ice.Received.Name(), ice.Expected.Name())
}
