package behaviours

import "github.com/vanhtuan0409/gotokyo/pkg"

func ScanNearestPlayer(ctx *pkg.Context) *pkg.PlayerInfo {
	if len(ctx.SortedAlivePlayers) < 1 {
		return nil
	}
	return ctx.SortedAlivePlayers[0]
}

func ScanNearestPlayerWithin(ctx *pkg.Context, distance float32) *pkg.PlayerInfo {
	found := ScanNearestPlayer(ctx)
	if found != nil && ctx.GetPlayerDistance(found.Id) < distance {
		return found
	}
	return nil
}
