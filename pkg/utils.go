package pkg

import (
	"math"
)

func DegreeToRad(degree int) float32 {
	ret := math.Pi / 180 * float32(degree)
	return ret
}

func RadToDegree(val float32) int {
	ret := 180 * val / math.Pi
	return int(ret)
}

func NormalizeRad(val float32) float32 {
	for val <= 0 {
		val += 2 * math.Pi
	}
	for val >= (2 * math.Pi) {
		val -= 2 * math.Pi
	}
	return val
}
