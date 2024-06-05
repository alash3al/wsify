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
	interceptorEndpointURL string
	webServerListenAddress string
	webServerPublishingKey string
}

func NewFromEnv(envFilename string) (*Config, error) {
	if err := godotenv.Load(envFilename); err != nil {
		return nil, err
	}

	return &Config{
		logger:                 slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		brokerDriver:           os.Getenv("BROKER_DRIVER"),
		brokerDSN:              os.Getenv("BROKER_DSN"),
		interceptorEndpointURL: os.Getenv("INTERCEPTOR_ENDPOINT_URL"),
		webServerListenAddress: os.Getenv("SERVER_LISTEN_ADDR"),
		webServerPublishingKey: os.Getenv("SERVER_PUBLISHING_KEY"),
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
func (c *Config) GetInterceptorEndpointURL() string {
	return c.interceptorEndpointURL
}

func (c *Config) GetWebServerListenAddr() string {
	return c.webServerListenAddress
}

func (c *Config) GetWebServerPublishingKey() string { return c.webServerPublishingKey }
