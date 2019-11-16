package api

import (
	"testing"

	"github.com/mohsalsaleem/ddgo-cli/command"
)

func TestSearch(t *testing.T) {
	com := command.Command{QueryString: "hello+world"}
	_, err := Search(&com)
	if err != nil {
		t.Error("Error in Search \n", err.Error())
	}
}
