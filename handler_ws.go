package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alash3al/wsify/broker"
	"github.com/alash3al/wsify/config"
	"github.com/gofiber/contrib/websocket"
)

func handleWebSocketRoute(cfg *config.Config, brokerConn broker.Driver) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		var parsedMsg WebSocketMessage
		cop := make(chan any)
		defer close(cop)

		go (func() {
			for msg := range cop {
				if err := conn.WriteJSON(msg); err != nil {
					// TODO handle this
					cfg.GetLogger().Error(err.Error(), "step", "writeJson")
				}
			}
		})()

		subscribedChannels := map[string]context.CancelFunc{}

		for {
			_, rawMsg, err := conn.ReadMessage()
			if err != nil {
				break
			}

			if err := json.Unmarshal(rawMsg, &parsedMsg); err != nil {
				// TODO handle this error
				cfg.GetLogger().Error(err.Error(), "step", "jsonUnmarshal")
			}

			switch parsedMsg.Command {
			case WebSocketMessageCommandTypeBroadcast:
				for _, channel := range parsedMsg.GetAvailableChannels() {
					if err := brokerConn.Publish(context.Background(), channel, parsedMsg.GetAvailableWebSocketMessage()); err != nil {
						// TODO handle this error
						cfg.GetLogger().Error(err.Error(), "step", "publish")
					}
				}
			case WebSocketMessageCommandTypeLeave:
				fmt.Println("Leaving ...")
				fmt.Println(subscribedChannels)
				for _, channel := range parsedMsg.GetAvailableChannels() {
					if cancel, found := subscribedChannels[channel]; found {
						cancel()
						delete(subscribedChannels, channel)
					}
				}
				fmt.Println(subscribedChannels)
				fmt.Println("=========================")
			case WebSocketMessageCommandTypeJoin:
				for _, channel := range parsedMsg.GetAvailableChannels() {
					ctx, cancel := context.WithCancel(context.Background())
					subscribedChannels[channel] = cancel
					msgsChan, err := brokerConn.Subscribe(ctx, []string{channel})
					if err != nil {
						// TODO handle this error
						cfg.GetLogger().Error(err.Error(), "step", "Subscribe")
					}

					go (func() {
						for msg := range msgsChan {
							cop <- msg
						}
						fmt.Println("Channel Listener Ended: ", channel)
					})()
				}
			}
		}
	}
}
