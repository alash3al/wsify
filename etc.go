package main

import (
	"net/http"
	"net/url"
	"encoding/json"
)

import "github.com/go-redis/redis"

// the message that will be polled from redis
type Message struct {
	Payload		interface{} `json:"payload"`
	To 			[]string 	`json:"to"`
}

// checks whether the specified user is allowed 
// to recieve this message or not
func (m *Message) IsUserAllowed(u string) bool {
	for _, v := range m.To {
		if v == u {
			return true
		}
	}
	return false
}

// decode a json data
func JsonDecode(data []byte, o interface{}) error {
	return json.Unmarshal(data, o)
}

// check whether the specified hostname is in the
// allowed origins
func IsAllowedOrigin(hostname string) bool {
	for _, origin := range ALLOWED_ORIGINS {
		if origin == "*" || origin == hostname {
			return true
		}
	}
	return false
}

// check whether the specified client "authKey" is allowed to
// access the specified channel or not.
func IsAllowedClient(authBackend, authKey, channel string) bool {
	backend, _ := url.Parse(authBackend)
	query := backend.Query()
	query.Set("channel", url.QueryEscape(channel))
	backend.RawQuery = query.Encode()
	authBackend = backend.String()
	req, err := http.NewRequest("POST", authBackend, nil)
	if err != nil {
		return false
	}
	req.Header.Set("Authorization", authKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return false
	}
	return true
}

// Connect to redis server
func RedisConnect(dsn string) (error, *redis.Client) {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return err, nil
	}
	client := redis.NewClient(opt)
	_, err = client.Ping().Result()
	if err != nil {
		return err, nil
	}
	return nil, client
}

// Subscribe to a redis channel and recieve the messages on the specified
// `out` chan, also listen for close signal on `end` chan.
func RedisSubscribe(rc *redis.Client, channel string, out chan *redis.Message, end chan bool) {
	pubsub := rc.Subscribe(channel)
	messages := pubsub.Channel()
	ended := false
	defer pubsub.Close()
	for !ended {
		select {
		case <-end:
			ended = true
		case msg := <-messages:
			out <- msg
		}
	}
}
