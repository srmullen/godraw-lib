package vector

import (
	"math"
	"testing"

	"github.com/srmullen/godraw-lib/geometry/d2"
	"github.com/srmullen/godraw-lib/geometry/d2/point"
	"github.com/stretchr/testify/assert"
)

func TestAngle(t *testing.T) {
	v1 := point.NewPoint(1, 1)
	v2 := point.NewPoint(1, 0)
	// assert.Equal(t, math.Pi/4, Angle(v1, v2))
	assert.True(t, d2.WithinTolerance(math.Pi/4, Angle(v1, v2), 0.0000001))

	v1 = point.NewPoint(1, 0)
	v2 = point.NewPoint(-1, 0)
	assert.Equal(t, math.Pi, Angle(v1, v2))

	v1 = point.NewPoint(-1, 0)
	v2 = point.NewPoint(0, -1)
	assert.True(t, d2.WithinTolerance(math.Pi/2, Angle(v1, v2), 0.0000001))

	v1 = point.NewPoint(0, -1)
	v2 = point.NewPoint(0, 1)
	assert.Equal(t, math.Pi, Angle(v1, v2))
}

func TestDot(t *testing.T) {
	v1 := point.NewPoint(7, 2)
	v2 := point.NewPoint(3, 6)
	assert.Equal(t, 33., Dot(v1, v2))

	v1 = point.NewPoint(1, 1)
	v2 = point.NewPoint(1, 0)
	assert.True(t, d2.WithinTolerance(1.0, Dot(v1, v2), 0.0000001))

	v1 = point.NewPoint(1, 0)
	v2 = point.NewPoint(-1, 0)
	assert.True(t, d2.WithinTolerance(-1.0, Dot(v1, v2), 0.0000001))

	v1 = point.NewPoint(-1, 0)
	v2 = point.NewPoint(0, -1)
	assert.True(t, d2.WithinTolerance(0.0, Dot(v1, v2), 0.0000001))

	v1 = point.NewPoint(0, -1)
	v2 = point.NewPoint(0, 1)
	assert.True(t, d2.WithinTolerance(-1.0, Dot(v1, v2), 0.0000001))
}
