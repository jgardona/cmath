// This package contains a structure and methods for representing
// a pair of coordinates of integer and float types.
package points

import (
	"math"

	"github.com/jgardona/cmath/constraints"
)

// A Point structure for representing a pair of coordinates
// of integer and float types
type Point[T constraints.Numbers] struct {
	x T
	y T
}

// The NewPoint function returns a float or integer Point type,
// given its x and y coordinates.
func NewPoint[T constraints.Numbers](x, y T) Point[T] {
	return Point[T]{x, y}
}

// Sum is the sum operation for a Point[T] type.
func (p Point[T]) Sum(point Point[T]) Point[T] {
	nx := p.x + point.x
	ny := p.y + point.y
	return NewPoint(nx, ny)
}

// SumScalar is the sum with a scalar operation for a Point[T] type.
func (p Point[T]) SumScalar(scalar T) Point[T] {
	nx := p.x + scalar
	ny := p.y + scalar
	return NewPoint(nx, ny)
}

// Subtract is subtraction operation for a Point[T] type.
func (p Point[T]) Subtract(point Point[T]) Point[T] {
	nx := p.x - point.x
	ny := p.y - point.y
	return NewPoint(nx, ny)
}

// SubScalar is subtraction with a scalar operation for a Point[T] type.
func (p Point[T]) SubScalar(scalar T) Point[T] {
	nx := p.x - scalar
	ny := p.y - scalar
	return NewPoint(nx, ny)
}

// Multiply operation with a scalar for a Point[T] type.
func (p Point[T]) Multiply(factor T) Point[T] {
	nx := p.x * factor
	ny := p.y * factor
	return NewPoint(nx, ny)
}

// Divide is division operation with a scalar for a Point[T] type.
func (p Point[T]) Divide(factor T) Point[T] {
	if factor == 0.0 {
		panic("cant divide by zero")
	}
	nx := p.x / factor
	ny := p.y / factor
	return NewPoint(nx, ny)
}

// DistanceTo calculates the euclidean distance between two points.
func (p Point[T]) DistanceTo(point Point[T]) float64 {
	dx := p.x - point.x
	dy := p.y - point.y
	return math.Sqrt(float64(dx*dx) + float64(dy*dy))
}

// SquaredDistanceTo calculates the squared Euclidean distance between two points.
func (p Point[T]) SquaredDistanceTo(point Point[T]) float64 {
	dx := p.x - point.x
	dy := p.y - point.y
	return float64(dx*dx) + float64(dy*dy)
}

// X returns the T value from x coordinate.
func (p Point[T]) X() T {
	return p.x
}

// Y returns the T value from y coordinate.
func (p Point[T]) Y() T {
	return p.y
}

// Equals verifies if point is equals to a. For a more accurate
// comparison, the Delta(a, b) function is recommended.
func (p Point[T]) Equals(a Point[T]) bool {
	return p.x == a.x && p.y == a.y
}
