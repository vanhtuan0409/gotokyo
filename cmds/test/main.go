package main

import (
	"fmt"

	"github.com/vanhtuan0409/gotokyo/pkg"
)

func main() {
	pos := pkg.Position{X: 0, Y: 0}
	v1 := pkg.NewVectorFromAngle(pos, pkg.DegreeToRad(50), 1)
	v2 := pkg.NewVectorFromAngle(pos, pkg.DegreeToRad(60), 1)
	angle := pkg.AngleBetween(v1, v2)
	fmt.Println(pkg.RadToDegree(angle))
}
