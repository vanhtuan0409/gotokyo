package pkg

import (
	"context"
	"log"

	"golang.org/x/time/rate"
)

type EventSource struct {
	client  *Client
	limiter *rate.Limiter
	opt     *SourceOpt
}

type SourceOpt struct {
	BufferSize uint
	Rate       uint // number of event sample per second
}

func NewEventSource(client *Client, opt SourceOpt) *EventSource {
	r := rate.Inf
	if opt.Rate > 0 {
		r = rate.Limit(opt.Rate)
	}

	s := &EventSource{
		client:  client,
		opt:     &opt,
		limiter: rate.NewLimiter(r, 1),
	}
	return s
}

func (s *EventSource) Stream(ctx context.Context) <-chan *Event {
	ch := make(chan *Event, s.opt.BufferSize)
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				event, err := s.client.ReadEvent(ctx)
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

func (s *EventSource) dispatch(ch chan<- *Event, e *Event) {
	// limit exceed
	if !s.limiter.Allow() {
		log.Println("Event rate exceed rate limit. Drop event")
		return
	}

	select {
	case ch <- e:
	default:
		log.Println("Event buffer overflow. Drop event")
	}
}
