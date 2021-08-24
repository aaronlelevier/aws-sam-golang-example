package main

import (
	"io"
	"encoding/json"
)

type Event struct {
	Action string
}

func EventAction(r io.Reader) string {
	decoder := json.NewDecoder(r)
	var t Event
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	return t.Action
}