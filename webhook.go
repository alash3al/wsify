package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

// Event ...
type Event struct {
	Action string `json:"action"`
	Key    string `json:"key,omitempty"`
	Value  string `json:"value,omitempty"`
}

// TriggerWebhook ...
func TriggerWebhook(ev Event) bool {
	if *FlagWebhookURL == "" {
		return true
	}
	ev.Action = strings.ToLower(ev.Action)
	jdata, _ := json.Marshal(ev)
	reader := bytes.NewReader(jdata)
	if _, found := WebhookEvents[ev.Action]; !found {
		return true
	}
	resp, err := http.Post(*FlagWebhookURL, "application/json", reader)
	defer func() {
		if resp != nil {
			resp.Body.Close() // nolint: errcheck
		}
	}()
	if err != nil || resp.StatusCode > 200 {
		return false
	}
	return true
}
