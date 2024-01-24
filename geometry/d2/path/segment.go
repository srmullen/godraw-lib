package path

import "github.com/srmullen/godraw-lib/geometry/d2/point"

type Segment struct {
	point.Point
	Curve *Curve
}

func (s *Segment) Scale(m float64) *Segment {
	return &Segment{
		Point: *s.Point.ScalarMult(m),
		// Curve: s.Curve.Scale(m),
		Curve: s.Curve,
	}
}
