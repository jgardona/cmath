// # Ranges
//
// This package contains a generic data struct that represents an integer or float interval,
// with minimum and maximum values.
package ranges

import "github.com/jgardona/cmath/constraints"

// Range represents an integer or float interval with minimum and maximum values.
//
// # Note
//
// The structure represents a generic type with inclusive limits -
// both minimum and maximum values for the interval are included into it.
// The mathematical notation of such interval is `[min, max]`.
type Range[T constraints.Numbers] struct {
	min T
	max T
}

// NewRange instantiates a new Range with min and max limits.
func NewRange[T constraints.Numbers](min, max T) Range[T] {
	return Range[T]{min, max}
}

// Min returns the minimum limit from Range.
func (r Range[T]) Min() T {
	return r.min
}

// Max returns the maximum limit from Range.
func (r Range[T]) Max() T {
	return r.max
}

// The Length of the Range(difference between maximum and minimum values).
func (r Range[T]) Length() T {
	return r.max - r.min
}

// IsInside checks if the specified scalar is inside the range.
func (r Range[T]) IsInside(scalar T) bool {
	return scalar > r.min && scalar < r.max
}

// IsRangeInside checks if the specified range is inside this range.
func (r Range[T]) IsRangeInside(a Range[T]) bool {
	return (r.IsInside(a.min)) && (r.IsInside(a.max))
}

// IsOverlapping checks if the specified range is inside this range.
func (r Range[T]) IsOverlapping(a Range[T]) bool {
	return (r.IsInside(a.min)) || (r.IsInside(a.max)) ||
		(a.IsInside(r.min)) || (a.IsInside(r.max))
}

// Equals checks if specified range is equals to this range. For more
// accurate results using floats, Delta(a, b) is recommended.
func (r Range[T]) Equals(a Range[T]) bool {
	return r.min == a.min && r.max == a.max
}
