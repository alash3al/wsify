package main

import "net/http"
import "strings"
import "github.com/gorilla/websocket"
import "github.com/go-redis/redis"

var (
	WSUpgrader = websocket.Upgrader{
		ReadBufferSize:    100,
		WriteBufferSize:   1024 * 8,
		CheckOrigin:       func(r *http.Request) bool { return IsAllowedOrigin(r.Host) },
		EnableCompression: true,
	}
)

// the websocket request handler
func WSHandler(w http.ResponseWriter, r *http.Request) {
	defer recover()
	channel := strings.Trim(r.URL.Path, "/")
	authorization := r.Header.Get("Authorization")
	token := r.FormValue("authorization")
	if authorization == "" {
		authorization = "Bearer " + token
	}
	if !IsAllowedClient(*FLAG_AUTH_BACKEND, authorization, channel) {
		http.Error(w, "Your aren't allowed to access this resource", 401)
		return
	}

	conn, err := WSUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	disconnected := make(chan bool)
	unsub := make(chan bool)
	inbox := make(chan *redis.Message)

	conn.SetCloseHandler(func(_ int, _ string) error {
		disconnected <- true
		return nil
	})

	go (func() {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}
		disconnected <- true
	})()

	go RedisSubscribe(REDIS_CLIENT, channel, inbox, unsub)

	ended := false

	for !ended {
		select {
		case <-disconnected:
			ended = true
			unsub <- true
			close(inbox)
			close(unsub)
		case msg := <-inbox:
			var parsedMsg *Message
			if err := JsonDecode([]byte(msg.Payload), &parsedMsg); err != nil {
				continue
			} else if (len(parsedMsg.To) >= 1) && ! parsedMsg.IsUserAllowed(authorization) {
				continue
			} else if conn.WriteJSON(parsedMsg.Payload) != nil {
				disconnected <- true
			}
		}
	}
}

// start tne websocket server
func InitWsServer(addr string) error {
	return http.ListenAndServe(addr, http.HandlerFunc(WSHandler))
}
