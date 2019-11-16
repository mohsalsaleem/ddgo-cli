package main

import (
	"fmt"
	"os"

	"github.com/mohsalsaleem/ddgo-cli/api"
	"github.com/mohsalsaleem/ddgo-cli/renderer"

	"github.com/mohsalsaleem/ddgo-cli/command"
)

func main() {
	com, err := command.Parse(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
	html, err := api.Search(com)
	if err != nil {
		fmt.Println(err.Error())
	}
	renderer.Render(html)
}
