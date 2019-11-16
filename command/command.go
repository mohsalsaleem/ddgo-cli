package command

import (
	"errors"
	"net/url"
)

// Command - Command structure
type Command struct {
	QueryString string
	IgnoreCache bool
}

// Parse - Parse the command line arguments
func Parse(args []string) (*Command, error) {
	if len(args) < 2 {
		return nil, errors.New("Wrong number of arguments. Example: ddgo-cli \"What is github.com?\"")
	}
	command := Command{}
	command.QueryString = url.QueryEscape(args[1])
	if len(args) > 2 && (args[2] == "--ignore-cache" || args[2] == "-ic") {
		command.IgnoreCache = true
	}
	return &command, nil
}
