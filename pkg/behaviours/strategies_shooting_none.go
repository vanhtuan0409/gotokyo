package behaviours

import (
	"github.com/vanhtuan0409/gotokyo/pkg"
)

var (
	StratShootNone = ShootingStratFunc(func(tick uint64, ctx *pkg.Context, b *pkg.Bot) (pkg.Object, bool) {
		return nil, false
	})
)
