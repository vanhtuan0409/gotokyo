package analysis

import (
	"github.com/vanhtuan0409/gotokyo/pkg"
)

type FacingResult struct {
	FacingVector pkg.Vector
	MovingVector pkg.Vector
	IsFacing     bool
	VectorAngle  float32
}

func AnalyzeFacing(mover pkg.Mover, target pkg.Object, maxAngle float32) *FacingResult {
	facingV := pkg.NewVectorFromPoint(mover.GetPosition(), target.GetPosition())
	movingV := pkg.NewVectorFromAngle(mover.GetPosition(), mover.GetAngle(), facingV.Mag())
	vectorAngle := pkg.AngleBetween(facingV, movingV)
	isFacing := vectorAngle <= maxAngle
	return &FacingResult{
		FacingVector: facingV,
		MovingVector: movingV,
		IsFacing:     isFacing,
		VectorAngle:  vectorAngle,
	}
}
