package geometry

import (
	"math"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNormalizeRadians(t *testing.T) {
	assert.Equal(t, 0.0, NormalizeRadians(0))
	assert.Equal(t, math.Pi, NormalizeRadians(math.Pi))
	assert.Equal(t, -math.Pi, NormalizeRadians(-math.Pi))
	assert.Equal(t, 0.0, NormalizeRadians(2*math.Pi))
	assert.Equal(t, 0.0, NormalizeRadians(-2*math.Pi))
	assert.Equal(t, math.Pi/2, NormalizeRadians(5*math.Pi/2))
	assert.Equal(t, -math.Pi/2, NormalizeRadians(-5*math.Pi/2))
}

func TestDegreesToRadians(t *testing.T) {
	assert.Equal(t, 0.0, DegreesToRadians(0))
	assert.Equal(t, math.Pi, DegreesToRadians(180))
	assert.Equal(t, math.Pi/2, DegreesToRadians(90))
	assert.Equal(t, math.Pi*2, DegreesToRadians(360))
}

func TestRadiansToDegrees(t *testing.T) {
	assert.Equal(t, 0.0, RadiansToDegrees(0))
	assert.Equal(t, 180.0, RadiansToDegrees(math.Pi))
	assert.Equal(t, 90.0, RadiansToDegrees(math.Pi/2))
	assert.Equal(t, 360.0, RadiansToDegrees(math.Pi*2))
}
