package behaviours

import "github.com/vanhtuan0409/gotokyo/pkg"

var (
	StratMoveCorner = MovingStratFunc(func(tick uint64, ctx *pkg.Context, b *pkg.Bot) []pkg.Vector {
		corner := ctx.Map.Conners[0]
		v := pkg.NewVectorFromPoint(b.GetPosition(), corner)
		return []pkg.Vector{v}
	})
)
