package pkg

import (
	"math"
)

type Vector [2]float32

func NewVector(x, y float32) Vector {
	return [2]float32{x, y}
}

func NewVectorFromPoint(root, head Position) Vector {
	x := head.X - root.X
	y := head.Y - root.Y
	return NewVector(x, y)
}

func NewVectorFromAngle(root Position, angle float32, mag float32) Vector {
	v := NewVector(
		root.X+float32(math.Cos(float64(angle))),
		root.Y+float32(math.Sin(float64(angle))),
	)
	v = v.Mul(mag / v.Mag())
	return v
}

func (v Vector) X() float32 {
	return v[0]
}

func (v Vector) Y() float32 {
	return v[1]
}

func (v Vector) Add(v2 Vector) Vector {
	return NewVector(
		v.X()+v2.X(),
		v.Y()+v2.Y(),
	)
}

func (v Vector) Sub(v2 Vector) Vector {
	return NewVector(
		v.X()-v2.X(),
		v.Y()-v2.Y(),
	)
}

func (v Vector) Mul(val float32) Vector {
	return NewVector(
		v.X()*val,
		v.Y()*val,
	)
}

func (v Vector) Div(val float32) Vector {
	return NewVector(
		v.X()/val,
		v.Y()/val,
	)
}

func (v Vector) Dot(v2 Vector) Vector {
	return NewVector(
		v.X()*v2.X(),
		v.Y()*v2.Y(),
	)
}

func (v Vector) MagSquared() float32 {
	return v.Mag() * v.Mag()
}

func (v Vector) maq64() float64 {
	return math.Hypot(
		float64(v.X()),
		float64(v.Y()),
	)
}

func (v Vector) Mag() float32 {
	return float32(v.maq64())
}

func (v Vector) Normalize() Vector {
	return v.Div(v.Mag())
}

func (v Vector) Inverse() Vector {
	return NewVector(
		-v.X(),
		-v.Y(),
	)
}

// Heading return angle in rad of vector
func (v Vector) Heading() float32 {
	return float32(math.Atan2(
		float64(v.Y()),
		float64(v.X()),
	))
}

func AddAll(vs ...Vector) Vector {
	if len(vs) == 0 {
		return NewVector(0, 0)
	}
	if len(vs) == 1 {
		return vs[0]
	}

	ret := vs[0]
	for _, v := range vs[1:] {
		ret = ret.Add(v)
	}

	return ret
}

func AngleBetween(v1 Vector, v2 Vector) float32 {
	angle := math.Atan2(
		float64(v2.Y()),
		float64(v2.X()),
	) - math.Atan2(
		float64(v1.Y()),
		float64(v1.X()),
	)
	return NormalizeRad(float32(angle))
}
