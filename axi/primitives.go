package axi

import (
	"fmt"
)

type Line struct {
	x1, y1, x2, y2 float64
}

func (l Line) Draw(axi *Axi) {
	axi.ctx.Line(int(l.x1), int(l.y1), int(l.x2), int(l.y2), fmt.Sprintf("stroke:%s;stroke-width:%f", axi.pen.Color, axi.pen.Width))
}

type Circle struct {
	x, y, r float64
}

func (c Circle) Draw(axi *Axi) {
	axi.ctx.Circle(int(c.x), int(c.y), int(c.r), fmt.Sprintf("fill:none;stroke:%s;stroke-width:%f", axi.pen.Color, axi.pen.Width))
}

type Rect struct {
	x, y, w, h float64
}

func (r Rect) Draw(axi *Axi) {
	axi.ctx.Rect(int(r.x), int(r.y), int(r.w), int(r.h), fmt.Sprintf("fill:none;stroke:%s;stroke-width:%f", axi.pen.Color, axi.pen.Width))
}

type Path struct {
	data string
}

func (p Path) Draw(axi *Axi) {
	axi.ctx.Path(p.data, fmt.Sprintf("fill:none;stroke:%s;stroke-width:%f", axi.pen.Color, axi.pen.Width))
}
