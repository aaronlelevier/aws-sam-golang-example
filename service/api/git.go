package main

import (
	"encoding/json"
	"io"
)

type GithubEventPullRequest struct {
	Merged bool
	Url string
}

type GithubEvent struct {
	Action       string
	Pull_request GithubEventPullRequest
}

func DecodeGithubEvent(r io.Reader) GithubEvent {
	decoder := json.NewDecoder(r)
	var event GithubEvent
	err := decoder.Decode(&event)
	if err != nil {
		panic(err)
	}
	return event
}
