package cmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeltaFunction(t *testing.T) {
	t.Run("must 2.57134 be equals 2.57136 camparing delta", func(t *testing.T) {
		assert.True(t, Delta(2.57134, 2.57136))
	})

	t.Run("must 2.5234 be not equals 2.57156 camparing delta", func(t *testing.T) {
		assert.True(t, !Delta(2.57234, 2.57156))
	})
}

func BenchmarkDeltaFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Delta(2.57134, 57136)
	}
}

func TestPop(t *testing.T) {
	data := []float64{0, 1, 2, 3}
	value := Pop(&data)
	assert.Equal(t, []float64{0, 1, 2}, data)
	assert.Equal(t, 3.0, value)
}

func TestPush(t *testing.T) {
	data := []float64{0, 1, 2, 3}
	Push(&data, 4)
	assert.Equal(t, []float64{0, 1, 2, 3, 4}, data)
}

func TestReturnEmptyValue(t *testing.T) {
	data := []float64{}
	value := Pop(&data)
	assert.Equal(t, 0.0, value)
}
