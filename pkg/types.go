package pkg

import "encoding/json"

type Event struct {
	Type string          `json:"e"`
	Data json.RawMessage `json:"data"`
}

func (e *Event) UnmarlshalData(out interface{}) error {
	return json.Unmarshal(e.Data, out)
}

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type PlayerInfo struct {
	Id       uint    `json:"id"`
	Angle    float32 `json:"angle"`
	Throttle float32 `json:"throttle"`
	Position
}

type BulletInfo struct {
	Id       uint    `json:"id"`
	Angle    float32 `json:"angle"`
	PlayerId uint    `json:"player_id"`
	Position
}

type GameState struct {
	Bounds     [2]float32    `json:"bounds"`
	Players    []*PlayerInfo `json:"players"`
	Bullets    []*BulletInfo `json:"bullets"`
	ScoreBoard map[uint]uint `json:"scoreboard"`
}
