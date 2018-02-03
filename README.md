Websocketify (wsify) v1.0
=========================
> Just a tiny, simple and realtime pub/sub messaging service

Why
====
> I wanted to create a tiny solution that can replace `pusher` and similar services and learning more about the realtime world, so I dispatched this project.

Features
================
- No dependencies, just a single binary !
- Light and Tiny.
- Event-Driven Design `webhooks`.
- A client can listen on any resource.
- You control whether a client is allowed to `connect`, `subscribe`, `unsubscribe` using any programming language !.
- A client defines itself using `key` via the url query param i.e `?key=123`.
- Send messages to only certain users.


Installation
==============
- **Docker ?** > `docker run --network host alash3al/wsify -listen :8080 -auth-webhook "http://localhost/auth.php"`   
- **Binary ?** > goto the [releases](https://github.com/alash3al/wsify/releases) page and download yours.
- **From Source ?** > `go get -u github.com/alash3al/wsify`

Questions
==========

### (1)- How can a client/device connect to the websocket service?
> by simply connecting to the following endpoint `ws://your.wsify.service:port/subscribe`

### (2)- How can a client subscribe to a certain channel(s)/topic(s)?
> after connecting to the main websocket service `/subscribe`, you can send a simple json payload `commands` to ask wsify to `subscribe`/`unsubscribe` you to/from any channel/topic you want!

### (3)- What is the commands format?
>
```json
{
	"action": "subscribe",
	"value": "testchan"
}

```

### (4)- Can I control the client command so I can allow/disallow certain users?
> Yes, each client can define itself using a query param `?key=client1`, this key will be passed to the `webhook` endpoint
as well as the event being executed, and here is the event format:
```javascript
{
	// one of the following: connect|subscribe|unsubscribe|disconnect
	"action": "subscribe",

	// the channel if provided
	"value": "testchan",

	// the key provided by the client
	"key": "client1"
}
```

### (5)- How can I publish message to i.e `testchan`?
> Just a post request to `/publish` with the following format:
```javascript
{
	// the channel you want to publish to
	"channel": "testchan",

	// the data to be send (any format)
	"payload": "testchan",

	// array of clients "keys" (if you want certain clients only to receive the message)
	"to": []
}
```
i.e
```bash
curl -X POST \
	-H "Content-Type: application/json" \
	-d '{"payload": "hi from the terminal", "channel": "testchan"}' \
	http://localhost:4040/publish
```

### (6)- Can I skip the webhook events for testing?
> Yes, `wsify --events=""` empty events means "NO WEBHOOK, WSIFY!"

Demo
======
> I'm using [this chrome extension](https://chrome.google.com/webstore/detail/pfdhoblngboilpfeibdedpjgfnlcodoo) for websocket testing
![Quick Demo](https://i.imgur.com/jxyejg0.gif)

Author
=============
This project has been created by [Mohamed Al Ashaal](http://github.com/alash3al) a Crazy Gopher ^^!

Contribution
=============
- Fork the Repo
- Create a feature branch
- Push your changes to the created branch
- Create a pull request.

License
=============
Wsify is licensed under MIT license
