package statistics

import (
	"math"

	"github.com/jgardona/cmath/ranges"
)

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

func StdDev(histogram []int) float64 {
	return stdDevMean(histogram, Mean(histogram))
}

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

func Mode(histogram []int) int {
	panic("not yet implemented")
}
