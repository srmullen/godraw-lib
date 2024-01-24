package lib

import (
	"math"
)

type Unit struct {
	Name   string
	System string
	Factor float64
}

type anchor struct {
	unit  *Unit
	ratio float64
}

var (
	MM = &Unit{"mm", "metric", 1. / 1000}
	CM = &Unit{"cm", "metric", 1. / 100}
	M  = &Unit{"m", "metric", 1}
	PC = &Unit{"pc", "imperial", 1. / 6}
	PT = &Unit{"pt", "imperial", 1. / 72}
	IN = &Unit{"in", "imperial", 1}
	FT = &Unit{"ft", "imperial", 12}
	PX = &Unit{"px", "imperial", 96} // 96 is pixelsPerInch
)

func (u *Unit) WithPixelsPerInch(ppi float64) *Unit {
	return &Unit{
		u.Name,
		u.System,
		ppi,
	}
}

var anchors = map[string]anchor{
	"metric":   {M, 1 / 0.0254},
	"imperial": {IN, 0.0254},
}

type Length struct {
	Unit  *Unit
	Value float64
}

func (from *Length) To(to *Unit) *Length {
	if from.Unit == to {
		return from
	}

	fromFactor := 1.
	toFactor := 1.

	fromUnit := from.Unit
	toUnit := to
	if from.Unit.Name == PX.Name {
		fromFactor = 1. / from.Unit.Factor
		fromUnit = IN
	}
	if to.Name == PX.Name {
		toFactor = to.Factor
		toUnit = IN
	}

	anchor := from.Value * fromUnit.Factor * fromFactor
	if fromUnit.System != toUnit.System {
		anchor = anchor * anchors[from.Unit.System].ratio
	}

	result := anchor / toUnit.Factor * toFactor
	if to.Name == PX.Name {
		result = math.Round(result)
	}

	return &Length{
		to,
		result,
	}
}
