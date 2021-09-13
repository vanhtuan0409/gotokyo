package pkg

import "sort"

type Context struct {
	Map                *Map
	SortedAlivePlayers []*PlayerInfo
	SortedBullets      []*BulletInfo
	ScoreBoard         map[uint]uint

	PlayerDistances map[uint]float32
	BulletDistances map[uint]float32
}

func NewContext() *Context {
	return &Context{
		Map:                NewEmptyMap(),
		SortedAlivePlayers: []*PlayerInfo{},
		SortedBullets:      []*BulletInfo{},
		ScoreBoard:         map[uint]uint{},

		PlayerDistances: map[uint]float32{},
		BulletDistances: map[uint]float32{},
	}
}

func (c *Context) Sync(b *Bot, state *GameState) {
	c.Map.UpdateBound(state.Bounds)
	c.ScoreBoard = state.ScoreBoard
	c.syncPlayerDistances(b, state)
	c.syncBulletDistances(b, state)
}

func (c *Context) GetPlayerDistance(id uint) float32 {
	return c.PlayerDistances[id]
}

func (c *Context) GetBulletDistance(id uint) float32 {
	return c.BulletDistances[id]
}

func (c *Context) GetMoverInRange(r float32) []Mover {
	ret := []Mover{}
	for _, p := range c.SortedAlivePlayers {
		distance := c.GetPlayerDistance(p.Id)
		if distance <= r {
			ret = append(ret, p)
		}
	}
	for _, b := range c.SortedBullets {
		distance := c.GetBulletDistance(b.Id)
		if distance <= r {
			ret = append(ret, b)
		}
	}
	return ret
}

func (c *Context) syncPlayerDistances(b *Bot, state *GameState) {
	c.PlayerDistances = map[uint]float32{}
	for _, p := range state.Players {
		if p.Id != b.ID() {
			v := NewVectorFromPoint(b.GetPosition(), p.GetPosition())
			c.PlayerDistances[p.Id] = v.Mag()
		}
	}
	c.syncPlayerList(b, state)
}

func (c *Context) syncBulletDistances(b *Bot, state *GameState) {
	c.BulletDistances = map[uint]float32{}
	for _, bu := range state.Bullets {
		v := NewVectorFromPoint(b.GetPosition(), bu.GetPosition())
		c.BulletDistances[bu.Id] = v.Mag()
	}
	c.syncBulletList(state)
}

func (c *Context) syncPlayerList(b *Bot, state *GameState) {
	c.SortedAlivePlayers = []*PlayerInfo{}

	deadPlayers := map[uint]bool{}
	for _, p := range state.Deads {
		deadPlayers[p.Player.Id] = true
	}

	for _, p := range state.Players {
		if !deadPlayers[p.Id] && (p.Id != b.ID()) {
			c.SortedAlivePlayers = append(c.SortedAlivePlayers, p)
		}
	}

	sort.Slice(c.SortedAlivePlayers, func(i, j int) bool {
		player1 := c.SortedAlivePlayers[i]
		player2 := c.SortedAlivePlayers[j]
		return c.PlayerDistances[player1.Id] < c.PlayerDistances[player2.Id]
	})
}

func (c *Context) syncBulletList(state *GameState) {
	c.SortedBullets = state.Bullets
	sort.Slice(c.SortedBullets, func(i, j int) bool {
		bullet1 := c.SortedBullets[i]
		bullet2 := c.SortedBullets[j]
		return c.BulletDistances[bullet1.Id] < c.BulletDistances[bullet2.Id]
	})
}
