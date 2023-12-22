package cmath

import "math"

// Delta is used for compare floats. The epsilon constant
// used is 0.0001 decimal precision.
func Delta(a, b float64) bool {
	epsilon := 0.0001
	delta := math.Abs(a-b) < epsilon

	return delta
}
