package pkg

import "context"

type Behaviour interface {
	Process(tick uint64, b *Bot, state *GameState) error
}

type BehaviourFunc func(tick uint64, b *Bot, state *GameState) error

func (f BehaviourFunc) Process(tick uint64, b *Bot, state *GameState) error {
	return f(tick, b, state)
}

var (
	StandStillBehaviour = BehaviourFunc(func(tick uint64, b *Bot, state *GameState) error {
		b.AdjustSpeed(context.Background(), 0)
		return nil
	})
)
