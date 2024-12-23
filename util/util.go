package util

import (
	"fmt"
)

func RGBToHex(r, g, b int) string {
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func UNUSED(t ...any) {}

func MapRange(value, low1, high1, low2, high2 float64) float64 {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

func Lerp(a, b, t float64) float64 {
	return a + t*(b-a)
}

// Return evenly spaced numbers over a specified interval.
func Linspace(start, end float64, n int) []float64 {
	ret := make([]float64, n)
	for i := 0; i < n; i += 1 {
		ret[i] = start + float64(i)*(end-start)/(float64(n)-1)
	}
	return ret
}

func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func Mod(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}

	// Get the absolute value of b
	b = abs(b)

	// First get remainder using built-in % operator
	result := a % b

	// If remainder is negative, add the modulus
	if result < 0 {
		result += b
	}

	return result
}

// abs returns the absolute value of x
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
