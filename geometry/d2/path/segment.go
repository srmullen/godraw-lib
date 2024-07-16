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

func NewCubicBezierSegment(p, ctrlIn, ctrlOut point.Point) Segment {
	curve := NewCubicBezier(ctrlIn, ctrlOut)
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
