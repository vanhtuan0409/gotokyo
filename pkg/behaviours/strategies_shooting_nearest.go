package behaviours

import (
	"github.com/vanhtuan0409/gotokyo/pkg"
)

var (
	StratShootNearest = ShootingStratFunc(func(tick uint64, ctx *pkg.Context, b *pkg.Bot) (pkg.Object, bool) {
		target := ScanNearestPlayer(ctx)
		return target, target != nil
	})
)
