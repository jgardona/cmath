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

	t.Run("test max function", func(t *testing.T) {
		result := v1.Max()
		assert.Equal(t, 3.0, result)
	})

	t.Run("test min function", func(t *testing.T) {
		result := v1.Min()
		assert.Equal(t, 1.0, result)
	})

	t.Run("test maxindex function", func(t *testing.T) {
		result := v1.MaxIndex()
		assert.Equal(t, 2, result)

		result = v2.MaxIndex()
		assert.Equal(t, 0, result)

		result = v3.MaxIndex()
		assert.Equal(t, 1, result)

		result = v4.MaxIndex()
		assert.Equal(t, 2, result)
	})

	t.Run("test minindex function", func(t *testing.T) {
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

	t.Run("test norm function", func(t *testing.T) {
		result := v1.Norm()
		assert.InDelta(t, 3.741, result, 0.001)
	})
}
