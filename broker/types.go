package broker

import (
	"context"
)

type Driver interface {
	Connect(dsn string) error
	Subscribe(ctx context.Context, channel string) (<-chan []byte, chan struct{}, error)
	Publish(ctx context.Context, channel string, msg []byte) error
	Close() error
}
