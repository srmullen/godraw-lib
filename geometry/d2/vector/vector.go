package vector

import (
	"math"

	"github.com/srmullen/godraw-lib/geometry"
)

// How to return struct with generic type?
// func Add[T geometry.Vector](v1, v2 T) (float64, float64) {
// 	x1, y1 := v1.Endpoint()
// 	x2, y2 := v2.Endpoint()
// 	return x1 + x2, y1 + y2
// }

// func Subtract[T geometry.Vector](v1, v2 T) (float64, float64) {
// 	x1, y1 := v1.Endpoint()
// 	x2, y2 := v2.Endpoint()
// 	return x1 - x2, y1 - y2
// }

func Dot[T geometry.Vector](v1, v2 T) float64 {
	m1 := v1.Magnitude()
	m2 := v2.Magnitude()
	d1 := v1.Direction()
	d2 := v2.Direction()
	return m1 * m2 * math.Cos(d1-d2)
}

// Returns the angle between two vectors in radians
func Angle[T geometry.Vector](v1, v2 T) float64 {
	return math.Acos(Dot(v1, v2) / (v1.Magnitude() * v2.Magnitude()))
}
