package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/mohsalsaleem/ddgo-cli/command"
)

const ddgAPI = "https://duckduckgo.com/html/?q="

func displayProgress(done chan bool, queryString string, delay time.Duration) {

	var stop bool = false

	for {
		select {
		case <-done:
			fmt.Fprint(os.Stdout, "\r \r")
			stop = true
		default:
			fmt.Printf("\r\u001b[36mSearching - '%s'\u001b[0m", queryString)
		}

		if stop {
			break
		}

		time.Sleep(delay)
	}
}

// Search - function to search duckduckgo
func Search(com *command.Command) (string, error) {

	done := make(chan bool)

	go displayProgress(done, com.QueryString, 5*time.Millisecond)

	var url string = ddgAPI + url.QueryEscape(com.QueryString)

	client := http.Client{}
	request, _ := http.NewRequest("GET", url, nil)

	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Accept-Language", "en-US,en;q=0.5")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Host", "duckduckgo.com")
	request.Header.Set("Pragma", "no-cache")
	request.Header.Set("TE", "Trailers")
	request.Header.Set("Upgrade-Insecure-Requests", "1")
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0")

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(bodyBytes)

	done <- true

	return bodyString, nil
}
