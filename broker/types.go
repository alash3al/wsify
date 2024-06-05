package broker

import (
	"context"
)

type Driver interface {
	Connect(dsn string) error
	Subscribe(ctx context.Context, channels []string) (<-chan any, error)
	Publish(ctx context.Context, channel string, msg any) error
	Close() error
}
