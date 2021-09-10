package behaviours

import (
	"context"

	"github.com/vanhtuan0409/gotokyo/pkg"
)

var (
	SampleBehaviour = pkg.BehaviourFunc(func(tick uint64, b *pkg.Bot, state *pkg.GameState) error {
		ctx := context.Background()
		worldMap := pkg.NewEmptyMap()
		worldMap.UpdateBound(state.Bounds)

		target := worldMap.Conners[0]
		b.AdjustSpeed(ctx, 1)
		b.FaceToward(ctx, target)

		return nil
	})
)
