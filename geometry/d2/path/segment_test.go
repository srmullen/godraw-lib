package path

import (
	"testing"

	"github.com/srmullen/godraw-lib/geometry/d2/point"
	"github.com/stretchr/testify/assert"
)

func TestInterpolate(t *testing.T) {
	t.Run("line interpolation", func(t *testing.T) {
		p := NewPath([]float64{0, 0, 1, 1}, false)

		x, y := p.Interpolate(0)
		assert.Equal(t, x, 0.)
		assert.Equal(t, y, 0.)

		x, y = p.Interpolate(1)
		assert.Equal(t, x, 1.)
		assert.Equal(t, y, 1.)

		x, y = p.Interpolate(0.5)
		assert.Equal(t, x, 0.5)
		assert.Equal(t, y, 0.5)
	})

	t.Run("multiple line segments", func(t *testing.T) {
		p := NewPath([]float64{0, 0, 1, 1, 3, 1, 3, 3}, false)

		x, y := p.Interpolate(0)
		assert.Equal(t, x, 0.)
		assert.Equal(t, y, 0.)

		x, y = p.Interpolate(1.5)
		assert.Equal(t, x, 2.)
		assert.Equal(t, y, 1.)

		x, y = p.Interpolate(1.75)
		assert.Equal(t, x, 2.5)
		assert.Equal(t, y, 1.)
	})

	t.Run("cubic bezier interpolation", func(t *testing.T) {
		segments := []Segment{
			NewCubicBezierSegment(point.NewPoint(0, 0), point.NewPoint(25, 50), point.NewPoint(75, 50)),
			NewSegment(100, 100),
		}

		p := FromSegments(segments, false)

		x, y := p.Interpolate(0)
		assert.Equal(t, x, 0.)
		assert.Equal(t, y, 0.)

		x, y = p.Interpolate(0.25)
		assert.Equal(t, x, 22.65625)
		assert.Equal(t, y, 29.6875)

		x, y = p.Interpolate(0.5)
		assert.Equal(t, x, 50.)
		assert.Equal(t, y, 50.)

		x, y = p.Interpolate(0.75)
		assert.Equal(t, x, 77.34375)
		assert.Equal(t, y, 70.3125)

		x, y = p.Interpolate(1)
		assert.Equal(t, x, 100.)
		assert.Equal(t, y, 100.)
	})
}
