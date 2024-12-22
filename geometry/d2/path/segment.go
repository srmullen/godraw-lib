package path

import (
	"github.com/srmullen/godraw-lib/geometry/d2/line"
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

// dx, dy is the point the segment ends at
// x1, y1, x2, y2 is the line to find segments with
func (s Segment) LineIntersections(dx, dy, x1, y1, x2, y2 float64) []point.Point {
	ret := make([]point.Point, 0)
	if s.Curve == nil {
		// Just a line
		if x, y, ok := line.GetIntersection(s.X, s.Y, dx, dy, x1, y1, x2, y2); ok {
			ret = append(ret, point.NewPoint(x, y))
		}
	} else {
		// TODO: Need to handle curves
	}
	return ret
}

func (s Segment) BoundIntersections(toX, toY, top, right, bottom, left float64) []point.Point {
	intersections := s.LineIntersections(toX, toY, left, top, right, top)
	intersections = append(intersections, s.LineIntersections(toX, toY, right, top, right, bottom)...)
	intersections = append(intersections, s.LineIntersections(toX, toY, right, bottom, left, bottom)...)
	intersections = append(intersections, s.LineIntersections(toX, toY, left, bottom, left, top)...)
	return intersections
}
