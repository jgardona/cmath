// # Statistics
//
// Set of statistics functions for golang
package statistics

import (
	"math"

	"github.com/jgardona/cmath/ranges"
)

// Mean calculates the mean.
//
// The input array is treated as histogram, i.e. its
// indexes are treated as velues of stochastic function, but
// array values are threated as probabilities (total amount of hits).
func Mean(histogram []int) float64 {
	var (
		mean  float64 = 0.0
		total int     = 0.0
	)

	for i, e := range histogram {
		total += e
		mean += float64(i * e)
	}

	if total == 0 {
		return 0.0
	} else {
		return mean / float64(total)
	}
}

// stdDevMean is only used inside statistics package.
// It calculates the standard deviation using the mean value.
//
// The input array is treated as histogram, i.e. its
// indexes are treated as velues of stochastic function, but
// array values are threated as probabilities (total amount of hits).
func stdDevMean(histogram []int, mean float64) float64 {
	var (
		total  int     = 0
		stddev float64 = 0.0
		diff   float64 = 0.0
	)
	for i, e := range histogram {
		diff = float64(i) - mean
		stddev += diff * diff * float64(e)
		total += e
	}

	if total == 0 {
		return 0
	} else {
		return math.Sqrt(stddev / float64(total))
	}
}

// StdDev calculates the standard deviation.
// The input array is treated as histogram, i.e. its
// indexes are treated as velues of stochastic function, but
// array values are threated as probabilities (total amount of hits).
func StdDev(histogram []int) float64 {
	return stdDevMean(histogram, Mean(histogram))
}

// Median calculates the median value
//
// The input array is treated as histogram, i.e. its
// indexes are treated as velues of stochastic function, but
// array values are threated as probabilities (total amount of hits).
//
// The median value is calculated accumulating histogram's values
// starting from the left point until the sum reaches 50% of histogram's sum.
func Median(histogram []int) int {
	var total int = 0

	for _, e := range histogram {
		total += e
	}

	htotal := total / 2
	median := 0
	v := 0

	for ; median < len(histogram); median++ {
		v += histogram[median]
		if v >= htotal {
			break
		}
	}
	return median
}

// Range gets the range around the median containing specified percentage of values.
//
// The method calculates ranges of stochastic variable, wihch summary probability
// comprises the specified percentage of histogram's hits.
func Range(histogram []int, percent float64) ranges.Range[int] {
	var total, n int = 0, len(histogram)

	// accumulate total
	for _, e := range histogram {
		total += e
	}

	var min, max, hits int = 0, 0, total
	h := int(float64(total) * (percent + (1.0-percent)/2.0))

	for min, hits = 0, total; min < n; min++ {
		hits -= histogram[min]
		if hits < h {
			break
		}
	}

	for max, hits = n-1, total; max >= 0; max-- {
		hits -= histogram[max]
		if hits < h {
			break
		}
	}

	return ranges.NewRange(min, max)
}

// Entropy calculates the entropy value.
//
// The input array is treated as histogram, i.e. its
// indexes are treated as velues of stochastic function, but
// array values are threated as probabilities (total amount of hits).
func Entropy(histogram []int) float64 {
	var (
		total   float64 = 0.0
		entropy float64 = 0.0
		p       float64 = 0.0
	)

	for _, e := range histogram {
		total += float64(e)
	}

	if total != 0 {
		for _, e := range histogram {
			p = float64(e) / total
			if p != 0 {
				entropy += (-p * math.Log2(p))
			}
		}
	}
	return entropy
}

// Mode calculates mode value.
//
// Returns mode value of the histogram array.
//
// The input array is treated as histogram, i.e. its
// indexes are treated as velues of stochastic function, but
// array values are threated as probabilities (total amount of hits).
//
// # Note
//
// Returns the minimum mode value if the specified histogram is multimodal.
func Mode(histogram []int) int {
	var mode, curmax int = 0, 0
	for i, e := range histogram {
		if e > curmax {
			curmax = e
			mode = i
		}
	}
	return mode
}
