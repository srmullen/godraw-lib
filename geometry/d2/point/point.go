package point

import (
	"math"

	"github.com/srmullen/godraw-lib/geometry/d2"
)

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

// NewPointFromAngle creates a new point from an angle (radians) and magnitude
func NewPointFromAngle(radians, magnitude float64) *Point {
	return &Point{
		magnitude * math.Cos(radians),
		magnitude * math.Sin(radians),
	}
}

func (p *Point) Rotate(radians float64) *Point {
	return &Point{
		p.X*math.Cos(radians) - p.Y*math.Sin(radians),
		p.X*math.Sin(radians) + p.Y*math.Cos(radians),
	}
}

// Returns the angle of the vector in radians
func (p *Point) Angle() float64 {
	return math.Atan2(p.Y, p.X)
}

func (p *Point) Coords() (float64, float64) {
	return p.X, p.Y
}

func (p *Point) Equals(other *Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p *Point) EqualsWithTolerance(other *Point, tolerance float64) bool {
	return d2.WithinTolerance(p.X, other.X, tolerance) && d2.WithinTolerance(p.Y, other.Y, tolerance)
}

func (p1 *Point) Add(x, y float64) *Point {
	return &Point{
		p1.X + x,
		p1.Y + y,
	}
}

func (p1 *Point) Subtract(x, y float64) *Point {
	return &Point{
		p1.X - x,
		p1.Y - y,
	}
}

func (p1 *Point) Divide(x, y float64) *Point {
	return &Point{
		p1.X / x,
		p1.Y / y,
	}
}

func (p1 *Point) ScalarMult(s float64) *Point {
	return &Point{
		p1.X * s,
		p1.Y * s,
	}
}

func (p1 *Point) AddPoint(p2 *Point) *Point {
	return &Point{
		p1.X + p2.X,
		p1.Y + p2.Y,
	}
}

func (p1 *Point) SubtractPoint(p2 *Point) *Point {
	return &Point{
		p1.X - p2.X,
		p1.Y - p2.Y,
	}
}

func (p1 *Point) DividePoint(p2 *Point) *Point {
	return &Point{
		p1.X / p2.X,
		p1.Y / p2.Y,
	}
}

func (p1 *Point) Magnitude() float64 {
	return math.Sqrt(p1.X*p1.X + p1.Y*p1.Y)
}

func (p1 *Point) Distance(p2 *Point) float64 {
	return p1.SubtractPoint(p2).Magnitude()
}

func (p1 *Point) Normalize() *Point {
	return p1.ScalarMult(1 / p1.Magnitude())
}

func (p1 *Point) Dot(p2 *Point) float64 {
	return p1.X*p2.X + p1.Y*p2.Y
}

func (p1 *Point) Cross(p2 *Point) *Point {
	return &Point{
		p1.X*p2.Y - p1.Y*p2.X,
		p1.X*p2.Y - p1.Y*p2.X,
	}
}

func (p1 *Point) Hadamard(p2 *Point) *Point {
	return &Point{
		p1.X * p2.X,
		p1.Y * p2.Y,
	}
}
