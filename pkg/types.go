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
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type PlayerPosition struct {
	Id       uint    `json:"id"`
	Angle    float64 `json:"angle"`
	Throttle float64 `json:"throttle"`
	Position
}

type BulletPosition struct {
	Id       uint    `json:"id"`
	Angle    float64 `json:"angle"`
	PlayerId uint    `json:"player_id"`
	Position
}

type GameState struct {
	Bounds     [2]float64        `json:"bounds"`
	Players    []*PlayerPosition `json:"players"`
	Bullets    []*BulletPosition `json:"bullets"`
	ScoreBoard map[uint]int      `json:"scoreboard"`
}
