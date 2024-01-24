package d2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithinTolerance(t *testing.T) {
	assert.True(t, WithinTolerance(1.0, 1.0, 0.0001))
	assert.True(t, WithinTolerance(1.0, 1.0001, 0.0001))
	assert.False(t, WithinTolerance(1.0, 1.0001, 0.00001))
	assert.False(t, WithinTolerance(1.0, 10, 2))
}

func TestRoundToDecimal(t *testing.T) {
	assert.Equal(t, 1.2, RoundToDecimal(1.23456789, 1))
	assert.Equal(t, 1.23, RoundToDecimal(1.23456789, 2))
	assert.Equal(t, 1.235, RoundToDecimal(1.23456789, 3))
	assert.Equal(t, 1.2346, RoundToDecimal(1.23456789, 4))
	assert.Equal(t, 1.23457, RoundToDecimal(1.23456789, 5))
	assert.Equal(t, 1.234568, RoundToDecimal(1.23456789, 6))
	assert.Equal(t, 1.2345679, RoundToDecimal(1.23456789, 7))
	assert.Equal(t, 1.23456789, RoundToDecimal(1.23456789, 8))
}
