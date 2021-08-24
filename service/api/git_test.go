package main

import (
	"strings"
	"testing"
)

func TestEventAction(t *testing.T) {
	const jsonSteam = `{"Action": "create"}`
	value := EventAction(strings.NewReader(jsonSteam))
	if value != "create" {
		t.Fatalf("Want %v Got %v", "create", value)
	}
}
