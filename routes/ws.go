package routes

import (
	"github.com/alash3al/wsify/broker"
	"github.com/alash3al/wsify/config"
	"github.com/alash3al/wsify/session"
	"github.com/alash3al/wsify/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"net/http"
)

func WebsocketRouteHandler(cfg *config.Config, brokerConn broker.Driver) echo.HandlerFunc {
	return func(c echo.Context) error {
		canConnect, err := utils.ShouldAcceptPayload(cfg.GetAuthorizerEndpointURL(), session.Message{
			Command: session.MessageCommandTypeConnect,
			Args: map[string]any{
				"headers": c.Request().Header,
				"query":   c.QueryParams(),
			},
		})

		if err != nil {
			cfg.GetLogger().Error(err.Error(), "utils.ShouldAcceptPayload")
			return c.NoContent(http.StatusForbidden)
		}

		if !canConnect {
			return c.NoContent(http.StatusForbidden)
		}

		return echo.WrapHandler(websocket.Server{
			Handshake: func(c *websocket.Config, request *http.Request) error { return nil },
			Handler: websocket.Handler(func(conn *websocket.Conn) {
				sess := session.Session{
					Context:      conn.Request().Context(),
					Broker:       brokerConn,
					Config:       cfg,
					Conn:         conn,
					DoneChannels: make(map[string]chan struct{}),
					ErrChan:      make(chan error),
					Writer:       make(chan []byte),
				}

				go (func() {
					for err := range sess.ErrChan {
						cfg.GetLogger().Error(err.Error(), "func", "sessionErrorListener")
					}
				})()

				if err := sess.Serve(); err != nil {
					cfg.GetLogger().Error(err.Error(), "func", "session.Serve")
					return
				}
			}),
		})(c)
	}
}
