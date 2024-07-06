package geometry

import "math"

// NormalizeRadians normalizes a radian value to be between -π and π
func NormalizeRadians(radians float64) float64 {
	for radians < -math.Pi {
		radians += 2 * math.Pi
	}
	for radians > math.Pi {
		radians -= 2 * math.Pi
	}
	return radians
}

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func RadiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
