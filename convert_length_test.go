package lib

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	inch := Length{IN, 1}
	converted := inch.To(PX)
	assert.Equal(t, PX, converted.Unit)
	assert.Equal(t, 96., converted.Value)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func TestPXtoPhysical(t *testing.T) {
	// t.equal(convert(1, 'in', 'px'), 96);
	// t.equal(convert(1, 'px', 'in'), 1 / 96);
	// t.equal(convert(4, 'in', 'px'), 384);
	// t.equal(convert(1, 'in', 'px', { pixelsPerInch: 72 }), 72);
	// t.equal(convert(6, 'cm', 'px'), 227);
	// t.equal(convert(10, 'mm', 'px'), 38);
	// t.equal(convert(10, 'mm', 'px', { pixelsPerInch: 72 }), 28);
	// t.equal(convert(10, 'px', 'mm', { precision: 2 }), 2.65);
	// t.equal(convert(11, 'px', 'in', { precision: 3 }), 0.115);

	assert.Equal(t, 96., (&Length{IN, 1}).To(PX).Value)
	assert.Equal(t, 1./96, (&Length{PX, 1}).To(IN).Value)
	assert.Equal(t, 384., (&Length{IN, 4}).To(PX).Value)
	assert.Equal(t, 72., (&Length{IN, 1}).To(PX.WithPixelsPerInch(72)).Value) // different pixelsPerInch
	assert.Equal(t, 227., (&Length{CM, 6}).To(PX).Value)
	assert.Equal(t, 38., (&Length{MM, 10}).To(PX).Value)
	assert.Equal(t, 28., (&Length{MM, 10}).To(PX.WithPixelsPerInch(72)).Value)
	assert.Equal(t, 2.6458, toFixed((&Length{PX, 10}).To(MM).Value, 4))
	assert.Equal(t, 0.1146, toFixed((&Length{PX, 11}).To(IN).Value, 4))
}

func TestBasicConversion(t *testing.T) {
	// t.equal(convert(1, 'm', 'm'), 1);
	// t.equal(convert(1, 'cm', 'm'), 0.01);
	// t.equal(convert(1, 'mm', 'm'), 0.001);
	// t.equal(convert(1, 'm', 'cm'), 100);
	// t.equal(convert(1, 'm', 'mm'), 1000);
	// t.equal(convertWithPrecision(1, 'm', 'ft'), 3.2808);
	// t.equal(convert(1, 'ft', 'in'), 12);
	// t.equal(convert(1, 'in', 'ft'), 1 / 12);
	// t.equal(convertWithPrecision(1, 'm', 'ft'), 3.2808);
	// t.equal(convertWithPrecision(1, 'm', 'ft'), 3.2808);
	// t.equal(convertWithPrecision(1, 'm', 'in'), 39.3701);
	// t.equal(convert(1, 'in', 'm'), 0.0254);
	// t.equal(convert(1, 'cm', 'm'), 0.01);
	// t.equal(convert(1, 'm', 'cm'), 100);
	// t.equal(convert(72, 'pt', 'in'), 1);
	// t.equal(convert(1, 'in', 'pt'), 72);
	// t.equal(convert(1, 'in', 'pc'), 6);
	// t.equal(convert(6, 'pc', 'in'), 1);
	// t.equal(convert(6, 'pc', 'pc'), 6);
	// t.equal(convert(6, 'in', 'in'), 6);

	assert.Equal(t, 1., (&Length{M, 1}).To(M).Value)
	assert.Equal(t, 0.01, (&Length{CM, 1}).To(M).Value)
	assert.Equal(t, 0.001, (&Length{MM, 1}).To(M).Value)
	assert.Equal(t, 100., (&Length{M, 1}).To(CM).Value)
	assert.Equal(t, 1000., (&Length{M, 1}).To(MM).Value)
	assert.Equal(t, 3.2808, toFixed((&Length{M, 1}).To(FT).Value, 4))
	assert.Equal(t, 12., (&Length{FT, 1}).To(IN).Value)
	assert.Equal(t, 1./12, (&Length{IN, 1}).To(FT).Value)
	assert.Equal(t, 3.2808, toFixed((&Length{M, 1}).To(FT).Value, 4))
	assert.Equal(t, 39.3701, toFixed((&Length{M, 1}).To(IN).Value, 4))
	assert.Equal(t, 0.0254, toFixed((&Length{IN, 1}).To(M).Value, 4))
	assert.Equal(t, 0.01, (&Length{CM, 1}).To(M).Value)
	assert.Equal(t, 100., (&Length{M, 1}).To(CM).Value)
	assert.Equal(t, 1., (&Length{PT, 72}).To(IN).Value)
	assert.Equal(t, 72., (&Length{IN, 1}).To(PT).Value)
	assert.Equal(t, 6., (&Length{IN, 1}).To(PC).Value)
	assert.Equal(t, 1., (&Length{PC, 6}).To(IN).Value)
	assert.Equal(t, 6., (&Length{PC, 6}).To(PC).Value)
	assert.Equal(t, 6., (&Length{IN, 6}).To(IN).Value)
}
