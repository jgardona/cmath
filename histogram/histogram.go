// # Histogram
//
// # A histogram for discrete random values
//
// This package contains a data structure representing a histogram,
// which is a wrapper for an integer array, where indexes of the array are
// treated as values of stochastic function. but array
package histogram

import (
	"github.com/jgardona/cmath/ranges"
	"github.com/jgardona/cmath/statistics"
)

type Histogram struct {
	values []int
	mean   float64
	stddev float64
	median int
	min    int
	max    int
	total  int
}

func NewHistogram(histogram []int) *Histogram {
	data := &Histogram{values: histogram}
	data.Update()
	return data
}

func (h Histogram) Values() []int {
	return h.values
}

func (h Histogram) Mean() float64 {
	return h.mean
}

func (h Histogram) StdDev() float64 {
	return h.stddev
}

func (h Histogram) Median() int {
	return h.median
}

func (h Histogram) Min() int {
	return h.min
}

func (h Histogram) Max() int {
	return h.max
}

func (h Histogram) Total() int {
	return h.total
}

func (h Histogram) Range(percent float64) ranges.Range[int] {
	return statistics.Range(h.values, percent)
}

func (h *Histogram) Update() {
	var i int
	var n int = len(h.values)
	h.min = n
	for i = 0; i < n; i++ {
		if h.values[i] != 0 {
			if i > h.max {
				h.max = i
			}

			if i < h.min {
				h.min = i
			}

			h.total += h.values[i]
		}
	}

	h.mean = statistics.Mean(h.values)
	h.stddev = statistics.StdDev(h.values)
	h.median = statistics.Median(h.values)
}
