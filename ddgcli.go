package main

import (
	"fmt"

	"github.com/mohsalsaleem/ddgo-cli/api"
	"github.com/mohsalsaleem/ddgo-cli/command"
	"github.com/mohsalsaleem/ddgo-cli/renderer"
)

func main() {
	com, err := command.Parse()
	if err != nil {
		fmt.Println(err.Error())
	}
	html, err := api.Search(com)
	if err != nil {
		fmt.Println(err.Error())
	}
	renderer.Render(html)
}
