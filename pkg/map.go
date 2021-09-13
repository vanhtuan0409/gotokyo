package pkg

import (
	"math"
)

type Map struct {
	Bounds  [2]float32
	Center  Position
	Conners [4]Position
}

func NewEmptyMap() *Map {
	m := &Map{
		Bounds: [2]float32{0, 0},
	}
	m.Initialize()
	return m
}

func (m *Map) Initialize() {
	// identify center
	m.Center = Position{
		X: float32(math.Round(float64(m.Bounds[0]) / 2)),
		Y: float32(math.Round(float64(m.Bounds[1]) / 2)),
	}

	// identify conners
	m.Conners = [4]Position{}
	m.Conners[0] = Position{X: 0, Y: 0}                     // top left
	m.Conners[1] = Position{X: m.Bounds[0], Y: 0}           // top right
	m.Conners[2] = Position{X: 0, Y: m.Bounds[1]}           // bot left
	m.Conners[3] = Position{X: m.Bounds[0], Y: m.Bounds[1]} // bot right
}

func (m *Map) UpdateBound(bounds [2]float32) {
	if m.Bounds[0] != bounds[0] || m.Bounds[1] != bounds[1] {
		m.Bounds = bounds
		m.Initialize()
	}
}

func (m *Map) Diagonal() float32 {
	ret := math.Sqrt(math.Hypot(
		float64(m.Bounds[0]),
		float64(m.Bounds[1]),
	))
	return float32(ret)
}

func (m *Map) MaxX() float32 {
	return m.Bounds[0]
}

func (m *Map) MaxY() float32 {
	return m.Bounds[1]
}
