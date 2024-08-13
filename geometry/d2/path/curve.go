package path

import (
	"fmt"
	"math"

	"github.com/srmullen/godraw-lib/geometry/d2/point"
)

func CubicBezierPolynomialInterpolation(p1, p2, p3, p4 point.Point, t float64) (float64, float64) {
	x := p1.X + t*(-3*p1.X+3*p2.X) + t*t*(3*p1.X+-6*p2.X+3*p3.X) + t*t*t*(-p1.X+3*p2.X+-3*p3.X+p4.X)
	y := p1.Y + t*(-3*p1.Y+3*p2.Y) + t*t*(3*p1.Y+-6*p2.Y+3*p3.Y) + t*t*t*(-p1.Y+3*p2.Y+-3*p3.Y+p4.Y)
	return x, y
}

func CubicBezierLength(c1, c2, c3, c4 point.Point) float64 {
	epsilon := 1e-6
	chord := c1.Distance(c4)
	control := c1.Distance(c2) + c2.Distance(c3) + c3.Distance(c4)
	if control-chord <= epsilon {
		return (chord + control) / 2
	}
	left, right := Subdivide([]point.Point{c1, c2, c3, c4}, 0.5)
	return CubicBezierLength(left[0], left[1], left[2], left[3]) + CubicBezierLength(right[0], right[1], right[2], right[3])
}

func Subdivide(b []point.Point, t float64) ([]point.Point, []point.Point) {
	p01 := interpolate(b[0], b[1], t)
	p12 := interpolate(b[1], b[2], t)
	p23 := interpolate(b[2], b[3], t)
	p012 := interpolate(p01, p12, t)
	p123 := interpolate(p12, p23, t)
	p0123 := interpolate(p012, p123, t)

	return []point.Point{b[0], p01, p012, p0123},
		[]point.Point{p0123, p123, p23, b[3]}
}

func interpolate(p1, p2 point.Point, t float64) point.Point {
	return point.NewPoint(
		p1.X*(1-t)+p2.X*t,
		p1.Y*(1-t)+p2.Y*t,
	)
}

type Curve struct {
	*CubicBezier
	*QuadraticBezier
	*Arc
}

// TODO: Curve pathdata should be relative units
// I think that will make it easier to translate curves
func (c *Curve) PathData() string {
	if c.CubicBezier != nil {
		return c.CubicBezier.PathData()
	} else if c.QuadraticBezier != nil {
		return c.QuadraticBezier.PathData()
	} else if c.Arc != nil {
		return c.Arc.PathData()
	}
	return ""
}

func (c *Curve) Translate(x, y float64) *Curve {
	if c.CubicBezier != nil {
		return NewCubicBezier(
			point.Point{
				X: c.CubicBezier.C1.X + x,
				Y: c.CubicBezier.C1.Y + y,
			},
			point.Point{
				X: c.CubicBezier.C2.X + x,
				Y: c.CubicBezier.C2.Y + y,
			},
		)
	} else if c.QuadraticBezier != nil {
		control := point.Point{
			X: c.QuadraticBezier.C.X + x,
			Y: c.QuadraticBezier.C.Y + y,
		}
		return NewQuadraticBezier(control)
	} else if c.Arc != nil {
		return nil
	}
	return nil
}

func (c *Curve) Interpolate(p1, p2 point.Point, t float64) (float64, float64) {
	if c.CubicBezier != nil {
		return c.CubicBezier.Interpolate(p1, p2, t)
	} else if c.QuadraticBezier != nil {
		return c.QuadraticBezier.Interpolate(p1, p2, t)
	} else if c.Arc != nil {
		return c.Arc.Interpolate(p1, p2, t)
	}
	return 0, 0
}

func (c *Curve) Length(p1, p2 point.Point) float64 {
	if c.CubicBezier != nil {
		// return c.CubicBezier.Length(p1, p2)
		return CubicBezierLength(p1, c.CubicBezier.C1, c.CubicBezier.C2, p2)
	} else if c.QuadraticBezier != nil {
		// return c.QuadraticBezier.Length(p1, p2)
		return 0
	} else if c.Arc != nil {
		// return c.Arc.Length(p1, p2)
		return 0
	}
	return 0
}

type CubicBezier struct {
	C1 point.Point
	C2 point.Point
}

func NewCubicBezier(c1, c2 point.Point) *Curve {
	return &Curve{
		CubicBezier: &CubicBezier{
			C1: c1,
			C2: c2,
		},
	}
}

func (c *CubicBezier) PathData() string {
	ret := "C"
	ret += fmt.Sprintf("%d %d ", int(math.Round(c.C1.X)), int(math.Round(c.C1.Y)))
	ret += fmt.Sprintf("%d %d ", int(math.Round(c.C2.X)), int(math.Round(c.C2.Y)))
	return ret
}

func (c *CubicBezier) Interpolate(p1, p2 point.Point, t float64) (float64, float64) {
	return CubicBezierPolynomialInterpolation(p1, c.C1, c.C2, p2, t)
}

type QuadraticBezier struct {
	C point.Point
}

func NewQuadraticBezier(c point.Point) *Curve {
	return &Curve{
		QuadraticBezier: &QuadraticBezier{
			C: c,
		},
	}
}

func (q *QuadraticBezier) PathData() string {
	ret := "Q"
	ret += fmt.Sprintf("%d %d ", int(math.Round(q.C.X)), int(math.Round(q.C.Y)))
	return ret
}

// TODO: Written by copilot. Need to verify
func (q *QuadraticBezier) Interpolate(p1, p2 point.Point, t float64) (float64, float64) {
	x := (1-t)*(1-t)*p1.X + 2*(1-t)*t*q.C.X + t*t*p2.X
	y := (1-t)*(1-t)*p1.Y + 2*(1-t)*t*q.C.Y + t*t*p2.Y
	return x, y
}

type Arc struct {
	Rx, Ry float64
	Xrot   float64
	Large  bool
	Sweep  bool
}

func NewArc(rx, ry, xrot float64, large, sweep bool) *Curve {
	return &Curve{
		Arc: &Arc{
			Rx:    rx,
			Ry:    ry,
			Xrot:  xrot,
			Large: large,
			Sweep: sweep,
		},
	}
}

func (a *Arc) PathData() string {
	ret := "A"
	ret += fmt.Sprintf("%d %d ", int(math.Round(a.Rx)), int(math.Round(a.Ry)))
	ret += fmt.Sprintf("%d ", int(math.Round(a.Xrot)))
	if a.Large {
		ret += "1 "
	} else {
		ret += "0 "
	}
	if a.Sweep {
		ret += "1 "
	} else {
		ret += "0 "
	}
	return ret
}

// TODO: Needs implementation. Need to figure out how arc is calculated.
func (a *Arc) Interpolate(p1, p2 point.Point, t float64) (float64, float64) {
	return 0, 0
}

func (c1 *Curve) GetIntersections(c2 *Curve) {

}
