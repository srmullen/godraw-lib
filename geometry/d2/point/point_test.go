package point

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAngle(t *testing.T) {
	p := NewPoint(1, 1)
	assert.Equal(t, math.Pi/4, p.Angle())

	p = NewPoint(1, 0)
	assert.Equal(t, 0.0, p.Angle())

	p = NewPoint(-1, 0)
	assert.Equal(t, math.Pi, p.Angle())

	p = NewPoint(0, -1)
	assert.Equal(t, -math.Pi/2, p.Angle())

	p = NewPoint(0, 1)
	assert.Equal(t, math.Pi/2, p.Angle())
}
