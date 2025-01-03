package point

// A Point that also has a T value from the range 0 to 1 indicating
// That the point is interpolated along some scale.
type InterpolationPoint struct {
	T float64
	Point
}

func NewInterpolationPoint(x, y, t float64) InterpolationPoint {
	return InterpolationPoint{
		T:     t,
		Point: NewPoint(x, y),
	}
}
