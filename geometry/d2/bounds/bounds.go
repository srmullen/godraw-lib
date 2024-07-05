package bounds

// import "godraw/lib/geometry/d2/point"
import (
	"github.com/srmullen/godraw-lib/geometry/d2"
	"github.com/srmullen/godraw-lib/geometry/d2/point"
)

type Bounds struct {
	Top    float64
	Right  float64
	Bottom float64
	Left   float64
}

func (bounds *Bounds) Width() float64 {
	return bounds.Right - bounds.Left
}

func (bounds *Bounds) Height() float64 {
	return bounds.Bottom - bounds.Top
}

func (b *Bounds) TopLeft() *point.Point {
	return &point.Point{X: b.Left, Y: b.Top}
}

func (b *Bounds) TopRight() *point.Point {
	return &point.Point{X: b.Right, Y: b.Top}
}

func (b *Bounds) BottomLeft() *point.Point {
	return &point.Point{X: b.Left, Y: b.Bottom}
}

func (b *Bounds) BottomRight() *point.Point {
	return &point.Point{X: b.Right, Y: b.Bottom}
}

func (b *Bounds) Translate(x, y float64) {
	b.Left += x
	b.Right += x
	b.Top += y
	b.Bottom += y
}

// Center returns the center point of the bounds
func (b *Bounds) Center() *point.Point {
	return &point.Point{
		X: b.Left + b.Width()/2,
		Y: b.Top + b.Height()/2,
	}
}

// Contains returns true if the point passed in is contained in the bounds
func (b *Bounds) ContainsPoint(point *point.Point) bool {
	return point.X > b.Left && point.X < b.Right && point.Y > b.Top && point.Y < b.Bottom
}

func (b *Bounds) Contains(x, y float64) bool {
	return b.ContainsX(x) && b.ContainsY(y)
}

func (b *Bounds) ContainsCoords(coords d2.Coords) bool {
	x, y := coords.Coords()
	return b.ContainsX(x) && b.ContainsY(y)
}

func (b *Bounds) ContainsX(x float64) bool {
	return x >= b.Left && x <= b.Right
}

func (b *Bounds) ContainsY(y float64) bool {
	return y >= b.Top && y <= b.Bottom
}

func (b *Bounds) ContainsInclusive(point *point.Point) bool {
	return point.X >= b.Left && point.X <= b.Right && point.Y >= b.Top && point.Y <= b.Bottom
}

// ContainsBounds returns true if the bounds passed in are completely contained
func (b *Bounds) ContainsBounds(bounds *Bounds) bool {
	return b.ContainsInclusive(bounds.TopLeft()) && b.ContainsInclusive(bounds.BottomRight())
}

func (b *Bounds) Intersects(bounds *Bounds) bool {
	return b.ContainsPoint(bounds.TopLeft()) || b.ContainsPoint(bounds.BottomRight()) || b.ContainsPoint(bounds.TopRight()) || b.ContainsPoint(bounds.BottomLeft())
}

// FIXME: This doesn't capture the situation where the bounds form a cross.
func (b *Bounds) Overlaps(bounds *Bounds) bool {
	// return b.Intersects(bounds) ||
	// 	b.ContainsBounds(bounds) ||
	// 	bounds.ContainsBounds(b) ||
	return bounds.Left+bounds.Width() > b.Left &&
		bounds.Top+bounds.Height() > b.Top &&
		bounds.Left < b.Left+b.Width() &&
		bounds.Top < b.Top+b.Height()
}

func (b *Bounds) ScalePoint(pnt *point.Point) *point.Point {
	return &point.Point{
		X: b.Left + pnt.X*b.Width(),
		Y: b.Top + pnt.Y*b.Height(),
	}
}
