package main

import (
	"flag"
	"strings"
)

var (
	FLAG_HTTP_ADDR        = flag.String("listen", ":4040", "the http address to listen on")
	FLAG_ALLOWED_ORIGIN   = flag.String("origin", "*", "the allowed websocket origin(s), it accepts a comma separated list of domains, * means anything")
	FLAG_WEBHOOK_URL      = flag.String("webhook", "http://localhost:8000", "the webhook")
	FLAG_WEBHOOK_EVENTS   = flag.String("events", "connect,disconnect,subscribe,unsubscribe", "the events to be sent to the webhook")
	FLAG_PUBLISH_ENDPOINT = flag.String("publish", "/publish", "the publish endpoint, just make it as secure as you can")
)

var (
	VERSION        = "2.0"
	WEBHOOK_EVENTS = map[string]bool{}
)

func InitFlags() {
	flag.Parse()
	for _, e := range strings.Split(strings.ToLower(*FLAG_WEBHOOK_EVENTS), ",") {
		WEBHOOK_EVENTS[strings.TrimSpace(e)] = true
	}
}
