package redisbroker

import (
	"context"
	"fmt"
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

func (d *Driver) Subscribe(ctx context.Context, channels []string) (<-chan []byte, chan struct{}, error) {
	pubSub := d.conn.Subscribe(ctx, channels...)
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
		if err := pubSub.Close(); err != nil {
			// TODO handle this
			fmt.Println("Unable to close from redis pubSub")
		}
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
