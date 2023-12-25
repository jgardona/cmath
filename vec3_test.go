package cmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec3(t *testing.T) {
	v1 := NewVec3(1, 2, 3)
	v2 := NewVec3(3, 2, 1)
	v3 := NewVec3(1, 3, 2)
	v4 := NewVec3(2, 1, 3)
	v5 := NewVec3(2, 3, 1)
	v6 := NewVec3(0, 0, 0)

	t.Run("test max method", func(t *testing.T) {
		result := v1.Max()
		assert.Equal(t, 3.0, result)
	})

	t.Run("test min method", func(t *testing.T) {
		result := v1.Min()
		assert.Equal(t, 1.0, result)
	})

	t.Run("test maxindex method", func(t *testing.T) {
		result := v1.MaxIndex()
		assert.Equal(t, 2, result)

		result = v2.MaxIndex()
		assert.Equal(t, 0, result)

		result = v3.MaxIndex()
		assert.Equal(t, 1, result)

		result = v4.MaxIndex()
		assert.Equal(t, 2, result)
	})

	t.Run("test minindex method", func(t *testing.T) {
		result := v1.MinIndex()
		assert.Equal(t, 0, result)

		result = v2.MinIndex()
		assert.Equal(t, 2, result)

		result = v3.MinIndex()
		assert.Equal(t, 0, result)

		result = v4.MinIndex()
		assert.Equal(t, 1, result)

		result = v5.MinIndex()
		assert.Equal(t, 2, result)
	})

	t.Run("test norm method", func(t *testing.T) {
		result := v1.Norm()
		assert.InDelta(t, 3.741, result, 0.001)
	})

	t.Run("test dot method", func(t *testing.T) {
		result := v3.Dot()
		assert.Equal(t, 14.0, result)

		result = v1.Dot()
		assert.Equal(t, 14.0, result)
	})

	t.Run("test asarray method", func(t *testing.T) {
		result := v5.AsArray()
		assert.Equal(t, [3]float64{2, 3, 1}, result)
	})

	t.Run("test sum method", func(t *testing.T) {
		result := v4.Sum(v1)
		assert.Equal(t, NewVec3(3, 3, 6), result)
	})

	t.Run("test sum scalar method", func(t *testing.T) {
		result := v1.SumScalar(2.0)
		assert.Equal(t, NewVec3(3.0, 4.0, 5.0), result)
	})

	t.Run("test subtract method", func(t *testing.T) {
		result := v1.Subtract(v2)
		assert.Equal(t, NewVec3(-2, 0, 2), result)
	})

	t.Run("test subscalar method", func(t *testing.T) {
		result := v1.SubScalar(1)
		assert.Equal(t, NewVec3(0, 1, 2), result)
	})

	t.Run("test multiply method", func(t *testing.T) {
		result := v2.Multiply(v3)
		assert.Equal(t, NewVec3(3, 6, 2), result)
	})

	t.Run("test mulscalar method", func(t *testing.T) {
		result := v2.MulScalar(2)
		assert.Equal(t, NewVec3(6, 4, 2), result)
	})

	t.Run("test divide method", func(t *testing.T) {
		result := v2.Divide(v1)
		assert.Equal(t, 3.0, result.X())
		assert.Equal(t, 1.0, result.Y())
		assert.InDelta(t, 0.333, result.Z(), 0.001)

		// fail dividing by zero
		assert.Panics(t, func() {
			_ = v1.Divide(v6)
		})
	})

	t.Run("test divide scalar method", func(t *testing.T) {
		result := v2.DivScalar(2)
		assert.InDelta(t, 1.5, result.X(), 0.001)
		assert.InDelta(t, 1.0, result.Y(), 0.001)
		assert.InDelta(t, 0.5, result.Z(), 0.001)

		// fail dividing by zero
		assert.Panics(t, func() {
			_ = v1.DivScalar(0.0)
		})
	})

	t.Run("test equals method", func(t *testing.T) {
		result := v2.DivScalar(3)
		vec := NewVec3(1.0, 0.666, 0.333)
		assert.True(t, result.Equals(vec, 0.001))
	})
}
