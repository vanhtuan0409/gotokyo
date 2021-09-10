package pkg

import (
	"context"
	"math"
	"sync"
	"sync/atomic"
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

func (b *Bot) LastPosition() Position {
	return b.lastPosition
}

func (b *Bot) LastSpeed() float32 {
	return b.lastSpeed
}

func (b *Bot) Fire(ctx context.Context) error {
	return b.client.SendCommand(ctx, NewFireCommand())
}

func (b *Bot) AdjustSpeed(ctx context.Context, speed float32) error {
	b.setSpeed(speed)
	return b.client.SendCommand(ctx, NewThrottleCommand(b.lastSpeed))
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

func (b *Bot) setAngle(angle float32) float32 {
	if angle > (2 * math.Pi) {
		angle -= 2 * math.Pi
	}
	b.lastAngle = angle
	return b.lastAngle
}

func (b *Bot) setSpeed(speed float32) float32 {
	if speed > 1 {
		speed = 1
	}
	if speed < 0 {
		speed = 0
	}
	b.lastSpeed = speed
	return b.lastSpeed
}

func degreeToRad(degree int) float32 {
	ret := math.Pi / 180 * float32(degree)
	return ret
}

func (b *Bot) setProcessing(status bool) {
	val := uint32(0)
	if status {
		val = 1
	}
	atomic.StoreUint32(&b.processing, val)
}

func (b *Bot) isProcessing() bool {
	return atomic.LoadUint32(&b.processing) == 1
}

func (b *Bot) syncInfo(info *PlayerInfo) {
	if info == nil {
		return
	}

	b.lastAngle = info.Angle
	b.lastSpeed = info.Throttle
	b.lastPosition = info.Position
}
