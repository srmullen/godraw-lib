package geometry

type Area interface {
	Width() float64
	Height() float64
	Area() float64
}

type Vector interface {
	Magnitude() float64
	Direction() float64

	// Normalize() Vector
	// ScalarMult(float64) Vector
	// Add(Vector) Vector
	// Subtract(Vector) Vector
	// Dot(Vector) float64
	// Cross(Vector) Vector
	// Hadamard(Vector) Vector
}

type Path interface {
	Translate(Vector) Path
	Scale(float64) Path
	Rotate(float64) Path
}
