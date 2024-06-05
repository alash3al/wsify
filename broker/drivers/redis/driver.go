package redisbroker

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Driver struct {
	conn *redis.Client
}

func (d *Driver) Connect(dsn string) error {
	opts, err := redis.ParseURL(dsn)
	if err != nil {
		return err
	}

	d.conn = redis.NewClient(opts)

	if err := d.conn.Ping(context.Background()).Err(); err != nil {
		return err
	}

	return nil
}

func (d *Driver) Subscribe(ctx context.Context, channel string) (<-chan []byte, chan struct{}, error) {
	pubSub := d.conn.Subscribe(ctx, channel)
	messagesChan := make(chan []byte)
	doneChan := make(chan struct{})

	go (func() {
		for {
			res, err := pubSub.ReceiveMessage(ctx)
			if err != nil {
				break
			}
			messagesChan <- []byte(res.Payload)
		}
	})()

	done := func() {
		_ = pubSub.Close()
	}

	go (func() {
		select {
		case <-ctx.Done():
			done()
		case <-doneChan:
			done()
		}
	})()

	return messagesChan, doneChan, nil
}

func (d *Driver) Publish(ctx context.Context, channel string, msg []byte) error {
	return d.conn.Publish(ctx, channel, string(msg)).Err()
}

func (d *Driver) Close() error {
	return d.conn.Close()
}
