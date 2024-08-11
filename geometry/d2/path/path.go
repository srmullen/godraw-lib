package path

import (
	"fmt"
	"log"
	"math"

	"github.com/srmullen/godraw-lib/geometry/d2/bounds"
	"github.com/srmullen/godraw-lib/geometry/d2/line"
	"github.com/srmullen/godraw-lib/geometry/d2/point"
)

// https://www.w3.org/TR/SVG11/paths.html
type PathData interface {
	PathData() string
}

// A Path is an ordered collection of 2d points
type Path struct {
	Segments []Segment
	Closed   bool
}

// TODO: Implement
func FromSegments(segments []Segment, closed bool) *Path {
	// return NewPath([]float64{}, closed)
	return &Path{
		segments,
		closed,
	}
}

func NewPath(coords []float64, closed bool) *Path {
	segments := make([]Segment, 0)
	for i := 0; i < len(coords); i += 2 {
		x := coords[i]
		y := coords[i+1]
		segments = append(segments, Segment{
			Point: point.NewPoint(x, y),
			Curve: nil,
		})
	}
	return &Path{
		Segments: segments,
		Closed:   closed,
	}
}

func NewOpenPath(coords []float64) *Path {
	return NewPath(coords, false)
}

func NewClosedPath(coords []float64) *Path {
	return NewPath(coords, true)
}

func (p *Path) Segment(i int) Segment {
	return p.Segments[i]
}

func (p *Path) PathData() string {
	ret := ""

	for i, segment := range p.Segments {
		if i == 0 {
			// Move to: starts a new path
			ret += "M"
			ret += fmt.Sprintf("%d %d ", int(math.Round(segment.X)), int(math.Round(segment.Y)))
		} else {
			ret += fmt.Sprintf(" %d %d ", int(math.Round(segment.X)), int(math.Round(segment.Y)))
		}

		if segment.Curve != nil {
			ret += segment.Curve.PathData()
		} else {
			// Line to: continues the path
			ret += "L"
			ret += fmt.Sprintf("%d %d", int(math.Round(segment.X)), int(math.Round(segment.Y)))
		}
	}
	if p.Closed {
		// Close path: draws a line from the last point to the first point
		ret += "Z"
	}
	return ret
}

func PathTranslate(path *Path, x, y float64) *Path {
	ret := &Path{
		Segments: make([]Segment, len(path.Segments)),
		Closed:   path.Closed,
	}
	for i, segment := range path.Segments {
		var curve *Curve = nil
		if segment.Curve != nil {
			curve = segment.Curve.Translate(x, y)
		}
		ret.Segments[i] = Segment{
			point.Point{
				X: segment.X + x,
				Y: segment.Y + y,
			},
			curve,
		}
	}
	return ret
}

// Bounds returns the bounding box of the path
// TODO: This does not take curves into account.a
func (p *Path) GetBounds() *bounds.Bounds {
	if p == nil || len(p.Segments) == 0 {
		return nil
	}
	minX := p.Segments[0].X
	minY := p.Segments[0].Y
	maxX := p.Segments[0].X
	maxY := p.Segments[0].Y
	for _, segment := range p.Segments {
		if segment.X < minX {
			minX = segment.X
		}
		if segment.Y < minY {
			minY = segment.Y
		}
		if segment.X > maxX {
			maxX = segment.X
		}
		if segment.Y > maxY {
			maxY = segment.Y
		}
	}
	return &bounds.Bounds{
		Top:    minY,
		Right:  maxX,
		Bottom: maxY,
		Left:   minX,
	}
}

func ScaleSegments(segments []Segment, scalex, scaley float64) []Segment {
	ret := make([]Segment, len(segments))
	for i, segment := range segments {
		ret[i] = Segment{
			Point: segment.Hadamard(point.Point{X: scalex, Y: scaley}),
			Curve: segment.Curve,
		}
	}
	return ret
}

func (path *Path) Translate(x, y float64) *Path {
	ret := &Path{
		Segments: make([]Segment, len(path.Segments)),
		Closed:   path.Closed,
	}
	for i, segment := range path.Segments {
		var curve *Curve = nil
		if segment.Curve != nil {
			curve = segment.Curve.Translate(x, y)
		}
		ret.Segments[i] = Segment{
			point.Point{
				X: segment.X + x,
				Y: segment.Y + y,
			},
			curve,
		}
	}
	return ret
}

func (p *Path) Scale(scalex, scaley float64) *Path {
	return &Path{
		Segments: ScaleSegments(p.Segments, scalex, scaley),
		Closed:   p.Closed,
	}
}

func (p *Path) Xs() []float64 {
	ret := make([]float64, len(p.Segments))
	for i, segment := range p.Segments {
		ret[i] = segment.X
	}
	return ret
}

func (p *Path) Ys() []float64 {
	ret := make([]float64, len(p.Segments))
	for i, segment := range p.Segments {
		ret[i] = segment.Y
	}
	return ret
}

func (p *Path) Length() float64 {
	ret := 0.0
	for i := 1; i < len(p.Segments); i++ {
		ret += p.Segments[i-1].Distance(p.Segments[i].Point)
	}
	return ret
}

// FIXME: This is a naive implementation that only works for straight lines
func (p *Path) GetIntersections(other *Path) []point.Point {
	var intersections []point.Point
	// If bounds don't intersect, then there are no intersections
	if !p.GetBounds().Overlaps(*other.GetBounds()) {
		return intersections
	}
	for i := 0; i < len(p.Segments); i++ {
		from := p.Segments[i].Point
		to := p.Segments[(i+1)%len(p.Segments)].Point
		for j := 0; j < len(other.Segments); j++ {
			log.Println("i", i, "j", j)
			jFrom := other.Segments[j].Point
			jTo := other.Segments[(j+1)%len(other.Segments)].Point
			// intersection := segment.Point.Intersection(other.Segments[j].Point)
			// intersection := line.NewLine(from.X, from.Y, to.X, to.Y).Intersection(line.NewLine(jFrom.X, jFrom.Y, jTo.X, jTo.Y))
			if x, y, ok := line.GetIntersection(from.X, from.Y, to.X, to.Y, jFrom.X, jFrom.Y, jTo.X, jTo.Y); ok {
				intersections = append(intersections, point.NewPoint(x, y))
			}
		}
	}
	return intersections
}

func (p *Path) Points() []point.Point {
	ret := make([]point.Point, len(p.Segments))
	for i, segment := range p.Segments {
		ret[i] = segment.Point
	}
	return ret
}

// Interpolate returns the x, y coordinates of the point at t along the path
// t is a value between 0 and len(p.Segments)
// the number to the left of the decimal point is the index of the segment
// the number to the right of the decimal point is the interpolation value
func (p *Path) Interpolate(t float64) (float64, float64) {
	if t < 0 || t > float64(len(p.Segments)) {
		return 0, 0
	}
	segmentIndex := int(t)
	segment := p.Segments[segmentIndex]
	next := p.Segments[(segmentIndex+1)%len(p.Segments)]
	v := t - float64(segmentIndex)
	return segment.Interpolate(next.Point, v)
}
