package pkg

type Behaviour interface {
	Process(b *Bot, state *GameState) error
}
