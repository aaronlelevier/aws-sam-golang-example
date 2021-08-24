package main

import (
	"strings"
	"testing"
)

func TestDecodePayload(t *testing.T) {
	const jsonSteam = `{"foo": "bar"}`
	value := DecodePayload(strings.NewReader(jsonSteam))
	if value != "bar" {
		t.Fatalf("Want %v Got %v", "bar", value)
	}
}
