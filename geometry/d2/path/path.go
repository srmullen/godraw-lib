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
	Segments []*Segment
	Closed   bool
}

func NewPath(coords []*point.Point, closed bool) *Path {
	segments := make([]*Segment, len(coords))
	for i, coord := range coords {
		x, y := coord.Coords()
		segments[i] = &Segment{
			Point: point.Point{
				X: x,
				Y: y,
			},
			Curve: nil,
		}
	}
	return &Path{
		Segments: segments,
		Closed:   closed,
	}
}

func NewOpenPath(coords []*point.Point) *Path {
	return NewPath(coords, false)
}

func NewClosedPath(coords []*point.Point) *Path {
	return NewPath(coords, true)
}

func NewLine(x1, y1, x2, y2 float64) *Path {
	return &Path{
		Segments: []*Segment{
			{
				point.Point{
					X: x1,
					Y: y1,
				},
				nil,
			},
			{
				point.Point{
					X: x2,
					Y: y2,
				},
				nil,
			},
		},
		Closed: false,
	}
}

func (p *Path) Segment(i int) *Segment {
	return p.Segments[i]
}

func (p *Path) PathData() string {
	ret := ""
	for i, segment := range p.Segments {
		if i == 0 {
			// Move to: starts a new path
			ret += "M"
			ret += fmt.Sprintf("%d %d", int(math.Round(segment.X)), int(math.Round(segment.Y)))
		} else if segment.Curve != nil {
			ret += segment.Curve.PathData()
			ret += fmt.Sprintf("%d %d", int(math.Round(segment.X)), int(math.Round(segment.Y)))
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
		Segments: make([]*Segment, len(path.Segments)),
		Closed:   path.Closed,
	}
	for i, segment := range path.Segments {
		var curve *Curve = nil
		if segment.Curve != nil {
			curve = segment.Curve.Translate(x, y)
		}
		ret.Segments[i] = &Segment{
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

func ScaleSegments(segments []*Segment, scalex, scaley float64) []*Segment {
	ret := make([]*Segment, len(segments))
	for i, segment := range segments {
		ret[i] = &Segment{
			Point: *segment.Hadamard(&point.Point{X: scalex, Y: scaley}),
			Curve: segment.Curve,
		}
	}
	return ret
}

func (path *Path) Translate(x, y float64) *Path {
	ret := &Path{
		Segments: make([]*Segment, len(path.Segments)),
		Closed:   path.Closed,
	}
	for i, segment := range path.Segments {
		var curve *Curve = nil
		if segment.Curve != nil {
			curve = segment.Curve.Translate(x, y)
		}
		ret.Segments[i] = &Segment{
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
		ret += p.Segments[i-1].Distance(&p.Segments[i].Point)
	}
	return ret
}

// FIXME: This is a naive implementation that only works for straight lines
func (p *Path) GetIntersections(other *Path) []*point.Point {
	var intersections []*point.Point
	// If bounds don't intersect, then there are no intersections
	if !p.GetBounds().Overlaps(other.GetBounds()) {
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
