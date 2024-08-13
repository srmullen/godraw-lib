package path

import (
	"github.com/srmullen/godraw-lib/geometry/d2/point"
)

type Segment struct {
	point.Point
	Curve *Curve
}

func NewSegment(x, y float64) Segment {
	return Segment{
		Point: point.Point{
			X: x,
			Y: y,
		},
		Curve: nil,
	}
}

func NewCubicBezierSegment(p, c1, c2 point.Point) Segment {
	curve := NewCubicBezier(c1, c2)
	return Segment{
		Point: p,
		Curve: curve,
	}
}

func (s Segment) Scale(m float64) Segment {
	return Segment{
		Point: s.Point.ScalarMult(m),
		Curve: s.Curve,
	}
}

func (s Segment) Interpolate(to point.Point, t float64) (float64, float64) {
	if s.Curve == nil {
		x := s.X + t*(to.X-s.X)
		y := s.Y + t*(to.Y-s.Y)
		return x, y
	} else {
		// TODO: Need to handle Quadratic and Arc curves
		return s.Curve.Interpolate(s.Point, to, t)
	}
}

func (s Segment) Length(to point.Point) float64 {
	if s.Curve == nil {
		return s.Distance(to)
	} else {
		return s.Curve.Length(s.Point, to)
	}
}
