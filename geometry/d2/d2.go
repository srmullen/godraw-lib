package d2

type Coords interface {
	Coords() (float64, float64)
}

type CoordData interface {
	Data() []float64
}
