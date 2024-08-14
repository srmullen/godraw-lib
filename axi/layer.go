package axi

type Layer struct {
	Index int
	Name  string
	Pen   string
	Items []Drawer
}

func (a *Axi) NewLayer(name, pen string) *Layer {
	a.layer += 1
	layer := &Layer{
		Index: a.layer,
		Name:  name,
		Pen:   pen,
	}
	a.layers[name] = layer
	a.activeLayer = layer
	return layer
}

func (l *Layer) Draw(item Drawer) {
	l.Items = append(l.Items, item)
}
