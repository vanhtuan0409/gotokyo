package pkg

import (
	"encoding/json"
)

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

func (p Position) GetType() ObjectType {
	return ObjectPosition
}

func (p Position) GetPosition() Position {
	return p
}

type PlayerInfo struct {
	Id       uint    `json:"id"`
	Angle    float32 `json:"angle"`
	Throttle float32 `json:"throttle"`
	Position
}

func (p PlayerInfo) GetType() ObjectType {
	return ObjectPlayer
}

func (p PlayerInfo) GetAngle() float32 {
	return p.Angle
}

func (p PlayerInfo) GetSpeed() float32 {
	return p.Throttle
}

type BulletInfo struct {
	Id       uint    `json:"id"`
	Angle    float32 `json:"angle"`
	PlayerId uint    `json:"player_id"`
	Position
}

func (b BulletInfo) GetType() ObjectType {
	return ObjectBullet
}

func (b BulletInfo) GetAngle() float32 {
	return b.Angle
}

func (b BulletInfo) GetSpeed() float32 {
	return BulletSpeed
}

type DeadInfo struct {
	Player *PlayerInfo `json:"player"`
}

type GameState struct {
	Bounds     [2]float32    `json:"bounds"`
	Players    []*PlayerInfo `json:"players"`
	Bullets    []*BulletInfo `json:"bullets"`
	Deads      []*DeadInfo   `json:"dead"`
	ScoreBoard map[uint]uint `json:"scoreboard"`
}
