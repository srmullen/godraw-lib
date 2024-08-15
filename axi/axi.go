package axi

import (
	"fmt"
	"io"
	"os"
	"strings"

	svg "github.com/ajstarks/svgo"
)

type Axi struct {
	Width       float64
	Height      float64
	ctx         *svg.SVG
	pens        map[string]*Pen
	pen         *Pen
	layers      map[string]*Layer
	activeLayer *Layer
	layer       int
	position    struct {
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
	axi := &Axi{
		Width:  width,
		Height: height,
		pens:   make(map[string]*Pen),
		layers: make(map[string]*Layer),
	}

	// Create default pen and layer
	axi.NewPen("default", "black", 1)
	axi.NewLayer("default", "default")

	return axi
}

func (axi *Axi) Done() {
	ctx := svg.New(os.Stdout)
	ctx.Start(int(axi.Width), int(axi.Height), "xmlns:inkscape=\"http://www.inkscape.org/namespaces/inkscape\"")

	axi.ctx = ctx

	// Iterate over layers and render the items they contain
	for _, layer := range axi.layers {
		// Create the Group for the layer
		if len(layer.Items) == 0 {
			continue
		}
		axi.WithPen(layer.Pen)
		attrs := []string{
			"inkscape:groupmode=\"layer\"",
			fmt.Sprintf("id=\"layer%d\"", layer.Index),
			fmt.Sprintf("inkscape:label=\"%d-%s\"", layer.Index, axi.pen.Name),
		}
		axi.ctx.Group(strings.Join(attrs, " "))
		// Render the items
		for _, item := range layer.Items {
			item.Draw(axi)
		}
		axi.ctx.Gend()
	}
	axi.ctx.End()
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

// Creates a Layer and Pen with the same name and associates them
func (axi *Axi) NewPenLayer(name, color string, width float64) {
	pen := axi.NewPen(name, color, width)
	axi.NewLayer(name, pen.Name)
}

func (axi *Axi) OnLayer(name string) {
	axi.activeLayer = axi.layers[name]
}

func (axi *Axi) WithPen(name string) {
	if axi.pen != nil && axi.pen.Name == name {
		return
	}
	axi.pen = axi.pens[name]
}

func (axi *Axi) Pen() *Pen {
	return axi.pen
}

// Drawing functions

func (axi *Axi) Line(x1, y1, x2, y2 float64) {
	item := Line{x1, y1, x2, y2}
	axi.drawItem(item)
}

func (axi *Axi) Circle(x, y, r float64) {
	item := Circle{x, y, r}
	axi.drawItem(item)
}

func (axi *Axi) Rect(x, y, w, h float64) {
	item := Rect{x, y, w, h}
	axi.drawItem(item)
}

func (axi *Axi) drawItem(item Drawer) {
	if axi.ctx == nil {
		axi.activeLayer.Draw(item)
	} else {
		// Rendering Phase
		item.Draw(axi)
	}
}

func (axi *Axi) Path(p PathData) {
	path := Path{p.PathData()}
	axi.drawItem(path)
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
