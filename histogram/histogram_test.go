package histogram

import (
	"testing"

	"github.com/jgardona/cmath/ranges"
	"github.com/stretchr/testify/assert"
)

func TestHistogram(t *testing.T) {
	histogram := []int{0, 0, 1, 3, 6, 8, 11, 0, 0, 0}
	t.Run("the mean must return float value", func(t *testing.T) {
		h := NewHistogram(histogram)
		mean := h.Mean()
		assert.InDelta(t, 4.862, mean, 0.001)
	})

	t.Run("total must return int value", func(t *testing.T) {
		h := NewHistogram(histogram)
		total := h.Total()
		assert.Equal(t, 29, total)
	})

	t.Run("max must return int value", func(t *testing.T) {
		h := NewHistogram(histogram)
		max := h.Max()
		assert.Equal(t, 6, max)
	})

	t.Run("min must return int value", func(t *testing.T) {
		h := NewHistogram(histogram)
		min := h.Min()
		assert.Equal(t, 2, min)
	})

	t.Run("median must return int value", func(t *testing.T) {
		h := NewHistogram(histogram)
		median := h.Median()
		assert.Equal(t, 5, median)
	})

	t.Run("stddev must return float value", func(t *testing.T) {
		h := NewHistogram(histogram)
		stddev := h.StdDev()
		assert.InDelta(t, 1.136, stddev, 0.001)
	})

	t.Run("rang must return an IntRange", func(t *testing.T) {
		h := NewHistogram(histogram)
		r := h.Range(0.5)
		assert.Equal(t, ranges.NewRange(4, 6), r)
	})
}

func BenchmarkHistogram(b *testing.B) {
	histogram := []int{0, 0, 1, 3, 6, 8, 11, 0, 0, 0}
	h := NewHistogram(histogram)

	b.Run("benchmark range function", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = h.Range(0.5)
		}
	})

	b.Run("benchmark mean function", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = h.Mean()
		}
	})

	b.Run("benchmark median function", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = h.Median()
		}
	})

	b.Run("benchmark stddev function", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = h.StdDev()
		}
	})

	b.Run("benchmark update function", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h.Update()
		}
	})

	b.Run("benchmark NewHistogram/Update function", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := NewHistogram([]int{0, 0, 1, 2, 3, 2, 1, 0, 0})
			h.Update()
		}
	})

}
