package memorybroker

import (
	"context"
	"github.com/savsgio/gotils/uuid"
	"sync"
)

type Driver struct {
	sync.RWMutex
	subscriptions map[string]map[string]chan any
	cancelFuncs   map[string]context.CancelFunc
}

func (d *Driver) Connect(_ string) error {
	d.subscriptions = make(map[string]map[string]chan any)
	d.cancelFuncs = make(map[string]context.CancelFunc)

	return nil
}

func (d *Driver) Subscribe(ctx context.Context, channels []string) (<-chan any, error) {
	d.Lock()
	defer d.Unlock()

	id := uuid.V4()
	messagesChan := make(chan any)
	ctx, cancel := context.WithCancel(ctx)

	d.cancelFuncs[id] = cancel

	for _, channel := range channels {
		if _, found := d.subscriptions[channel]; !found {
			d.subscriptions[channel] = make(map[string]chan any)
		}
		d.subscriptions[channel][id] = messagesChan
	}

	go (func() {
		select {
		case <-ctx.Done():
			d.Lock()
			close(messagesChan)
			for _, channel := range channels {
				delete(d.subscriptions[channel], id)
				delete(d.cancelFuncs, id)
			}
			d.Unlock()
		}
	})()

	return messagesChan, nil
}

func (d *Driver) Publish(_ context.Context, channel string, msg any) error {
	d.Lock()
	defer d.Unlock()

	for _, subscriber := range d.subscriptions[channel] {
		go (func() {
			subscriber <- msg
		})()
	}

	return nil
}

func (d *Driver) Close() error {
	d.Lock()
	defer d.Unlock()

	for _, cancel := range d.cancelFuncs {
		cancel()
	}

	return nil
}
