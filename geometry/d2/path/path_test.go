package path

import (
	"testing"

	"github.com/srmullen/godraw-lib/geometry/d2/point"

	"github.com/stretchr/testify/assert"
)

func TestGetBounds(t *testing.T) {
	p := NewClosedPath([]*point.Point{})
	bounds := p.GetBounds()
	assert.Nil(t, bounds)
}
