package main

import (
	"log"
	"net/http"
	"strings"
)

import (
	"github.com/alash3al/go-pubsub"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	// Default websocket upgrader
	WSUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			for _, origin := range strings.Split(*FLAG_ALLOWED_ORIGIN, ",") {
				origin = strings.TrimSpace(origin)
				if origin == "*" || origin == r.Host {
					return true
				}
			}
			return false
		},
		EnableCompression: true,
	}

	// Default broker
	Broker = pubsub.NewBroker()
)

// the websocket request handler
func WSHandler(c echo.Context) error {
	defer (func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	})()
	key := c.QueryParam("key")
	allowed := TriggerWebhook(Event{
		Action: "connect",
		Key:    key,
	})
	if !allowed {
		return c.JSON(403, "You aren't allowed to access this resource")
	}
	conn, err := WSUpgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return nil
	}
	defer conn.Close()
	subscriber, err := Broker.Attach()
	if err != nil {
		conn.WriteJSON(map[string]string{
			"error": "Sorry, couldn't allocate resources for you",
		})
		return nil
	}
	closeCh := make(chan bool)
	closed := false
	go (func() {
		var action Event
		stop := false
		for !stop {
			if conn.ReadJSON(&action) != nil {
				stop = true
				break
			}
			if action.Action == "subscribe" || action.Action == "unsubscribe" {
				if !TriggerWebhook(Event{Action: action.Action, Key: key, Value: action.Value}) {
					conn.WriteJSON(map[string]string{
						"error": "You aren't allowed to access the requested resource",
					})
					continue
				}
			}
			if action.Action == "subscribe" {
				Broker.Subscribe(subscriber, action.Value)
			} else if action.Action == "unsubscribe" {
				Broker.Unsubscribe(subscriber, action.Value)
			}
		}
		closeCh <- true
	})()
	for !closed {
		select {
		case <-closeCh:
			closed = true
			close(closeCh)
			Broker.Detach(subscriber)
			TriggerWebhook(Event{Action: "disconnect", Key: key})
		case data := <-subscriber.GetMessages():
			msg := (data.GetPayload()).(Message)
			if !msg.IsUserAllowed(key) {
				continue
			}
			msg.Topic = data.GetTopic()
			msg.Time = data.GetCreatedAt()
			msg.To = nil
			if conn.WriteJSON(msg) != nil {
				closeCh <- true
			}
		}
	}
	return nil
}

// publish handler
func PublishHandler(c echo.Context) error {
	var msg Message
	if err := c.Bind(&msg); err != nil {
		return c.JSON(422, map[string]interface{}{
			"sucess": false,
			"error":  err.Error(),
		})
	}
	Broker.Broadcast(msg, msg.Topic)
	return c.JSON(200, map[string]interface{}{
		"sucess": true,
		"data":   msg,
	})
}

// welcome handler
func WelcomeHandler(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"success": true,
		"message": "let's go ?",
	})
}

// start the websocket server
func InitWsServer(addr string) error {
	e := echo.New()

	e.Debug = true
	e.HideBanner = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 9}))

	e.GET("/subscribe", WSHandler)
	e.POST(*FLAG_PUBLISH_ENDPOINT, PublishHandler)

	return e.Start(addr)
}
