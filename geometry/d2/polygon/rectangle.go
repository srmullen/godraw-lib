package polygon

import (
	"github.com/srmullen/godraw-lib/geometry/d2/path"
	"github.com/srmullen/godraw-lib/geometry/d2/point"
)

type Rectangle struct {
	*Polygon
	x      float64
	y      float64
	width  float64
	height float64
}

func NewRectangle(x, y, width, height float64) *Rectangle {
	poly := &Polygon{
		Path: path.NewPath([]*point.Point{
			{X: x, Y: y},
			{X: x + width, Y: y},
			{X: x + width, Y: y + height},
			{X: x, Y: y + height},
		}, true),
	}
	return &Rectangle{
		poly,
		x,
		y,
		width,
		height,
	}
}

// NewRectangleFromCenter creates a rectangle with the center at x, y
func NewRectangleFromCenter(centerx, centery, width, height float64) *Rectangle {
	x := centerx - width/2
	y := centery - height/2
	return NewRectangle(x, y, width, height)
}

func (r *Rectangle) X() float64 {
	return r.x
}

func (r *Rectangle) Y() float64 {
	return r.y
}

func (r *Rectangle) Width() float64 {
	return r.width
}

func (r *Rectangle) Height() float64 {
	return r.height
}

func (r *Rectangle) Center() *point.Point {
	return point.NewPoint(r.x+r.width/2, r.y+r.height/2)
}

func (r *Rectangle) Contains(x, y float64) bool {
	return r.x <= x && x <= r.x+r.width && r.y <= y && y <= r.y+r.height
}
