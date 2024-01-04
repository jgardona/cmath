package cmath

import "math"

type Vec4 struct {
	x float64
	y float64
	z float64
	w float64
}

func NewVec4(x, y, z, w float64) Vec4 {
	return Vec4{x, y, z, w}
}

func (v Vec4) X() float64 {
	return v.x
}

func (v Vec4) Min() float64 {
	a := math.Min(v.x, v.y)
	b := math.Min(v.z, v.w)
	return math.Min(a, b)
}

func (v Vec4) Max() float64 {
	a := math.Max(v.x, v.y)
	b := math.Max(v.z, v.w)
	return math.Max(a, b)
}

func (v Vec4) Y() float64 {
	return v.y
}

func (v Vec4) Z() float64 {
	return v.z
}

func (v Vec4) W() float64 {
	return v.w
}
