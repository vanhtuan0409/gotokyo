package pkg

import (
	"context"
	"log"
	"math"
)

const (
	maxBullets = 4
)

type Bot struct {
	id        uint
	name      string
	lastAngle float32
	behaviour Behaviour

	client *Client
}

func NewBot(name string, behaviour Behaviour, client *Client) *Bot {
	b := Bot{
		name:      name,
		client:    client,
		behaviour: behaviour,
	}
	return &b
}

func (b *Bot) ID() uint {
	return b.id
}

func (b *Bot) Name() string {
	return b.name
}

func (b *Bot) LastAngle() float32 {
	return b.lastAngle
}

func (b *Bot) Run(ctx context.Context, source *EventSource) error {
	stream := source.Stream(ctx)
	if err := b.init(ctx, stream); err != nil {
		return err
	}
	log.Printf("Bot initialized successfully. Bot id %d", b.id)

	var state GameState
	for event := range stream {
		if err := event.UnmarlshalData(&state); err != nil {
			break // invalid game state
		}
		b.behaviour.Process(b, &state)
	}

	return nil
}

func (b *Bot) init(ctx context.Context, stream <-chan *Event) error {
	e := <-stream
	e.UnmarlshalData(&b.id)
	return nil
}

func (b *Bot) Fire(ctx context.Context) error {
	return b.client.SendCommand(ctx, NewFireCommand())
}

func (b *Bot) AdjustSpeed(ctx context.Context, speed float32) error {
	return b.client.SendCommand(ctx, NewThrottleCommand(speed))
}

func (b *Bot) RotateAbs(ctx context.Context, angle float32) error {
	b.setAngle(angle)
	return b.client.SendCommand(ctx, NewRotateCommand(b.lastAngle))
}

func (b *Bot) Rotate(ctx context.Context, angle float32) error {
	b.setAngle(b.lastAngle + angle)
	return b.client.SendCommand(ctx, NewRotateCommand(b.lastAngle))
}

func (b *Bot) RotateDegAbs(ctx context.Context, degree int) error {
	return b.RotateAbs(ctx, degreeToRad(degree))
}

func (b *Bot) RotateDeg(ctx context.Context, degree int) error {
	return b.Rotate(ctx, degreeToRad(degree))
}

func (b *Bot) ChangeBehaviour(behaviour Behaviour) {
	b.behaviour = behaviour
}

func (b *Bot) setAngle(angle float32) {
	b.lastAngle = angle
	if b.lastAngle > (2 * math.Pi) {
		b.lastAngle -= 2 * math.Pi
	}
}

func degreeToRad(degree int) float32 {
	ret := math.Pi / 180 * float32(degree)
	return ret
}
