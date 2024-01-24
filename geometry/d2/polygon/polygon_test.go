package polygon

import (
	"testing"

	"github.com/srmullen/godraw-lib/geometry/d2/point"

	"github.com/stretchr/testify/assert"
)

func TestGetIntersections(t *testing.T) {

	t.Run("Bounds do not intersect", func(t *testing.T) {
		p1 := NewPolygon([]*point.Point{
			{
				X: 0,
				Y: 0,
			},
			{
				X: 1,
				Y: 1,
			},
			{
				X: 1,
				Y: 0,
			},
		})

		p2 := NewPolygon([]*point.Point{
			{
				X: 0,
				Y: 1,
			},
			{
				X: 1,
				Y: 1,
			},
			{
				X: 1,
				Y: 2,
			},
		})

		intersections := p1.GetIntersections(p2)
		assert.Equal(t, 0, len(intersections))
	})
	t.Run("Intersects", func(t *testing.T) {
		p1 := NewPolygon([]*point.Point{
			{
				X: 0,
				Y: 0,
			},
			{
				X: 2,
				Y: 2,
			},
			{
				X: 2,
				Y: 0,
			},
		})

		p2 := NewPolygon([]*point.Point{
			{
				X: 0.5,
				Y: 1,
			},
			{
				X: 3,
				Y: 3,
			},
			{
				X: 1,
				Y: 0.5,
			},
		})

		intersections := p1.GetIntersections(p2)
		assert.Equal(t, 2, len(intersections))
		assert.Equal(t, 0.75, intersections[0].X)
		assert.Equal(t, 0.75, intersections[0].Y)
	})

	t.Run("Contains Point", func(t *testing.T) {
		// Rectangle
		// rect := NewRectangle(0, 0, 100, 100)
		// assert.True(t, rect.ContainsPoint(50, 50))
		// assert.False(t, rect.ContainsPoint(150, 50))
		// // Test point on edge
		// assert.True(t, rect.ContainsPoint(100, 50))

		// Star
		star := NewStar(100, 100, 100, 50, 5)
		// log.Println(star.Points())
		assert.True(t, star.ContainsPoint(100, 100))
		assert.True(t, star.ContainsPoint(100, 150))
		assert.False(t, star.ContainsPoint(100, 180)) // within bounds, but not in polygon
		assert.True(t, star.ContainsPoint(50, 100))   // on border
		assert.False(t, star.ContainsPoint(49, 100))  // on border
		assert.True(t, star.ContainsPoint(51, 100))   // on border
	})
}
