package axi

type Pen struct {
	Name  string
	Color string
	Width float64
}

func NewPen(name, color string, width float64) *Pen {
	return &Pen{
		Name:  name,
		Color: color,
		Width: width,
	}
}

// func (pen *Pen) Line(x1, y1, x2, y2 float64) {
// 	pen.ctx.Line(
// 		int(math.Round(x1)),
// 		int(math.Round(y1)),
// 		int(math.Round(x2)),
// 		int(math.Round(y2)),
// 		"stroke:"+pen.Color+";stroke-width:"+fmt.Sprintf("%f", pen.Width),
// 	)
// }
