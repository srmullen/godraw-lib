package path

type Line struct {
	x1, y1, x2, y2 float64
	Path
}

func NewLine(x1, y1, x2, y2 float64) *Line {
	path := Path{
		Segments: []*Segment{
			NewSegment(x1, y1),
			NewSegment(x2, y2),
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
