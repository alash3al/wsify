package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type Event struct {
	Action string `json:"action"`
	Key    string `json:"key,omitempty"`
	Value  string `json:"value,omitempty"`
}

func TriggerWebhook(ev Event) bool {
	ev.Action = strings.ToLower(ev.Action)
	jdata, _ := json.Marshal(ev)
	reader := bytes.NewReader(jdata)
	if _, found := WEBHOOK_EVENTS[ev.Action]; !found {
		return true
	}
	resp, err := http.Post(*FLAG_WEBHOOK_URL, "application/json", reader)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode > 200 {
		return false
	}
	return true
}
