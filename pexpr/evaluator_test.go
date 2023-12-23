package pexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolishEvaluator(t *testing.T) {
	t.Run("test must do 3 + 2 + 1.5 equals 6.5", func(t *testing.T) {
		ev := NewPolishEvaluator("3 2 + $0 +", []float64{1.5})
		result, err := ev.Evaluate()
		assert.Nil(t, err)
		assert.InDelta(t, 6.5, result, 0.001)
	})

	t.Run("test must do 3 - 2 equals 1", func(t *testing.T) {
		ev := NewPolishEvaluator("3 $0 -", []float64{2.0})
		result, err := ev.Evaluate()
		assert.Nil(t, err)
		assert.InDelta(t, 1.0, result, 0.001)
	})

	t.Run("test subtraction must fail with invalid syntax error", func(t *testing.T) {
		ev := NewPolishEvaluator("3$0 -", []float64{2.0})
		result, err := ev.Evaluate()
		assert.Equal(t, 0.0, result)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "strconv.ParseFloat: parsing \"3$0\": invalid syntax")
	})

	t.Run("test addition must fail with cant parse integer invalid syntax", func(t *testing.T) {
		ev := NewPolishEvaluator("3 $a +", []float64{2.0})
		result, err := ev.Evaluate()
		assert.Equal(t, 0.0, result)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "strconv.ParseInt: parsing \"a\": invalid syntax")
	})

	t.Run("test addition must fail with function not implemented", func(t *testing.T) {
		ev := NewPolishEvaluator("3 $0 ]", []float64{2.0})
		result, err := ev.Evaluate()
		assert.Equal(t, 0.0, result)
		assert.NotNil(t, err)
		assert.Equal(t, err, ErrBadFunction)
	})

	t.Run("test division results of 10 and 2 must be equal 5", func(t *testing.T) {
		ev := NewPolishEvaluator("10 $0 /", []float64{2.0})
		result, err := ev.Evaluate()
		assert.Nil(t, err)
		assert.InDelta(t, 5.0, result, 0.001)
	})

	t.Run("test cant divide by zero", func(t *testing.T) {
		ev := NewPolishEvaluator("10 $0 /", []float64{0.0})
		result, err := ev.Evaluate()
		assert.NotNil(t, err)
		assert.Equal(t, ErrDivisionByZero, err)
		assert.InDelta(t, 0.0, result, 0.001)
	})

	t.Run("test multiplication of 2 and 3 must equal 6", func(t *testing.T) {
		ev := NewPolishEvaluator("$0 $1 *", []float64{2.0, 3.0})
		result, err := ev.Evaluate()
		assert.Nil(t, err)
		assert.InDelta(t, 6.0, result, 0.001)
	})

	t.Run("test cos function", func(t *testing.T) {
		ev := NewPolishEvaluator("$0 cos", []float64{10.0})
		result, err := ev.Evaluate()
		assert.Nil(t, err)
		assert.InDelta(t, -0.8390715290764524, result, 0.001)
	})

	t.Run("test exp function", func(t *testing.T) {
		ev := NewPolishEvaluator("$0 exp", []float64{5.0})
		result, err := ev.Evaluate()
		assert.Nil(t, err)
		assert.InDelta(t, 148.4131591025766, result, 0.001)
	})

	t.Run("test ln function", func(t *testing.T) {
		ev := NewPolishEvaluator("$0 ln", []float64{3.0})
		result, err := ev.Evaluate()
		assert.Nil(t, err)
		assert.InDelta(t, 1.0986122886681098, result, 0.001)
	})

	t.Run("test sine function", func(t *testing.T) {
		ev := NewPolishEvaluator("$0 sin", []float64{3.0})
		result, err := ev.Evaluate()
		assert.Nil(t, err)
		assert.InDelta(t, 0.1411200080598672, result, 0.001)
	})

	t.Run("test sine function", func(t *testing.T) {
		ev := NewPolishEvaluator("$0 sqrt", []float64{25.0})
		result, err := ev.Evaluate()
		assert.Nil(t, err)
		assert.InDelta(t, 5.0, result, 0.001)
	})
}
func BenchmarkPolishExpression(b *testing.B) {
	ev1 := NewPolishEvaluator("$0 $1 *", []float64{3.0, 2.0})
	ev2 := NewPolishEvaluator("$0 1 *", []float64{3.0})
	ev3 := NewPolishEvaluator("$0 $1 + $2 $3 * cos sin ln sqrt", []float64{3.0, 1.0, 2.0, 5.0})

	b.Run("benchmark two variable expression", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = ev1.Evaluate()
		}
	})

	b.Run("benchmark one variable expression", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = ev2.Evaluate()
		}
	})

	b.Run("benchmark multiple variable expression", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = ev3.Evaluate()
		}
	})
}
