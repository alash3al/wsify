package main

import (
	"flag"
	"strings"
)

var (
	//FlagHTTPAddr ...
	FlagHTTPAddr = flag.String("listen", ":4040", "the http address to listen on")
	//FlagAllowedOrigin ...
	FlagAllowedOrigin = flag.String("origin", "*", "the allowed websocket origin(s), it accepts a comma separated list of domains, * means anything")
	//FlagWebhookURL ...
	FlagWebhookURL = flag.String("webhook", "http://localhost:8000", "the webhook")
	//FlagWebhookEvents ...
	FlagWebhookEvents = flag.String("events", "connect,disconnect,subscribe,unsubscribe", "the events to be sent to the webhook")
	//FlagPublishEndpoint ...
	FlagPublishEndpoint = flag.String("publish", "/publish", "the publish endpoint, just make it as secure as you can")
	//FlagDebug
	FlagDebug = flag.Bool("debug", false, "enable debugging mode")
	//Version ...
	Version = "2.3"
	//WebhookEvents ..
	WebhookEvents = map[string]bool{}
)

//InitFlags ...
func InitFlags() {
	flag.Parse()
	for _, e := range strings.Split(strings.ToLower(*FlagWebhookEvents), ",") {
		WebhookEvents[strings.TrimSpace(e)] = true
	}
}
