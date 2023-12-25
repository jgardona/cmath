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
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vec3) AsArray() [3]float64 {
	return [3]float64{v.x, v.y, v.z}
}

func (v Vec3) Sum(vec Vec3) Vec3 {
	sx := v.x + vec.x
	sy := v.y + vec.y
	sz := v.z + vec.z

	return NewVec3(sx, sy, sz)
}

func (v Vec3) SumScalar(scalar float64) Vec3 {
	sx := v.x + scalar
	sy := v.y + scalar
	sz := v.z + scalar
	return NewVec3(sx, sy, sz)
}

func (v Vec3) Subtract(vec Vec3) Vec3 {
	sx := v.x - vec.x
	sy := v.y - vec.y
	sz := v.z - vec.z

	return NewVec3(sx, sy, sz)
}

func (v Vec3) SubScalar(scalar float64) Vec3 {
	sx := v.x - scalar
	sy := v.y - scalar
	sz := v.z - scalar
	return NewVec3(sx, sy, sz)
}

func (v Vec3) Multiply(vec Vec3) Vec3 {
	sx := v.x * vec.x
	sy := v.y * vec.y
	sz := v.z * vec.z
	return NewVec3(sx, sy, sz)
}

func (v Vec3) MulScalar(scalar float64) Vec3 {
	sx := v.x * scalar
	sy := v.y * scalar
	sz := v.z * scalar
	return NewVec3(sx, sy, sz)
}

func (v Vec3) Divide(vec Vec3) Vec3 {
	if vec.x == 0.0 || vec.y == 0.0 || vec.z == 0.0 {
		panic("cant divide by zero")
	}
	sx := v.x / vec.x
	sy := v.y / vec.y
	sz := v.z / vec.z
	return NewVec3(sx, sy, sz)
}

func (v Vec3) DivScalar(scalar float64) Vec3 {
	if scalar == 0.0 {
		panic("cant divide by zero")
	}
	sx := v.x / scalar
	sy := v.y / scalar
	sz := v.z / scalar
	return NewVec3(sx, sy, sz)
}

func (v Vec3) Equals(vec Vec3, epsilon float64) bool {
	dx := math.Abs(v.x-vec.x) < epsilon
	dy := math.Abs(v.y-vec.y) < epsilon
	dz := math.Abs(v.z-vec.z) < epsilon
	return dx && dy && dz
}

func (v *Vec3) Normalize() float64 {
	norm := v.Norm()
	invn := 1.0 / norm
	v.x *= invn
	v.y *= invn
	v.z *= invn
	return norm
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

func (v Vec3) X() float64 {
	return v.x
}

func (v Vec3) Y() float64 {
	return v.y
}

func (v Vec3) Z() float64 {
	return v.z
}
