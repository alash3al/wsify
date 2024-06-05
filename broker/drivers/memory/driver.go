package memorybroker

import (
	"context"
	"github.com/savsgio/gotils/uuid"
	"sync"
)

type Driver struct {
	sync.RWMutex
	subscriptions map[string]map[string]chan []byte
	doneChannels  map[string]chan struct{}
}

func (d *Driver) Connect(_ string) error {
	d.subscriptions = make(map[string]map[string]chan []byte)
	d.doneChannels = make(map[string]chan struct{})

	return nil
}

func (d *Driver) Subscribe(ctx context.Context, channel string) (<-chan []byte, chan struct{}, error) {
	d.Lock()
	defer d.Unlock()

	id := uuid.V4()
	messagesChan := make(chan []byte)
	doneChan := make(chan struct{})

	d.doneChannels[id] = doneChan

	if _, found := d.subscriptions[channel]; !found {
		d.subscriptions[channel] = make(map[string]chan []byte)
	}

	d.subscriptions[channel][id] = messagesChan

	done := func() {
		d.Lock()
		defer d.Unlock()

		close(messagesChan)

		delete(d.subscriptions[channel], id)
		delete(d.doneChannels, id)

	}

	go (func() {
		select {
		case <-doneChan:
			done()
		case <-ctx.Done():
			done()
		}
	})()

	return messagesChan, doneChan, nil
}

func (d *Driver) Publish(_ context.Context, channel string, msg []byte) error {
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

	for _, ch := range d.doneChannels {
		ch <- struct{}{}
	}

	return nil
}
