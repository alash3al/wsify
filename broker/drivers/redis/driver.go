package redisbroker

import (
	"context"
	"encoding/json"
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

func (d *Driver) Subscribe(ctx context.Context, channels []string) (<-chan any, error) {
	pubSub := d.conn.Subscribe(ctx, channels...)
	messagesChan := make(chan any)

	go (func() {
		var msg any
		for {
			res, err := pubSub.ReceiveMessage(ctx)
			if err != nil {
				break
			}
			if err := json.Unmarshal([]byte(res.Payload), &msg); err != nil {
				// TODO do something with the error
			}
			messagesChan <- msg
		}
	})()

	go (func() {
		select {
		case <-ctx.Done():
			if err := pubSub.Close(); err != nil {
				// TODO handle this
				fmt.Println("Unable to close from redis driver")
			}
		}
	})()

	return messagesChan, nil
}

func (d *Driver) Publish(ctx context.Context, channel string, msg any) error {
	j, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return d.conn.Publish(ctx, channel, string(j)).Err()
}

func (d *Driver) Close() error {
	return d.conn.Close()
}
