package path

import (
	"testing"

	"github.com/srmullen/godraw-lib/geometry/d2/point"
	"github.com/stretchr/testify/assert"
)

func TestGetBounds(t *testing.T) {
	p := NewClosedPath([]float64{})
	bounds := p.GetBounds()
	assert.Nil(t, bounds)
}

func TestLength(t *testing.T) {
	t.Run("path with no segements", func(t *testing.T) {
		p := NewClosedPath([]float64{})
		length := p.Length()
		assert.Equal(t, length, 0.)
	})

	t.Run("path made up of only lines", func(t *testing.T) {
		p := NewClosedPath([]float64{0, 0, 1, 1, 2, 2})
		length := p.Length()
		assert.Equal(t, length, 2.8284271247461903*2.)

		p = NewClosedPath([]float64{0, 0, 1, 0, 1, 1, 0, 1})
		length = p.Length()
		assert.Equal(t, length, 4.)

		p = NewOpenPath([]float64{0, 0, 1, 1, 2, 2})
		length = p.Length()
		assert.Equal(t, length, 2.8284271247461903)

		p = NewOpenPath([]float64{0, 0, 2, 0, 2, 2})
		length = p.Length()
		assert.Equal(t, length, 4.)
	})

	t.Run("path made up of lines and curves", func(t *testing.T) {
		segments := []Segment{
			{
				Point: point.Point{0, 0},
				Curve: NewCubicBezier(point.Point{25, 50}, point.Point{75, 50}),
			},
			{
				Point: point.Point{100, 100},
				Curve: nil,
			},
		}
		p := FromSegments(segments, false)
		length := p.Length()
		assert.Equal(t, 143.32764325352355, length)

		segments = []Segment{
			{
				Point: point.Point{0, 0},
				Curve: NewCubicBezier(point.Point{25, 50}, point.Point{75, 50}),
			},
			{
				Point: point.Point{100, 100},
				Curve: nil,
			},
		}
		p = FromSegments(segments, true)
		length = p.Length()
		assert.Equal(t, 284.74899949083306, length)
	})
}
