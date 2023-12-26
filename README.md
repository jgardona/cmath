# cmath

A core math module for golang, with statistics, histogram, points, ranges and more math tools. This module is almost a port of the core and math libraries of Aforge.Net.

## Usage

### Points

A data structure for representing a pair of coordinates.

- Creating two points. An int point and a float point.

```go
p2 := NewPoint(2, 2)
p3 := NewPoint(2.5, 3.775)
```

- Adding two points

```go
result := p1.Sum(p2)
```

- Calculates the euclidean distance between two points.

```go
distance := p2.DistanceTo(p3)
```

### Histogram

A histogram for discrete values.

- Initializing a new histogram.

```go
hist := NewHistogram([]int{1, 2, 3, 10, 3, 2, 1})

// returning the mean, standard deviation, min, and max values.

mean := hist.Mean()
stddev := hist.StdDev()
min := hist.Min()
max := hist.Max()
```
### Ranges

Generic data structure that represents an integer or float interval.

- Initializes a new range and check if some values are in inside the interval.

```go
r1 := NewRange(-0.5, 1.0)
r2 := NewRange(1.0, 1.5)

if r1.IsRangeInside(r2) {
    println("r2 is inside r1")
}

if r1.IsInside(0.2) {
    println("0.2 is inside the interval")
}
```
### Statistics

Set of statistics functions for golang.

- Calculates histogram's entropy.

```go
entropy := statistics.Entropy([1, 2, 3, 5])
```

## References
[Aforge.Net](https://github.com/andrewkirillov/AForge.NET)