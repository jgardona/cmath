package points

import "math"

type Numbers interface {
	int | float64
}

type Point[T Numbers] struct {
	x T
	y T
}

func NewPoint[T Numbers](x, y T) Point[T] {
	return Point[T]{x, y}
}

func (p Point[T]) Sum(point Point[T]) Point[T] {
	nx := p.x + point.x
	ny := p.y + point.y
	return NewPoint(nx, ny)
}

func (p Point[T]) SumScalar(scalar T) Point[T] {
	nx := p.x + scalar
	ny := p.y + scalar
	return NewPoint(nx, ny)
}

func (p Point[T]) Subtract(point Point[T]) Point[T] {
	nx := p.x - point.x
	ny := p.y - point.y
	return NewPoint(nx, ny)
}

func (p Point[T]) SubScalar(scalar T) Point[T] {
	nx := p.x - scalar
	ny := p.y - scalar
	return NewPoint(nx, ny)
}

func (p Point[T]) Multiply(factor T) Point[T] {
	nx := p.x * factor
	ny := p.y * factor
	return NewPoint(nx, ny)
}

func (p Point[T]) Divide(factor T) Point[T] {
	if factor == 0.0 {
		panic("cant divide by zero")
	}
	nx := p.x / factor
	ny := p.y / factor
	return NewPoint(nx, ny)
}

func (p Point[T]) DistanceTo(point Point[T]) float64 {
	dx := p.x - point.x
	dy := p.y - point.y
	return math.Sqrt(float64(dx*dx) + float64(dy*dy))
}

func (p Point[T]) SquaredDistanceTo(point Point[T]) float64 {
	dx := p.x - point.x
	dy := p.y - point.y
	return float64(dx*dx) + float64(dy*dy)
}

func (p Point[T]) X() T {
	return p.x
}

func (p Point[T]) Y() T {
	return p.y
}

func Delta(a, b Point[float64]) bool {
	epsilon := 0.0001
	dx := math.Abs(a.x-b.y) < epsilon
	dy := math.Abs(a.y-b.y) < epsilon
	return dx && dy
}
