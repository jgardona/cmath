package cmath

import (
	"math"
)

type Vec3 struct {
	x float64
	y float64
	z float64
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

func (v Vec3) Max() float64 {
	return math.Max(v.x, math.Max(v.y, v.z))
}

func (v Vec3) Min() float64 {
	return math.Min(v.x, math.Min(v.y, v.z))
}

func (v Vec3) MaxIndex() int {
	if v.x >= v.y {
		if v.x >= v.z {
			return 0
		} else {
			return 2
		}
	} else {
		if v.y >= v.z {
			return 1
		} else {
			return 2
		}
	}
}

func (v Vec3) MinIndex() int {
	if v.x <= v.y {
		if v.x <= v.z {
			return 0
		} else {
			return 2
		}
	} else {
		if v.y <= v.z {
			return 1
		} else {
			return 2
		}
	}
}

func (v Vec3) Norm() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Vec3) Dot() float64 {
	panic("not yet implemented")
}

func (v Vec3) AsArray() []float64 {
	panic("not yet implemented")
}

func (v Vec3) Sum(vec Vec3) Vec3 {
	panic("not yet implemented")
}

func (v Vec3) SumScalar(scalar float64) Vec3 {
	panic("not yet implemented")
}

func (v Vec3) Subtract(vec Vec3) Vec3 {
	panic("not yet implemented")
}

func (v Vec3) SubScalar(scalar float64) Vec3 {
	panic("not yet implemented")
}

func (v Vec3) Multiply(vec Vec3) Vec3 {
	panic("not yet implemented")
}

func (v Vec3) MulScalar(scalar float64) Vec3 {
	panic("not yet implemented")
}

func (v Vec3) Divide(vec Vec3) Vec3 {
	panic("not yet implemented")
}

func (v Vec3) DivScalar(scalar float64) Vec3 {
	panic("not yet implemented")
}

func (v Vec3) Equals(vec Vec3) bool {
	panic("not yet implemented")
}

func (v Vec3) Normalize() float64 {
	panic("not yet implemented")
}

func (v Vec3) Inverse() Vec3 {
	panic("not yet implemented")
}

func (v Vec3) Abs() Vec3 {
	panic("not yet implemented")
}

func (v Vec3) Cross(a, b, c Vec3) Vec3 {
	panic("not yet implemented")
}

func Dot(a, b Vec3) float64 {
	panic("not yet implemented")
}

func (v Vec3) AsVec4() Vec4 {
	panic("not yet implemented")
}
