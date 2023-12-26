// # vec3
//
// This package contains a 3D vector structure with x, y, z coordinates.
package cmath

import (
	"math"
)

// # Vec3
//
// A 3D data structure with x, y and z coordinates.
// The structure encapsulates x, y and z coordinates of a 3d vector and
// provides some linear algebra operations.
type Vec3 struct {
	x float64
	y float64
	z float64
}

// NewVec3 initializes a Vec3 structure.
func NewVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

// Max returns the maximum value of the vector.
func (v Vec3) Max() float64 {
	return math.Max(v.x, math.Max(v.y, v.z))
}

// Min returns the minimum value of the vector.
func (v Vec3) Min() float64 {
	return math.Min(v.x, math.Min(v.y, v.z))
}

// MaxIndex returns the vector's index which has the maximum value.
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

// MinIndex returns the vector's index which has the minimum value.
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

// Norm returns the norm of a vector.
func (v Vec3) Norm() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

// Dot return the dot product of vector himself.
func (v Vec3) Dot() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

// AsArray returns a vector representation as array
func (v Vec3) AsArray() [3]float64 {
	return [3]float64{v.x, v.y, v.z}
}

// Sum implements vector addition.
func (v Vec3) Sum(vec Vec3) Vec3 {
	sx := v.x + vec.x
	sy := v.y + vec.y
	sz := v.z + vec.z

	return NewVec3(sx, sy, sz)
}

// SumScalar adds a scalar.
func (v Vec3) SumScalar(scalar float64) Vec3 {
	sx := v.x + scalar
	sy := v.y + scalar
	sz := v.z + scalar
	return NewVec3(sx, sy, sz)
}

// Subtract implements vector subtraction.
func (v Vec3) Subtract(vec Vec3) Vec3 {
	sx := v.x - vec.x
	sy := v.y - vec.y
	sz := v.z - vec.z

	return NewVec3(sx, sy, sz)
}

// SubScalar subtracts a scalar.
func (v Vec3) SubScalar(scalar float64) Vec3 {
	sx := v.x - scalar
	sy := v.y - scalar
	sz := v.z - scalar
	return NewVec3(sx, sy, sz)
}

// Multiply implements vector multiplication.
func (v Vec3) Multiply(vec Vec3) Vec3 {
	sx := v.x * vec.x
	sy := v.y * vec.y
	sz := v.z * vec.z
	return NewVec3(sx, sy, sz)
}

// MulScalar implements scalar vector multiplication.
func (v Vec3) MulScalar(scalar float64) Vec3 {
	sx := v.x * scalar
	sy := v.y * scalar
	sz := v.z * scalar
	return NewVec3(sx, sy, sz)
}

// Divide implements vector division.
func (v Vec3) Divide(vec Vec3) Vec3 {
	if vec.x == 0.0 || vec.y == 0.0 || vec.z == 0.0 {
		panic("cant divide by zero")
	}
	sx := v.x / vec.x
	sy := v.y / vec.y
	sz := v.z / vec.z
	return NewVec3(sx, sy, sz)
}

// DivScalar implements vector scalar division.
func (v Vec3) DivScalar(scalar float64) Vec3 {
	if scalar == 0.0 {
		panic("cant divide by zero")
	}
	sx := v.x / scalar
	sy := v.y / scalar
	sz := v.z / scalar
	return NewVec3(sx, sy, sz)
}

// Equals compares float numbers using epsilon.
func (v Vec3) Equals(vec Vec3, epsilon float64) bool {
	dx := math.Abs(v.x-vec.x) < epsilon
	dy := math.Abs(v.y-vec.y) < epsilon
	dz := math.Abs(v.z-vec.z) < epsilon
	return dx && dy && dz
}

// Normalize normalizes this vector by dividing itś coordinates with the vector's norm.
// Returns the value of vector's norm before normalization.
func (v *Vec3) Normalize() float64 {
	norm := v.Norm()
	invn := 1.0 / norm
	v.x *= invn
	v.y *= invn
	v.z *= invn
	return norm
}

// Inverse inverts the vector.
// Returns a vector with all coordinate equal to 1.0 divided by the value of correspondent coordinate
// in this vector (or equal to 0.0 if this vector has corresponding coordinate also set to 0.0).
func (v Vec3) Inverse() Vec3 {
	var nx, ny, nz float64 = 0.0, 0.0, 0.0
	if v.x != 0.0 {
		nx = 1.0 / v.x
	}
	if v.y != 0.0 {
		ny = 1.0 / v.y
	}
	if v.z != 0.0 {
		nz = 1.0 / v.z
	}
	return NewVec3(nx, ny, nz)
}

// Abs calculates the vector's obsolute value.
// Returns a vector with all coordinates equal to absolute values of this vector's coordinates.
func (v Vec3) Abs() Vec3 {
	ax := math.Abs(v.x)
	ay := math.Abs(v.y)
	az := math.Abs(v.z)
	return NewVec3(ax, ay, az)
}

// Cross calculates the cross product of two vectors.
func Cross(a, b Vec3) Vec3 {
	cx := a.y*b.z - a.z*b.y
	cy := a.z*b.x - a.x*b.z
	cz := a.x*b.y - a.y*b.x
	return NewVec3(cx, cy, cz)
}

// Dot Calculates the dot product of two vectors.
func Dot(a, b Vec3) float64 {
	dot := a.x*b.x + a.y*b.y + a.z*b.z
	return dot
}

// AsVec4 returns a vec4 representation of this vector.
func (v Vec3) AsVec4() Vec4 {
	panic("not yet implemented")
}

// X returns the vector's x coordinate.
func (v Vec3) X() float64 {
	return v.x
}

// Y returns the vector's y coordinate.
func (v Vec3) Y() float64 {
	return v.y
}

// Z returns the vector's z coordinate
func (v Vec3) Z() float64 {
	return v.z
}
