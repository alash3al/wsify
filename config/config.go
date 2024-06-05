package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type Config struct {
	logger                 *slog.Logger
	brokerDriver           string
	brokerDSN              string
	webServerListenAddress string
	interceptorEndpointURL string
}

func NewFromEnv(envFilename string) (*Config, error) {
	if err := godotenv.Load(envFilename); err != nil {
		return nil, err
	}

	return &Config{
		logger:                 slog.Default(),
		brokerDriver:           os.Getenv("BROKER_DRIVER"),
		brokerDSN:              os.Getenv("BROKER_DSN"),
		webServerListenAddress: os.Getenv("SERVER_LISTEN_ADDR"),
		interceptorEndpointURL: os.Getenv("INTERCEPTOR_ENDPOINT_URL"),
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

func (c *Config) GetWebServerListenAddr() string {
	return c.webServerListenAddress
}

func (c *Config) GetInterceptorEndpointURL() string {
	return c.interceptorEndpointURL
}
