package axi

import (
	"fmt"
	"io"
	"os"
	"strings"

	svg "github.com/ajstarks/svgo"
)

type Axi struct {
	Width    float64
	Height   float64
	ctx      *svg.SVG
	pens     map[string]*Pen
	pen      *Pen
	layer    int
	position struct {
		X float64
		Y float64
	}
}

type Drawer interface {
	Draw(axi *Axi)
}

func (axi *Axi) Draw(d Drawer) {
	d.Draw(axi)
}

type PathData interface {
	PathData() string
}

func NewAxi(width, height float64) *Axi {
	s := svg.New(os.Stdout)
	s.Start(int(width), int(height), "xmlns:inkscape=\"http://www.inkscape.org/namespaces/inkscape\"")

	return &Axi{
		Width:  width,
		Height: height,
		ctx:    s,
		pens:   make(map[string]*Pen),
	}
}

func NewAxiWithWriter(w io.Writer, width, height float64) *Axi {
	s := svg.New(w)
	s.Start(int(width), int(height), "xmlns:inkscape=\"http://www.inkscape.org/namespaces/inkscape\"")

	return &Axi{
		Width:  width,
		Height: height,
		ctx:    s,
		pens:   make(map[string]*Pen),
	}
}

func (axi *Axi) NewPen(name, color string, width float64) *Pen {
	pen := newPen(name, color, width)
	axi.pens[name] = pen
	return pen
}

func (axi *Axi) WithPen(name string) {
	if axi.pen != nil && axi.pen.Name == name {
		return
	}
	if axi.pen != nil {
		axi.UnloadPen()
	}
	axi.pen = axi.pens[name]
	axi.layer += 1
	attrs := []string{
		"inkscape:groupmode=\"layer\"",
		fmt.Sprintf("id=\"layer%d\"", axi.layer),
		fmt.Sprintf("inkscape:label=\"%d-%s\"", axi.layer, axi.pen.Name),
	}
	axi.ctx.Group(strings.Join(attrs, " "))
}

func (axi *Axi) UnloadPen() {
	axi.pen = nil
	axi.ctx.Gend()
}

func (axi *Axi) Done() {
	if axi.pen != nil {
		axi.UnloadPen()
	}
	axi.ctx.End()
}

func (axi *Axi) Pen() *Pen {
	return axi.pen
}

// Drawing functions

func (axi *Axi) Line(x1, y1, x2, y2 float64) {
	axi.ctx.Line(int(x1), int(y1), int(x2), int(y2), fmt.Sprintf("stroke:%s;stroke-width:%f", axi.pen.Color, axi.pen.Width))
}

func (axi *Axi) Circle(x, y, r float64) {
	axi.ctx.Circle(int(x), int(y), int(r), fmt.Sprintf("fill:none;stroke:%s;stroke-width:%f", axi.pen.Color, axi.pen.Width))
}

func (axi *Axi) Rect(x, y, w, h float64) {
	axi.ctx.Rect(int(x), int(y), int(w), int(h), fmt.Sprintf("fill:none;stroke:%s;stroke-width:%f", axi.pen.Color, axi.pen.Width))
}

func (axi *Axi) SVGPath(path string) {
	axi.ctx.Path(path, fmt.Sprintf("fill:none;stroke:%s;stroke-width:%f", axi.pen.Color, axi.pen.Width))
}

func (axi *Axi) Path(p PathData) {
	axi.SVGPath(p.PathData())
}

func (axi *Axi) Paths(p []PathData) {
	for _, path := range p {
		axi.Path(path)
	}
}

func (axi *Axi) MoveTo(x, y float64) {
	axi.position.X = x
	axi.position.Y = y
}

func (axi *Axi) LineTo(x, y float64) {
	axi.ctx.Line(int(axi.position.X), int(axi.position.Y), int(x), int(y), fmt.Sprintf("stroke:%s;stroke-width:%f", axi.pen.Color, axi.pen.Width))
	axi.position.X = x
	axi.position.Y = y
}
