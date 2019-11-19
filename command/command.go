package command

import (
	"flag"
)

// Command - Command structure
type Command struct {
	QueryString string
	IgnoreCache bool
}

// Parse - Parse the cli arguments
func Parse() (*Command, error) {
	query := flag.String("q", "", "Enter a query: Example: ddgo-cli -q=\"What is github.com?\"")
	ignoreCache := flag.Bool("f", false, "Example: ddgo-cli -q=\"What is github.com?\" -f")
	flag.Parse()
	command := Command{}
	command.QueryString = *query
	command.IgnoreCache = *ignoreCache
	return &command, nil
}
