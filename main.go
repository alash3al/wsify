package main

import (
	"github.com/alash3al/wsify/broker"
	"github.com/alash3al/wsify/config"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"

	_ "github.com/alash3al/wsify/broker/drivers/memory"
	_ "github.com/alash3al/wsify/broker/drivers/redis"
)

func main() {
	envFilename := ".env"
	if len(os.Args) > 1 {
		envFilename = os.Args[1]
	}

	cfg, err := config.NewFromEnv(envFilename)
	if err != nil {
		panic(err.Error())
	}

	brokerConn, err := broker.Connect(cfg.GetBrokerDriver(), cfg.GetBrokerDSN())
	if err != nil {
		panic(err.Error())
	}

	server := fiber.New()

	server.Use(logger.New())

	// TODO check with interceptor
	server.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// TODO handle client id
	// TODO set connection close handler
	// TODO more code cleanup
	// TODO integrate interceptor (with retries if possible)
	server.Get("/ws/:clientId", websocket.New(handleWebSocketRoute(cfg, brokerConn)))

	panic(server.Listen(cfg.GetWebServerListenAddr()))
}
