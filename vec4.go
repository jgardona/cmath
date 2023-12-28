package cmath

type Vec4 struct {
	x float64
	y float64
	z float64
	w float64
}

func NewVec4(x, y, z, w float64) Vec4 {
	panic("not yet implemented")
}

func (v Vec4) X() float64 {
	return v.x
}

func (v Vec4) Min() float64 {
	panic("not yet implemented")
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
