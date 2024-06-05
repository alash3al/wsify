package routes

import (
	"encoding/json"
	"github.com/alash3al/wsify/broker"
	"github.com/alash3al/wsify/config"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func BroadcastHandler(cfg *config.Config, brokerConn broker.Driver) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.QueryParam("key") != cfg.GetWebServerBroadcastingKey() {
			return c.NoContent(http.StatusForbidden)
		}

		if strings.ToLower(c.Request().Header.Get("Content-Type")) != "application/json" {
			return c.NoContent(http.StatusUnsupportedMediaType)
		}

		var msg struct {
			Channel string `json:"channel"`
			Content any    `json:"content"`
		}

		if err := c.Bind(&msg); err != nil {
			cfg.GetLogger().Error(err.Error(), "func", "BroadcastHandler.Bind")
			return c.NoContent(http.StatusBadRequest)
		}

		j, err := json.Marshal(msg.Content)
		if err != nil {
			cfg.GetLogger().Error(err.Error(), "func", "BroadcastHandler.json.Unmarshal")
			return c.NoContent(http.StatusInternalServerError)
		}

		if err := brokerConn.Publish(c.Request().Context(), msg.Channel, j); err != nil {
			cfg.GetLogger().Error(err.Error(), "func", "BroadcastHandler.broker.Publish")
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusCreated)
	}
}
