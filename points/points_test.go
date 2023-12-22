package points

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type operation int

const (
	sum operation = iota
	sumScalar
	sub
	subScalar
	mul
	div
)

func TestTableDrivenPoint(t *testing.T) {
	testCases := []struct {
		desc     string
		p1       Point[float64]
		p2       Point[float64]
		factor   float64
		op       operation
		expected Point[float64]
	}{
		{
			desc:     "sum (2, 2) + (2, 2) = (4, 4)",
			p1:       NewPoint(2.0, 2.0),
			p2:       NewPoint(2.0, 2.0),
			op:       sum,
			expected: NewPoint(4.0, 4.0),
		},
		{
			desc:     "sum (2, 2) + (-2, -2) = (0, 0)",
			p1:       NewPoint(2.0, 2.0),
			p2:       NewPoint(-2.0, -2.0),
			op:       sum,
			expected: NewPoint(0.0, 0.0),
		},
		{
			desc:     "sum scalar (2, 2) + 2 = (4, 4)",
			p1:       NewPoint(2.0, 2.0),
			factor:   2.9,
			op:       sumScalar,
			expected: NewPoint(4.9, 4.9),
		},
		{
			desc:     "sub (2, 2) - (2, 2) = (0, 0)",
			p1:       NewPoint(2.0, 2.0),
			p2:       NewPoint(2.0, 2.0),
			op:       sub,
			expected: NewPoint(0.0, 0.0),
		},
		{
			desc:     "sub scalar (2.1, 2.3) - 4.5 = (-2.4, -2.2)",
			p1:       NewPoint(2.1, 2.3),
			factor:   4.5,
			op:       subScalar,
			expected: NewPoint(-2.4, -2.2),
		},
		{
			desc:     "div (10, 3) / 2 = (5, 1.5)",
			p1:       NewPoint(10.0, 3.0),
			factor:   2.0,
			op:       div,
			expected: NewPoint(5.0, 1.5),
		},
		{
			desc:     "mul (10, 3) * 2 = (20.0, 6.0)",
			p1:       NewPoint(10.0, 3.0),
			factor:   2.0,
			op:       mul,
			expected: NewPoint(20.0, 6.0),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var result Point[float64]
			switch tC.op {
			case sum:
				result = tC.p1.Sum(tC.p2)
			case sumScalar:
				result = tC.p1.SumScalar(tC.factor)
			case sub:
				result = tC.p1.Subtract(tC.p2)
			case subScalar:
				result = tC.p1.SubScalar(tC.factor)
			case mul:
				result = tC.p1.Multiply(tC.factor)
			case div:
				result = tC.p1.Divide(tC.factor)
			default:
				assert.Fail(t, "function not implemented")
			}
			assert.Equal(t, tC.expected, result)
		})
	}
}

func BenchmarkPoint(b *testing.B) {
	p1 := NewPoint(2, 2)
	p2 := NewPoint(2, 2)
	p3 := NewPoint(2.5, 3.775)
	b.Run("benchmark sum", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = p1.Sum(p2)
		}
	})

	b.Run("benchmark sum scalar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = p3.SumScalar(3.7)
		}
	})

	b.Run("benchmark subtract", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = p1.Subtract(p2)
		}
	})

	b.Run("benchmark subtract scalar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = p3.SubScalar(3.7)
		}
	})
}

func TestPoint(t *testing.T) {
	t.Run("test distance to another point", func(t *testing.T) {
		p1 := NewPoint(1, 2)
		p2 := NewPoint(3, 4)
		distance := p1.DistanceTo(p2)
		assert.InDelta(t, 2.8284, distance, 0.0001)
	})

	t.Run("test squared distance to another point", func(t *testing.T) {
		p1 := NewPoint(1, 2)
		p2 := NewPoint(3, 4)
		distance := p1.SquaredDistanceTo(p2)
		assert.InDelta(t, 8, distance, 0.0001)
	})

	t.Run("test expression (a + b) * 3 + (a - c)", func(t *testing.T) {
		p1 := NewPoint(1, 2)
		p2 := NewPoint(3, 4)
		p3 := NewPoint(5, 6)
		result := p1.Sum(p2).Multiply(3).Sum(p1.Subtract(p3))
		assert.Equal(t, NewPoint(8, 14), result)
		assert.Equal(t, 8, result.X())
		assert.Equal(t, 14, result.Y())
	})

	t.Run("test division by zero must panic", func(t *testing.T) {
		p1 := NewPoint(2.0, 5.0)
		assert.Panics(t, func() {
			p1.Divide(0.0)
		})
	})

	t.Run("test if p1 is equals to p2", func(t *testing.T) {
		p1 := NewPoint(1, 2)
		p2 := NewPoint(1, 2)
		result := p1.Equals(p2)
		assert.True(t, result)
	})
}
