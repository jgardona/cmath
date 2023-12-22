package ranges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	r1 := NewRange(-1, 2)
	r2 := NewRange(1, 1)
	r3 := NewRange(1, 3)
	r4 := NewRange(-1, 2)

	t.Run("scalar 1 must be inside r1", func(t *testing.T) {
		result := r1.IsInside(1)
		assert.True(t, result)
	})

	t.Run("range r3 must overlap range r1", func(t *testing.T) {
		result := r1.IsOverlapping(r3)
		assert.True(t, result)
	})

	t.Run("range r2 must be inside r1", func(t *testing.T) {
		result := r1.IsRangeInside(r2)
		assert.True(t, result)
	})

	t.Run("the lengh of r1 must be 1 (max - min)", func(t *testing.T) {
		result := r1.Length()
		assert.Equal(t, 3, result)
	})

	t.Run("r1 must be equals r4 and not equals r2", func(t *testing.T) {
		result := r1.Equals(r4)
		assert.True(t, result)
		result = r1.Equals(r2)
		assert.False(t, result)
	})
}

func BenchmarkRanges(b *testing.B) {
	r1 := NewRange(-1.5, 3.0)
	r2 := NewRange(1.0, 2.0)
	r3 := NewRange(-1.5, 3.5)

	b.Run("range isInside must not allocate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = r1.IsInside(1.0)
		}
	})

	b.Run("range isRangeInside must not allocate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = r1.IsRangeInside(r2)
		}
	})

	b.Run("range isOverlapping must not allocate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = r1.IsOverlapping(r3)
		}
	})
}
