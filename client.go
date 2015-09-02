package go_pushover_client

import (
	"http"
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

func (p *PushoverClient) send(to string, message string, priority int32) {
	payload := url.Values{
		"token":    {p.ApplicationToken},
		"user":     {to},
		"message":  {message},
		"priority": {priority},
	}
	http.PostForm(pushoverURL, payload)
}

func NewPushoverClient(token string) {
	out := PushoverClient{}
	out.new(token)
	return &out
}
