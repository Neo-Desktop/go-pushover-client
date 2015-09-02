package go_pushover_client

import (
	"log"
	"net/http"
	"net/url"
)

const (
	pushoverURL = "https://api.pushover.net/1/messages.json"
)

type PushoverClient struct {
	ApplicationToken string
}

func (p *PushoverClient) new(token string) {
	p.ApplicationToken = token
}

func (p *PushoverClient) Send(to string, message string, priority string) {

	payload := url.Values{
		"token":    {p.ApplicationToken},
		"user":     {to},
		"message":  {message},
		"priority": {priority},
	}

	resp, err := http.PostForm(pushoverURL, payload)
	if err != nil {
		log.Fatal(err)
	}

	responseBytes := make([]byte, 1024)
	n, err := resp.Body.Read(responseBytes)

	responseString := string(responseBytes[:n])

	log.Println(responseString)
}

func New(token string) *PushoverClient {
	out := PushoverClient{}
	out.new(token)
	return &out
}
