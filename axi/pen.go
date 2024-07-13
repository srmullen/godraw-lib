package axi

type Pen struct {
	Name  string
	Color string
	Width float64
}

func newPen(name, color string, width float64) *Pen {
	return &Pen{
		Name:  name,
		Color: color,
		Width: width,
	}
}
