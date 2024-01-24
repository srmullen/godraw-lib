package d2

import "math"

func WithinTolerance(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

func RoundToDecimal(x, decimals float64) float64 {
	return math.Round(x*math.Pow(10, decimals)) / math.Pow(10, decimals)
}
