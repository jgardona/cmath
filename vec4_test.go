package cmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec4(t *testing.T) {
	v1 := NewVec4(1, 2, 3, 4)
	v2 := NewVec4(4, 3, 2, 1)
	//v3 := NewVec4(3, 4, 2, 1)
	//v4 := NewVec4(2, 1, 3, 4)

	t.Run("test dot method", func(t *testing.T) {

	})

	t.Run("test min method", func(t *testing.T) {
		result := v1.Min()
		assert.Equal(t, 1.0, result)
	})

	t.Run("test max method", func(t *testing.T) {
		result := v2.Max()
		assert.Equal(t, 4.0, result)
	})
}
