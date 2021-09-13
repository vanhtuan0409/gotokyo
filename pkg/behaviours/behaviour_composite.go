package behaviours

import (
	"context"

	"github.com/vanhtuan0409/gotokyo/pkg"
)

type BehaviourComposite struct {
	ctx *pkg.Context

	movingStrat MovingStrategy
	shootStrat  ShootingStrategy
}

func NewBehaviourComposite(
	movingStrat MovingStrategy,
	shootStrat ShootingStrategy,
) *BehaviourComposite {
	return &BehaviourComposite{
		ctx: pkg.NewContext(),

		movingStrat: movingStrat,
		shootStrat:  shootStrat,
	}
}

func (c *BehaviourComposite) Process(tick uint64, b *pkg.Bot, state *pkg.GameState) error {
	c.ctx.Sync(b, state)
	ctx := context.Background()

	// shooting
	target, found := c.shootStrat.Scan(tick, c.ctx, b)
	if found {
		c.shoot(ctx, b, target)
		return nil
	}

	// moving
	vs := c.movingStrat.Process(tick, c.ctx, b)
	c.move(ctx, b, vs)

	return nil
}

func (c *BehaviourComposite) move(ctx context.Context, b *pkg.Bot, vs []pkg.Vector) error {
	if vs == nil || len(vs) == 0 {
		return nil
	}

	aggregated := pkg.AddAll(vs...)
	b.AdjustSpeed(ctx, 1)
	b.RotateVector(ctx, aggregated)
	return nil
}

func (c *BehaviourComposite) shoot(ctx context.Context, b *pkg.Bot, target pkg.Object) error {
	b.AdjustSpeed(ctx, 0)
	b.FaceToward(ctx, target)
	b.Fire(ctx)
	return nil
}
