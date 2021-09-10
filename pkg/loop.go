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
	for event := range stream {
		if err := event.UnmarlshalData(&state); err != nil {
			log.Printf("Unable to parse game state. ERR: %+v", err)
			continue
		}

		l.doTick()

		// There will be 2 thread accessing bot instance
		// 1 is main game loop thread
		// 2 is behaviour processing thread
		// Behaviour processing might run slower than the received event rate
		// Game loop should always ensure bot's behaviour based on latest received event
		go func(b *Bot, state *GameState) {
			if b.isProcessing() {
				return
			}

			// resync with data from server
			botInfo := FindPlayer(b.id, state.Players)
			b.syncInfo(botInfo)
			if score, ok := state.ScoreBoard[b.id]; ok {
				b.lastScore = score
			}

			// perform behaviour processing
			b.setProcessing(true)
			defer b.setProcessing(false)
			b.behaviour.Process(l.tick, b, state)
		}(b, &state)
	}

}

func (l *GameLoop) doTick() {
	l.tick += 1
	// reset tick if too large
	if l.tick >= 10_000_000 {
		l.tick = 0
	}
}
