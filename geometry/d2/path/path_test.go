package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBounds(t *testing.T) {
	p := NewClosedPath([]float64{})
	bounds := p.GetBounds()
	assert.Nil(t, bounds)
}
