package statistics

import (
	"github.com/jgardona/cmath/ranges"
)

func Mean(histogram []int) float64 {
	panic("not yet implemented")
}

func stdDevMean(histogram []int, mean float64) float64 {
	panic("not yet implemented")
}

func StdDev(histogram []int) float64 {
	return stdDevMean(histogram, Mean(histogram))
}

func Median(histogram []int) int {
	panic("not yet implemented")
}

func Range(histogram []int, percent float64) ranges.Range[int] {
	panic("not yet implemented")
}

func Entropy(histogram []int) float64 {
	panic("not yet implemented")
}

func Mode(histogram []int) int {
	panic("not yet implemented")
}
