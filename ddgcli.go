package main

import (
	"log"
	"os"

	"github.com/mohsalsaleem/ddgo-cli/api"
	"github.com/mohsalsaleem/ddgo-cli/renderer"

	"github.com/mohsalsaleem/ddgo-cli/command"
)

func main() {
	com, err := command.Parse(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	html, err := api.Search(com)
	if err != nil {
		log.Fatal(err)
	}
	renderer.Render(html)
}
