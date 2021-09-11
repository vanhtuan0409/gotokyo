package behaviours

import (
	"github.com/vanhtuan0409/gotokyo/pkg"
)

type StratShootThrottle struct {
	lastFire uint64
	interval uint64
	inner    ShootingStrategy
}

func NewStratShootingThrottle(tickInterval uint64, inner ShootingStrategy) *StratShootThrottle {
	return &StratShootThrottle{
		lastFire: 0,
		interval: tickInterval,
		inner:    inner,
	}
}

func (s *StratShootThrottle) Scan(tick uint64, ctx *pkg.Context, b *pkg.Bot) (pkg.Object, bool) {
	// throttled, should not fire
	if !s.shouldFire(tick) {
		return nil, false
	}

	target, found := s.inner.Scan(tick, ctx, b)
	if found {
		s.lastFire = tick
	}
	return target, found
}

func (s *StratShootThrottle) shouldFire(tick uint64) bool {
	return tick-s.lastFire >= s.interval
}
