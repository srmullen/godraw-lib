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

func TestAdd(t *testing.T) {
	p := NewPoint(1, 1)
	p2 := p.Add(1, 1)
	assert.Equal(t, NewPoint(2, 2), p2)
	assert.Equal(t, p, NewPoint(1, 1))
}

func TestSubtract(t *testing.T) {
	p := NewPoint(1, 1)
	p2 := p.Subtract(1, 1)
	assert.Equal(t, NewPoint(0, 0), p2)
	assert.Equal(t, p, NewPoint(1, 1))
}

func TestDivide(t *testing.T) {
	p := NewPoint(1, 1)
	p2 := p.Divide(2, 2)
	assert.Equal(t, NewPoint(0.5, 0.5), p2)
	assert.Equal(t, p, NewPoint(1, 1))
}

func TestMultiply(t *testing.T) {
	p := NewPoint(1, 1)
	p2 := p.Multiply(2, 2)
	assert.Equal(t, NewPoint(2, 2), p2)
	assert.Equal(t, p, NewPoint(1, 1))
}

func TestAddCoords(t *testing.T) {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(3, 4)
	p := p1.AddCoords(p2)
	assert.Equal(t, NewPoint(4, 6), p)
}

func TestSubtractCoords(t *testing.T) {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(3, 4)
	p := p1.SubtractCoords(p2)
	assert.Equal(t, NewPoint(-2, -2), p)
}

func TestDivideCoords(t *testing.T) {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(2, 2)
	p := p1.DivideCoords(p2)
	assert.Equal(t, NewPoint(0.5, 1), p)
}

func TestMultiplyCoords(t *testing.T) {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(2, 2)
	p := p1.MultiplyCoords(p2)
	assert.Equal(t, NewPoint(2, 4), p)
}
