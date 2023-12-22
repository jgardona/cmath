package statistics

import (
	"testing"

	"github.com/jgardona/cmath/ranges"
	"github.com/stretchr/testify/assert"
)

func TestStatisticsPackage(t *testing.T) {
	histogram := []int{1, 1, 2, 3, 6, 8, 11, 12, 7, 3}
	histogramEntropyTest1 := []int{3, 3}
	histogramEntropyTest2 := []int{1, 1, 1, 1}
	histogramEntropyTest3 := []int{1, 2, 3, 4}
	histogramNormal := []int{1, 3, 6, 9, 6, 3, 1}
	histogramEmpty := []int{}
	histogramTestMode := []int{1, 1, 2, 3, 6, 8, 11, 12, 7, 3}

	t.Run("the Mean must be", func(t *testing.T) {
		result := Mean(histogram)
		assert.InDelta(t, 5.759, result, 0.001)
	})

	t.Run("the Mean must be zero for empty histogram", func(t *testing.T) {
		result := Mean(histogramEmpty)
		assert.InDelta(t, 0.0, result, 0.001)
	})

	t.Run("the StdDevMean must be", func(t *testing.T) {
		result := stdDevMean(histogram, Mean(histogram))
		assert.InDelta(t, 1.999, result, 0.001)
	})

	t.Run("the StdDev must be", func(t *testing.T) {
		result := StdDev(histogram)
		assert.InDelta(t, 1.999, result, 0.001)
	})

	t.Run("the StdDev must be zero from an empty histogram", func(t *testing.T) {
		result := StdDev(histogramEmpty)
		assert.InDelta(t, 0.0, result, 0.001)
	})

	t.Run("the Median return must be", func(t *testing.T) {
		result := Median(histogram)
		assert.Equal(t, 6, result)
	})

	t.Run("the Median must be zero", func(t *testing.T) {
		result := Median(histogramEmpty)
		assert.Equal(t, 0, result)
	})

	t.Run("Range must return [4, 8]", func(t *testing.T) {
		result := Range(histogram, 0.75)
		assert.Equal(t, ranges.NewRange(4, 8), result)

		result = Range(histogramNormal, 0.5)
		assert.Equal(t, ranges.NewRange(2, 4), result)
	})

	t.Run("Entropy must be", func(t *testing.T) {
		result := Entropy(histogramEntropyTest1)
		assert.InDelta(t, 1.000, result, 0.001)

		result = Entropy(histogramEntropyTest2)
		assert.InDelta(t, 2.000, result, 0.001)

		result = Entropy(histogramEntropyTest3)
		assert.InDelta(t, 1.846, result, 0.001)
	})

	t.Run("Mode must return seven", func(t *testing.T) {
		result := Mode(histogramTestMode)
		assert.Equal(t, 7, result)
	})

}

func BenchmarkStatisticsPakcage(b *testing.B) {
	histogram := []int{1, 1, 2, 3, 6, 8, 11, 12, 7, 3}

	b.Run("mean must not allocate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Mean(histogram)
		}
	})

	b.Run("median must not allocate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Median(histogram)
		}
	})
}
