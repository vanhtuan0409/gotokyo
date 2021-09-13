package behaviours

import (
	"github.com/vanhtuan0409/gotokyo/pkg"
)

type MovingStrategy interface {
	Process(tick uint64, ctx *pkg.Context, b *pkg.Bot) []pkg.Vector
}

type MovingStratFunc func(tick uint64, ctx *pkg.Context, b *pkg.Bot) []pkg.Vector

func (s MovingStratFunc) Process(tick uint64, ctx *pkg.Context, b *pkg.Bot) []pkg.Vector {
	return s(tick, ctx, b)
}

func CombineMovingStrat(strats ...MovingStrategy) MovingStrategy {
	return MovingStratFunc(func(tick uint64, ctx *pkg.Context, b *pkg.Bot) []pkg.Vector {
		ret := []pkg.Vector{}
		for _, strat := range strats {
			ret = append(ret, strat.Process(tick, ctx, b)...)
		}
		return ret
	})
}

type ShootingStrategy interface {
	Scan(tick uint64, ctx *pkg.Context, b *pkg.Bot) (pkg.Object, bool)
}

type ShootingStratFunc func(tick uint64, ctx *pkg.Context, b *pkg.Bot) (pkg.Object, bool)

func (s ShootingStratFunc) Scan(tick uint64, ctx *pkg.Context, b *pkg.Bot) (pkg.Object, bool) {
	return s(tick, ctx, b)
}
