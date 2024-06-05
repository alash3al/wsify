package session

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/alash3al/wsify/broker"
	"github.com/alash3al/wsify/config"
	"github.com/alash3al/wsify/utils"
	"golang.org/x/net/websocket"
	"io"
)

type Session struct {
	Context      context.Context
	Broker       broker.Driver
	Config       *config.Config
	Conn         *websocket.Conn
	Message      Message
	DoneChannels map[string]chan struct{}
	ErrChan      chan error
	Writer       chan []byte
}

func (s *Session) Serve() error {
	defer (func() {
		_ = s.Conn.Close()

		close(s.ErrChan)
		close(s.Writer)

		for _, ch := range s.DoneChannels {
			ch <- struct{}{}
		}
	})()

	go (func() {
		for output := range s.Writer {
			if err := websocket.Message.Send(s.Conn, string(output)); err != nil {
				s.ErrChan <- err
				break
			}
		}
	})()

	for {
		if err := websocket.JSON.Receive(s.Conn, &s.Message); err != nil {
			if errors.Is(err, io.EOF) {
				return err
			}

			s.ErrChan <- err
		}

		canProceed, err := utils.ShouldAcceptPayload(s.Config.GetAuthorizerEndpointURL(), s.Message)
		if err != nil {
			return err
		}

		if !canProceed {
			continue
		}

		switch s.Message.Command {
		case MessageCommandTypeJoin:
			s.onJoin()
		case MessageCommandTypeLeave:
			s.onLeave()
		case MessageCommandTypeBroadcast:
			s.onBroadcast()
		}
	}
}

func (s *Session) onJoin() {
	channel := s.Message.GetArgsChannel()

	if channel == "" {
		s.ErrChan <- errors.New("requested join on an empty chan")
		return
	}

	feed, done, err := s.Broker.Subscribe(s.Context, channel)
	if err != nil {
		s.ErrChan <- err
		return
	}

	s.DoneChannels[channel] = done

	go (func() {
		for msg := range feed {
			s.Writer <- msg
		}
	})()
}

func (s *Session) onLeave() {
	channel := s.Message.GetArgsChannel()
	s.DoneChannels[channel] <- struct{}{}
	delete(s.DoneChannels, channel)
}

func (s *Session) onBroadcast() {
	channel := s.Message.GetArgsChannel()

	j, err := json.Marshal(s.Message.GetArgsContent())
	if err != nil {
		s.ErrChan <- err
		return
	}

	if err := s.Broker.Publish(s.Context, channel, j); err != nil {
		s.ErrChan <- err
	}

}
