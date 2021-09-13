package pkg

import (
	"context"
	"fmt"
	"log"
)

type GameLoop struct {
	tick uint64
}

func NewGameLoop() *GameLoop {
	return &GameLoop{}
}

func (l *GameLoop) Tick() uint64 {
	return l.tick
}

func (l *GameLoop) Run(ctx context.Context, b *Bot, source *EventSource) {
	stream := source.Stream(ctx)
	initEvent := <-stream
	if err := initEvent.UnmarlshalData(&b.id); err != nil {
		panic(fmt.Errorf("Cannot init bot id. ERR: %v", err))
	}
	log.Printf("Bot initialized successfully. Bot id %d", b.id)

	var state GameState
	for true {
		// discard outdated event, keep the latest
		eventLength := len(stream)
		if eventLength > 1 {
			for i := 0; i < eventLength-1; i++ {
				l.readEvent(stream)
			}
		}
		event, ok := l.readEvent(stream)
		if !ok {
			break
		}

		// unmarshal state
		if err := event.UnmarlshalData(&state); err != nil {
			log.Printf("Unable to parse game state. ERR: %+v", err)
			continue
		}

		// resync with data from server
		botInfo := FindPlayer(b.id, state.Players)
		b.syncInfo(botInfo)
		if score, ok := state.ScoreBoard[b.id]; ok {
			b.lastScore = score
		}

		// perform behaviour processing
		if err := b.behaviour.Process(l.tick, b, &state); err != nil {
			log.Printf("Bot's behaviour cannot process state. ERR: %+v", err)
		}
	}
}

func (l *GameLoop) readEvent(stream <-chan *Event) (*Event, bool) {
	e, ok := <-stream
	l.doTick()
	return e, ok
}

func (l *GameLoop) doTick() {
	l.tick += 1
	// reset tick if too large
	if l.tick >= 10_000_000 {
		l.tick = 0
	}
}
