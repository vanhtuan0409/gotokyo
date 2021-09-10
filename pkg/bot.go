package pkg

import (
	"context"
	"sync"
)

type Bot struct {
	id        uint
	name      string
	numBullet uint32

	lastPosition Position
	lastAngle    float32
	lastSpeed    float32
	lastScore    uint

	behaviour Behaviour

	client     *Client
	processing uint32
	lock       sync.Mutex
}

func NewBot(name string, behaviour Behaviour, client *Client) *Bot {
	b := Bot{
		name:      name,
		numBullet: BotMaxBullets,
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

func (b *Bot) GetPosition() Position {
	return b.lastPosition
}

func (b *Bot) LastSpeed() float32 {
	return b.lastSpeed
}

func (b *Bot) GetType() ObjectType {
	return ObjectBot
}

func (b *Bot) Fire(ctx context.Context) error {
	return b.client.SendCommand(ctx, NewFireCommand())
}

func (b *Bot) AdjustSpeed(ctx context.Context, speed float32) error {
	_, changed := b.setSpeed(speed)
	if !changed {
		return nil
	}
	return b.client.SendCommand(ctx, NewThrottleCommand(b.LastSpeed()))
}

func (b *Bot) RotateAbs(ctx context.Context, angle float32) error {
	b.setAngle(angle)
	return b.client.SendCommand(ctx, NewRotateCommand(b.LastAngle()))
}

func (b *Bot) Rotate(ctx context.Context, angle float32) error {
	b.setAngle(b.LastAngle() + angle)
	return b.client.SendCommand(ctx, NewRotateCommand(b.LastAngle()))
}

func (b *Bot) RotateDegAbs(ctx context.Context, degree int) error {
	return b.RotateAbs(ctx, DegreeToRad(degree))
}

func (b *Bot) RotateDeg(ctx context.Context, degree int) error {
	return b.Rotate(ctx, DegreeToRad(degree))
}

func (b *Bot) RotateVector(ctx context.Context, v Vector) error {
	angle := v.Heading()
	return b.RotateAbs(ctx, angle)
}

func (b *Bot) FaceToward(ctx context.Context, o Object) error {
	v := NewVectorFromPoint(b.GetPosition(), o.GetPosition())
	return b.RotateVector(ctx, v)
}

func (b *Bot) FaceAway(ctx context.Context, o Object) error {
	v := NewVectorFromPoint(b.GetPosition(), o.GetPosition()).Inverse()
	return b.RotateVector(ctx, v)
}

func (b *Bot) ChangeBehaviour(behaviour Behaviour) {
	b.behaviour = behaviour
}

func (b *Bot) setAngle(angle float32) float32 {
	b.lastAngle = NormalizeRad(angle)
	return b.lastAngle
}

func (b *Bot) setSpeed(speed float32) (float32, bool) {
	if speed > 1 {
		speed = 1
	}
	if speed < 0 {
		speed = 0
	}
	changed := b.lastSpeed != speed
	b.lastSpeed = speed
	return b.lastSpeed, changed
}

func (b *Bot) setProcessing(status bool) {
	b.lock.Lock()
	defer b.lock.Unlock()
	val := uint32(0)
	if status {
		val = 1
	}
	b.processing = val
}

func (b *Bot) isProcessing() bool {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.processing == 1
}

func (b *Bot) syncInfo(info *PlayerInfo) {
	if info == nil {
		return
	}

	b.setAngle(info.Angle)
	b.setSpeed(info.Throttle)
	b.lastPosition = info.Position
}
