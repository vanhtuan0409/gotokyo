package behaviours

import (
	"math/rand"

	"github.com/vanhtuan0409/gotokyo/pkg"
	"github.com/vanhtuan0409/gotokyo/pkg/analysis"
)

var (
	// StratMoveDodgePerpendicular should handle moving object within DangerRadius
	StratMoveDodgePerpendicular = MovingStratFunc(func(tick uint64, ctx *pkg.Context, b *pkg.Bot) []pkg.Vector {
		ret := []pkg.Vector{}
		movers := ctx.GetMoverInRange(DangerRadius)
		for _, m := range movers {
			result := analysis.AnalyzeFacing(m, b, FacingMaxAngle)

			if result.IsFacing {
				// if is facing. Move perpendicular
				clockwise := rand.Intn(100)%2 == 0
				v := getPerpenVec(result.MovingVector, clockwise)
				v = calcGravityFoce(v)
				v = v.Mul(DodgingFactor)
				ret = append(ret, v)
			} else {
				// if not, just move anti gravity
				v := calcGravityVector(b, m.GetPosition())
				ret = append(ret, v)
			}
		}
		return ret
	})
)

func getPerpenVec(v pkg.Vector, clockwise bool) pkg.Vector {
	if clockwise {
		return pkg.NewVector(v.Y(), -v.X())
	}
	return pkg.NewVector(-v.Y(), v.X())
}
