package cmath

import (
	"math"

	"github.com/jgardona/cmath/constraints"
)

// Delta is used for compare floats. The epsilon constant
// used is 0.0001 decimal precision.
func Delta(a, b float64) bool {
	epsilon := 0.0001
	delta := math.Abs(a-b) < epsilon

	return delta
}

// Pop pops the latest element from a slice.
func Pop[T ~string | ~int | float64](s *[]T) T {
	if len(*s) == 0 {
		var zero T
		return zero
	}
	poppedValue := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return poppedValue
}

// Push pushes an element to the end of a slice.
func Push[T ~string | ~int | float64](s *[]T, v T) {
	*s = append(*s, v)
}

// Sum return the sum of all slice elements.
func Sum[T constraints.Numbers](data ...T) T {
	var total T
	for _, e := range data {
		total += e
	}
	return total
}

// Prod return the product of all slice elements.
func Prod[T constraints.Numbers](data ...T) T {
	var total T = 1
	for _, e := range data {
		total *= e
	}
	return total
}
