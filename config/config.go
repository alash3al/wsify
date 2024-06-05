package config

import (
	"flag"
	"log/slog"
	"os"
)

type Config struct {
	logger                   *slog.Logger
	brokerDriver             string
	brokerDSN                string
	authorizerEndpointURL    string
	webServerListenAddress   string
	webServerBroadcastingKey string
}

func NewFromFlags() (*Config, error) {
	brokerDriver := flag.String("broker-driver", "memory", "the message broker driver (redis, memory)")
	brokerDSN := flag.String("broker-dsn", "", "the selected driver DSN (connection url), example: redis://localhost")
	authorizerURL := flag.String("authorizer-url", "", "the endpoint url that will be used as the main authorizer webhook")
	listenAddr := flag.String("listen-addr", ":3000", "the web server listen address")
	broadcastingKey := flag.String("broadcasting-key", "", "key that will authorize all `/broadcast` calls")

	flag.Parse()

	return &Config{
		logger:                   slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		brokerDriver:             *brokerDriver,
		brokerDSN:                *brokerDSN,
		authorizerEndpointURL:    *authorizerURL,
		webServerListenAddress:   *listenAddr,
		webServerBroadcastingKey: *broadcastingKey,
	}, nil
}

func (c *Config) GetLogger() *slog.Logger {
	return c.logger
}

func (c *Config) GetBrokerDriver() string {
	return c.brokerDriver
}

func (c *Config) GetBrokerDSN() string {
	return c.brokerDSN
}
func (c *Config) GetAuthorizerEndpointURL() string {
	return c.authorizerEndpointURL
}

func (c *Config) GetWebServerListenAddr() string {
	return c.webServerListenAddress
}

func (c *Config) GetWebServerBroadcastingKey() string { return c.webServerBroadcastingKey }
