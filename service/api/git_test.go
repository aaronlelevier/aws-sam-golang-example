package main

import (
	"strings"
	"testing"
)

func TestDecodeGithubEvent(t *testing.T) {
	const jsonEvent = `
		{
			"action": "create",
			"pull_request": {
				"merged": true,
				"url": "URL-A"
			}
	  }
	`

	event := DecodeGithubEvent(strings.NewReader(jsonEvent))

	assertStringEquals(t, event.Action, "create")
	assertTrue(t, event.Pull_request.Merged)
	assertStringEquals(t, event.Pull_request.Url, "URL-A")
}

// helpers

func assertStringEquals(t *testing.T, s string, s2 string) bool {
	ret := s == s2

	if s != s2 {
		t.Fatalf("Want %v Got %v", s2, s)
	}

	return ret
}

func assertTrue(t *testing.T, b bool) bool {
	if b != true {
		t.Fatalf("Want %v Got %v", true, b)
	}
	return true
}
