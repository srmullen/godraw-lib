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
		var control *point.Point = nil
		if c.QuadraticBezier.C != nil {
			control = &point.Point{
				X: c.QuadraticBezier.C.X + x,
				Y: c.QuadraticBezier.C.Y + y,
			}
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
	C *point.Point
}

func NewQuadraticBezier(c *point.Point) *Curve {
	return &Curve{
		QuadraticBezier: &QuadraticBezier{
			C: c,
		},
	}
}

func (q *QuadraticBezier) PathData() string {
	if q.C == nil {
		return "T"
	}
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
