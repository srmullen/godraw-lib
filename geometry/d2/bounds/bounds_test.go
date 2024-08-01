package bounds

import (
	"testing"

	"github.com/srmullen/godraw-lib/geometry/d2/point"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	t.Run("Contains", func(t *testing.T) {
		b := &Bounds{
			Top:    0,
			Right:  1,
			Bottom: 1,
			Left:   0,
		}

		p := &point.Point{
			X: 0.5,
			Y: 0.5,
		}
		assert.True(t, b.Contains(p.X, p.Y))
	})

	t.Run("Does not contain", func(t *testing.T) {
		b := &Bounds{
			Top:    0,
			Right:  1,
			Bottom: 1,
			Left:   0,
		}

		p := &point.Point{
			X: 2,
			Y: 2,
		}
		assert.False(t, b.Contains(p.X, p.Y))
	})

	t.Run("Point on edge", func(t *testing.T) {
		b := &Bounds{
			Top:    0,
			Right:  1,
			Bottom: 1,
			Left:   0,
		}

		p := &point.Point{
			X: 1,
			Y: 1,
		}
		assert.True(t, b.Contains(p.X, p.Y))
	})
}

func TestContainsBounds(t *testing.T) {
	t.Run("Contains", func(t *testing.T) {
		b1 := &Bounds{
			Top:    0,
			Right:  1,
			Bottom: 1,
			Left:   0,
		}

		b2 := &Bounds{
			Top:    0.1,
			Right:  0.5,
			Bottom: 0.5,
			Left:   0.1,
		}
		assert.True(t, b1.ContainsBounds(b2))
	})

	t.Run("Does not contain", func(t *testing.T) {
		b1 := &Bounds{
			Top:    0,
			Right:  1,
			Bottom: 1,
			Left:   0,
		}

		b2 := &Bounds{
			Top:    2.5,
			Right:  3.5,
			Bottom: 3.5,
			Left:   2.5,
		}
		assert.False(t, b1.ContainsBounds(b2))
	})

	t.Run("Bounds are the same", func(t *testing.T) {
		b1 := &Bounds{
			Top:    0,
			Right:  1,
			Bottom: 1,
			Left:   0,
		}

		b2 := &Bounds{
			Top:    0,
			Right:  1,
			Bottom: 1,
			Left:   0,
		}
		assert.True(t, b1.ContainsBounds(b2))
	})

	t.Run("Test inner does not contain outer", func(t *testing.T) {
		inner := &Bounds{
			Top:    0.1,
			Right:  0.5,
			Bottom: 0.5,
			Left:   0.1,
		}

		outer := &Bounds{
			Top:    0,
			Right:  1,
			Bottom: 1,
			Left:   0,
		}
		assert.False(t, inner.ContainsBounds(outer))
	})
}

func TestIntersect(t *testing.T) {
	t.Run("Intersects", func(t *testing.T) {
		b1 := &Bounds{
			Top:    0,
			Right:  1,
			Bottom: 1,
			Left:   0,
		}

		b2 := &Bounds{
			Top:    0.5,
			Right:  1.5,
			Bottom: 1.5,
			Left:   0.5,
		}
		assert.True(t, b1.Intersects(b2))
		assert.True(t, b2.Intersects(b1))
	})

	t.Run("Does not intersect", func(t *testing.T) {
		b1 := &Bounds{
			Top:    0,
			Right:  1,
			Bottom: 1,
			Left:   0,
		}

		b2 := &Bounds{
			Top:    2.5,
			Right:  3.5,
			Bottom: 3.5,
			Left:   2.5,
		}
		assert.False(t, b1.Intersects(b2))
		assert.False(t, b2.Intersects(b1))
	})
}

func TestOverlaps(t *testing.T) {
	t.Run("Overlaps", func(t *testing.T) {
		b1 := Bounds{
			Top:    0,
			Right:  1,
			Bottom: 1,
			Left:   0,
		}

		b2 := Bounds{
			Top:    0.5,
			Right:  1.5,
			Bottom: 1.5,
			Left:   0.5,
		}
		assert.True(t, b1.Overlaps(b2))
		assert.True(t, b2.Overlaps(b1))
	})

	t.Run("Does not overlap", func(t *testing.T) {
		b1 := Bounds{
			Top:    0,
			Right:  1,
			Bottom: 1,
			Left:   0,
		}

		b2 := Bounds{
			Top:    2.5,
			Right:  3.5,
			Bottom: 3.5,
			Left:   2.5,
		}
		assert.False(t, b1.Overlaps(b2))
		assert.False(t, b2.Overlaps(b1))
	})

	t.Run("Overlaps on edge", func(t *testing.T) {
		b1 := Bounds{
			Top:    0,
			Right:  1,
			Bottom: 1,
			Left:   0,
		}

		b2 := Bounds{
			Top:    0,
			Right:  2,
			Bottom: 1,
			Left:   1,
		}
		assert.False(t, b1.Overlaps(b2))
		assert.False(t, b2.Overlaps(b1))
	})

	t.Run("Overlaps on TopLeft/BottomRight corner", func(t *testing.T) {
		b1 := Bounds{
			Top:    0,
			Right:  2,
			Bottom: 2,
			Left:   0,
		}

		b2 := Bounds{
			Top:    1,
			Right:  3,
			Bottom: 3,
			Left:   1,
		}
		assert.True(t, b1.Overlaps(b2))
		assert.True(t, b2.Overlaps(b1))
	})

	t.Run("Overlaps on TopRight/BottomLeft corner", func(t *testing.T) {
		b1 := Bounds{
			Top:    0,
			Right:  4,
			Bottom: 2,
			Left:   2,
		}

		b2 := Bounds{
			Top:    1,
			Right:  3,
			Bottom: 3,
			Left:   1,
		}
		assert.True(t, b1.Overlaps(b2))
		assert.True(t, b2.Overlaps(b1))
	})

	t.Run("Only bounds interiors overlap", func(t *testing.T) {
		b1 := Bounds{
			Top:    2,
			Left:   2,
			Right:  3,
			Bottom: 6,
		}

		b2 := Bounds{
			Top:    3,
			Left:   1,
			Bottom: 4,
			Right:  4,
		}
		assert.True(t, b1.Overlaps(b2))
		assert.True(t, b2.Overlaps(b1))
	})

	t.Run("Only bounds interiors overlap", func(t *testing.T) {
		b1 := Bounds{
			Top:    1,
			Left:   1,
			Right:  4,
			Bottom: 2,
		}

		b2 := Bounds{
			Top:    0,
			Left:   2,
			Bottom: 4,
			Right:  3,
		}
		assert.True(t, b1.Overlaps(b2))
		assert.True(t, b2.Overlaps(b1))
	})
}
