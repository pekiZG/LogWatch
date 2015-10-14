package main

import (
	"net/http"
	"net/url"
	"strings"
)

func main() {
	slackAPI := "" // Put your incoming webHookURL here!

	client := http.Client{}

	form := url.Values{"text": {"This is a line of text in a channel."}}

	req, err := http.NewRequest("POST", slackAPI, strings.NewReader(form.Encode()))

	resp, err := client.Do(req)

	if err != nil {
		// handle error here, but I'm just lazy
	}
	defer resp.Body.Close()

}
