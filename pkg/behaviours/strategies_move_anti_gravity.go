package behaviours

import (
	"github.com/vanhtuan0409/gotokyo/pkg"
)

var (
	// StratMoveAntiGravity should handle move to avoid player outside of DangerRadius
	// For moving object within DangerRadius, should refer to DodgePerpendicular
	StratMoveAntiGravity = MovingStratFunc(func(tick uint64, ctx *pkg.Context, b *pkg.Bot) []pkg.Vector {
		ret := []pkg.Vector{}

		// other players anti gravity
		for _, p := range ctx.SortedAlivePlayers {
			if ctx.GetPlayerDistance(p.Id) > DangerRadius {
				ev := calcGravityVector(b, p)
				ret = append(ret, ev)
			}
		}

		// wall anti gravity
		topEv := calcGravityFoce(pkg.NewVector(0, b.GetPosition().Y))
		ret = append(ret, topEv) // top
		downEv := calcGravityFoce(pkg.NewVector(0, b.GetPosition().Y-ctx.Map.MaxY()))
		ret = append(ret, downEv) // down
		leftEv := calcGravityFoce(pkg.NewVector(b.GetPosition().X, 0))
		ret = append(ret, leftEv) // left
		rightEv := calcGravityFoce(pkg.NewVector(b.GetPosition().X-ctx.Map.MaxX(), 0))
		ret = append(ret, rightEv) // right

		return ret
	})
)

func calcGravityVector(b *pkg.Bot, obj pkg.Object) pkg.Vector {
	v := pkg.
		NewVectorFromPoint(b.GetPosition(), obj.GetPosition()).
		Inverse()
	return calcGravityFoce(v)
}

func calcGravityFoce(v pkg.Vector) pkg.Vector {
	mag := v.Mag()
	v = v.Div(mag * mag)
	return v
}
