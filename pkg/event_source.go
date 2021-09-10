package pkg

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/gorilla/websocket"
)

type EventSource struct {
	conn  *websocket.Conn
	debug bool
}

func NewEventSource(conn *websocket.Conn, debug bool) *EventSource {
	return &EventSource{
		conn:  conn,
		debug: debug,
	}
}

func (s *EventSource) Stream(ctx context.Context, bufferSize uint) <-chan *Event {
	ch := make(chan *Event, bufferSize)
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				event, err := s.readEvent()
				if err != nil {
					// TODO: try to reconnect
					return
				}

				s.dispatch(ch, event)
			}
		}
	}()
	return ch
}

func (s *EventSource) readEvent() (e *Event, err error) {
	_, r, err := s.conn.NextReader()
	if err != nil {
		return nil, err
	}
	if s.debug {
		r = io.TeeReader(r, os.Stdout)
	}

	err = json.NewDecoder(r).Decode(e)
	return
}

func (s *EventSource) dispatch(ch chan<- *Event, e *Event) {
	// drop event if cannot handle
	select {
	case ch <- e:
	default:
	}
}
