package main

import (
	"github.com/alash3al/wsify/broker"
	"github.com/alash3al/wsify/config"
	"github.com/alash3al/wsify/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
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

	srv := echo.New()
	srv.HideBanner = true

	srv.Use(middleware.CORS())
	srv.Use(middleware.Logger())

	srv.GET("/ws/:id", routes.WebsocketRouteHandler(cfg, brokerConn))
	srv.POST("/publish", routes.PublishHandler(cfg, brokerConn))

	log.Fatal(srv.Start(cfg.GetWebServerListenAddr()))
}
