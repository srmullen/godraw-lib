package line

import (
	"math"

	"github.com/srmullen/godraw-lib/geometry/d2"
)

const precision = 4

// GetVectorIntersection returns the intersection point of two infinite lines
func GetVectorIntersection(x1, y1, x2, y2, x3, y3, x4, y4 float64) (x, y float64, ok bool) {
	// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_two_points_on_each_line
	denom := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if denom == 0 {
		return 0, 0, false
	}

	// x and y are the intersection point of the infinte lines
	// x = d2.RoundToDecimal(((x1*y2-y1*x2)*(x3-x4)-(x1-x2)*(x3*y4-y3*x4))/denom, precision)
	// y = d2.RoundToDecimal(((x1*y2-y1*x2)*(y3-y4)-(y1-y2)*(x3*y4-y3*x4))/denom, precision)
	x = ((x1*y2-y1*x2)*(x3-x4) - (x1-x2)*(x3*y4-y3*x4)) / denom
	y = ((x1*y2-y1*x2)*(y3-y4) - (y1-y2)*(x3*y4-y3*x4)) / denom
	return x, y, true
}

// GetIntersection returns the intersection point of two line segments
func GetIntersection(x1, y1, x2, y2, x3, y3, x4, y4 float64) (x, y float64, ok bool) {
	x, y, ok = GetVectorIntersection(x1, y1, x2, y2, x3, y3, x4, y4)
	if !ok {
		return 0, 0, false
	}

	xr := d2.RoundToDecimal(x, precision)
	yr := d2.RoundToDecimal(y, precision)
	x1r := d2.RoundToDecimal(x1, precision)
	y1r := d2.RoundToDecimal(y1, precision)
	x2r := d2.RoundToDecimal(x2, precision)
	y2r := d2.RoundToDecimal(y2, precision)
	x3r := d2.RoundToDecimal(x3, precision)
	y3r := d2.RoundToDecimal(y3, precision)
	x4r := d2.RoundToDecimal(x4, precision)
	y4r := d2.RoundToDecimal(y4, precision)

	// make sure intersection point is on the line segments
	// This is failing because of floating point errors
	if xr < x1r && xr < x2r || xr > x1r && xr > x2r {
		// log.Println("here1", x, x1, x2)
		return 0, 0, false
	}
	if yr < y1r && yr < y2r || yr > y1r && yr > y2r {
		// log.Println("here2", y, y1, y2)
		return 0, 0, false
	}
	if xr < x3r && xr < x4r || xr > x3r && xr > x4r {
		// log.Println("here3", x, x3, x4)
		return 0, 0, false
	}
	if yr < y3r && yr < y4r || yr > y3r && yr > y4r {
		// log.Println("here4", y, y3, y4)
		return 0, 0, false
	}

	return x, y, true
}

// Standard form
// Ax + By = C
type Line struct {
	x1, y1, x2, y2 float64
}

func (l Line) GetIntersection(l2 Line) (x, y float64, ok bool) {
	return GetIntersection(l.x1, l.y1, l.x2, l.y2, l2.x1, l2.y1, l2.x2, l2.y2)
}

func NewLine(x1, y1, x2, y2 float64) Line {
	return Line{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}
}

func (l Line) Data() []float64 {
	return []float64{l.x1, l.y1, l.x2, l.y2}
}

func (l Line) PointSlopeForm(x float64) float64 {
	m := (l.y2 - l.y1) / (l.x2 - l.x1)
	b := l.y1 - m*l.x1
	return m*x + b
}

func (l Line) M() float64 {
	return (l.y2 - l.y1) / (l.x2 - l.x1)
}

// Implement the Vector interface

// Magnitude return the length of the line
func (l Line) Magnitude() float64 {
	return math.Sqrt(math.Pow(l.x2-l.x1, 2) + math.Pow(l.y2-l.y1, 2))
}

// return -pi to pi
func (l Line) Direction() float64 {
	return math.Atan2(l.y2-l.y1, l.x2-l.x1)
}

func (l Line) Origin() (float64, float64) {
	return l.x1, l.y1
}
