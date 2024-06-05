package routes

import (
	"github.com/alash3al/wsify/broker"
	"github.com/alash3al/wsify/config"
	"github.com/alash3al/wsify/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"net/http"
)

func WebsocketRouteHandler(cfg *config.Config, broker broker.Driver) echo.HandlerFunc {
	return func(c echo.Context) error {
		return echo.WrapHandler(websocket.Server{
			Handshake: func(c *websocket.Config, request *http.Request) error { return nil },
			Handler: websocket.Handler(func(conn *websocket.Conn) {
				session := session.Session{
					Context:      conn.Request().Context(),
					Broker:       broker,
					Conn:         conn,
					DoneChannels: make(map[string]chan struct{}),
					ErrChan:      make(chan error),
					Writer:       make(chan []byte),
				}

				go (func() {
					for err := range session.ErrChan {
						cfg.GetLogger().Error(err.Error(), "func", "sessionErrorListener")
					}
				})()

				if err := session.Serve(); err != nil {
					cfg.GetLogger().Error(err.Error(), "func", "session.Serve")
					return
				}
			}),
		})(c)
	}
}
