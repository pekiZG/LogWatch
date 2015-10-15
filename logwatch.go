package main

import (
	"net/http"
	"net/url"
)

func main() {
	slackIncomingWebhookAPI := "" // Put your incoming webHookURL here!
	message := `{"text": "Hello World from GoLang!"}`
	resp, err := http.PostForm(slackIncomingWebhookAPI, url.Values{"payload": {message}})

	if err != nil {
		// handle error here, but I'm just lazy
	}
	defer resp.Body.Close()
}
