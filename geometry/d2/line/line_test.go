package line

import (
	"math"
	"testing"

	"github.com/srmullen/godraw-lib/geometry"
	"github.com/srmullen/godraw-lib/geometry/d2"
	"github.com/stretchr/testify/assert"
)

func TestMagnitude(t *testing.T) {
	assert.Equal(t, 1., NewLine(0, 0, 0, 1).Magnitude())
	assert.True(t, d2.WithinTolerance(1.41421, NewLine(0, 0, 1, 1).Magnitude(), 0.00001))
	assert.True(t, d2.WithinTolerance(1.41421, NewLine(30, 30, 31, 31).Magnitude(), 0.00001))
	assert.True(t, d2.WithinTolerance(1.41421, NewLine(0, 0, -1, -1).Magnitude(), 0.00001))
}

func TestDirection(t *testing.T) {
	assert.Equal(t, 0., NewLine(0, 0, 1, 0).Direction())
	assert.Equal(t, 0., NewLine(10, 10, 11, 10).Direction())

	assert.Equal(t, math.Pi/2, NewLine(0, 0, 0, 1).Direction())
	assert.Equal(t, math.Pi/4, NewLine(0, 0, 1, 1).Direction())
	assert.Equal(t, math.Pi*3/4, NewLine(0, 0, -1, 1).Direction())
	assert.Equal(t, math.Pi, NewLine(0, 0, -1, 0).Direction())
	assert.Equal(t, geometry.NormalizeRadians(math.Pi+math.Pi/4), NewLine(0, 0, -1, -1).Direction())
}

func TestLineIntersection(t *testing.T) {

	t.Run("diagonal lines", func(t *testing.T) {
		x, y, ok := GetIntersection(0, 0, 1, 1, 0, 1, 1, 0)
		assert.True(t, ok)
		assert.Equal(t, 0.5, x)
		assert.Equal(t, 0.5, y)
	})

	t.Run("perpendicular lines", func(t *testing.T) {
		x, y, ok := GetIntersection(100, 100, 200, 100, 150, 100, 150, 200)
		assert.True(t, ok)
		assert.Equal(t, 150., x)
		assert.Equal(t, 100., y)
	})

	t.Run("should intersect", func(t *testing.T) {
		x1 := 426.6338998124982
		y1 := 321.3492466862989
		x2 := x1 + 200
		y2 := y1

		x3 := 448.45084971874735
		y3 := 293.38926261462365
		x4 := 438.90169943749476
		y4 := 359.10565162951536
		// draw.Circle(pnt.X, pnt.Y, 5)
		// starLine := path.NewLine(448.45084971874735, 293.38926261462365, 438.90169943749476, 359.10565162951536)
		// lne := path.NewLine(pnt.X, pnt.Y, pnt.X+200, pnt.Y)

		_, _, ok := GetIntersection(x1, y1, x2, y2, x3, y3, x4, y4)
		assert.True(t, ok)
	})
}
