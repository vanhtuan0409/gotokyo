package pkg

type Behaviour interface {
	Process(tick uint64, b *Bot, state *GameState) error
}
