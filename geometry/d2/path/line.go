package path

type Line struct {
	x1, y1, x2, y2 float64
	Path
}

func NewLine(x1, y1, x2, y2 float64) *Line {
	s1 := NewSegment(x1, y1)
	s2 := NewSegment(x2, y2)
	path := Path{
		Segments: []Segment{
			s1,
			s2,
		},
		Closed: false,
	}
	return &Line{
		x1:   x1,
		y1:   y1,
		x2:   x2,
		y2:   y2,
		Path: path,
	}
}
