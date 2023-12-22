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

// A histogram for discrete random values, with mean,
// standard deviation, median, min, max and total count values.
type Histogram struct {
	values []int
	mean   float64
	stddev float64
	median int
	min    int
	max    int
	total  int
}

// NewHistogram instantiates a histogram given an int array.
func NewHistogram(histogram []int) *Histogram {
	data := &Histogram{values: histogram}
	data.Update()
	return data
}

// Values returns the histogram array.
func (h Histogram) Values() []int {
	return h.values
}

// Mean returns the histogram mean.
func (h Histogram) Mean() float64 {
	return h.mean
}

// StdDev returns the histogram's standard deviation.
func (h Histogram) StdDev() float64 {
	return h.stddev
}

// Median returns the histogram's median.
func (h Histogram) Median() int {
	return h.median
}

// Min returns the histogram's minimum value.
func (h Histogram) Min() int {
	return h.min
}

// Max returns the histogram's maximum value.
func (h Histogram) Max() int {
	return h.max
}

// Total returns the histogram's total count.
func (h Histogram) Total() int {
	return h.total
}

// Range returns an int range around the mean, given the percent.
func (h Histogram) Range(percent float64) ranges.Range[int] {
	return statistics.Range(h.values, percent)
}

// Update updates all values in histogram. This function must only
// be called to recalculate the histogram.
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
