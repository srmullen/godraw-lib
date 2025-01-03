package bezier

import (
	"log"
	"math"

	"github.com/srmullen/godraw-lib/util"

	"github.com/srmullen/godraw-lib/geometry/d2/bounds"
	"github.com/srmullen/godraw-lib/geometry/d2/line"
	"github.com/srmullen/godraw-lib/geometry/d2/point"
)

// TODO: Implement reqcursively up to a depth

// p1, p2, p3, p4 are the control points
// t is the parameter range 0. to 1.
func Decasteljau(p1, p2, p3, p4 point.Point, t float64) (float64, float64) {
	x1 := util.Lerp(p1.X, p2.X, t)
	y1 := util.Lerp(p1.Y, p2.Y, t)

	x2 := util.Lerp(p2.X, p3.X, t)
	y2 := util.Lerp(p2.Y, p3.Y, t)

	x3 := util.Lerp(p3.X, p4.X, t)
	y3 := util.Lerp(p3.Y, p4.Y, t)

	x4 := util.Lerp(x1, x2, t)
	y4 := util.Lerp(y1, y2, t)

	x5 := util.Lerp(x2, x3, t)
	y5 := util.Lerp(y2, y3, t)

	x := util.Lerp(x4, x5, t)
	y := util.Lerp(y4, y5, t)

	return x, y
}

func Bernstein(p1, p2, p3, p4 point.Point, t float64) (float64, float64) {
	x := p1.X*(1-t)*(1-t)*(1-t) + 3*p2.X*t*(1-t)*(1-t) + 3*p3.X*t*t*(1-t) + p4.X*t*t*t
	y := p1.Y*(1-t)*(1-t)*(1-t) + 3*p2.Y*t*(1-t)*(1-t) + 3*p3.Y*t*t*(1-t) + p4.Y*t*t*t
	return x, y
}

func Polynomial(p1, p2, p3, p4 point.Point, t float64) (float64, float64) {
	x := p1.X + t*(-3*p1.X+3*p2.X) + t*t*(3*p1.X+-6*p2.X+3*p3.X) + t*t*t*(-p1.X+3*p2.X+-3*p3.X+p4.X)
	y := p1.Y + t*(-3*p1.Y+3*p2.Y) + t*t*(3*p1.Y+-6*p2.Y+3*p3.Y) + t*t*t*(-p1.Y+3*p2.Y+-3*p3.Y+p4.Y)
	return x, y
}

func DecasteljauSteps(p1, p2, p3, p4 point.Point, steps int) []point.Point {
	points := make([]point.Point, 0)
	for _, n := range util.Linspace(0, 1, steps) {
		x, y := Decasteljau(p1, p2, p3, p4, n)
		points = append(points, point.NewPoint(x, y))
	}
	return points
}

func BernsteinSteps(p1, p2, p3, p4 point.Point, steps int) []point.Point {
	points := make([]point.Point, 0)
	for _, n := range util.Linspace(0, 1, steps) {
		x, y := Bernstein(p1, p2, p3, p4, n)
		points = append(points, point.NewPoint(x, y))
	}
	return points
}

// Fastest because coefficients can be precomputed/cached
func PolynomialSteps(p1, p2, p3, p4 point.Point, steps int) []point.Point {
	points := make([]point.Point, 0)
	for _, n := range util.Linspace(0, 1, steps) {
		x, y := Polynomial(p1, p2, p3, p4, n)
		points = append(points, point.NewPoint(x, y))
	}
	return points
}

func Derivative(p1, p2, p3, p4 point.Point, t float64) (float64, float64) {
	mt := 1 - t
	x := 3 * (mt*mt*(p2.X-p1.X) + 2*mt*t*(p3.X-p2.X) + t*t*(p4.X-p3.X))
	y := 3 * (mt*mt*(p2.Y-p1.Y) + 2*mt*t*(p3.Y-p2.Y) + t*t*(p4.Y-p3.Y))
	return x, y
}

func DerivativeCubicBezier(p0, p1, p2, p3 point.Point, t float64) (float64, float64) {
	// Coefficients for cubic Bézier derivative
	c0 := -3 * math.Pow(1-t, 2)
	c1 := 3 * (1 - 4*t + 3*math.Pow(t, 2))
	c2 := 3 * (2*t - 3*math.Pow(t, 2))
	c3 := 3 * math.Pow(t, 2)

	x := c0*p0.X + c1*p1.X + c2*p2.X + c3*p3.X
	y := c0*p0.Y + c1*p1.Y + c2*p2.Y + c3*p3.Y
	return x, y
}

// This can only find one intersection per subdivision.
// To fix, if subdivision contains an intersection, could subdivide again to check for more points.
// Returns the points in order of their t value.
func LineIntersectionsNewtonsMethod(c1, c2, c3, c4, lstart, lend point.Point) []point.InterpolationPoint {
	maxIterations := 10
	epsilon := 1e-2

	var intersections []point.InterpolationPoint
	subdivisions := 100
	step := 1.0 / float64(subdivisions)

	for i := 0; i < subdivisions; i++ {
		tmin := float64(i) * step
		tmax := tmin + step
		ci := findLineIntersection(c1, c2, c3, c4, lstart, lend, tmin, tmax, epsilon, maxIterations)
		if ci != nil {
			// Check if this intersection is unique
			unique := true
			for _, existing := range intersections {
				if math.Abs(existing.X-ci.X) < epsilon && math.Abs(existing.Y-ci.Y) < epsilon {
					unique = false
					break
				}
			}
			if unique {
				intersections = append(intersections, *ci)
			}
		}
	}

	return intersections
}

// Recursively finds intersections of a bezier curve by dividing it into smaller lines.
// The smaller the distance between t values the more accurate the point will be.
func findLineIntersection(c1, c2, c3, c4, lstart, lend point.Point, tmin, tmax, epsilon float64, maxIterations int) *point.InterpolationPoint {
	x1, y1 := Polynomial(c1, c2, c3, c4, tmin)
	x2, y2 := Polynomial(c1, c2, c3, c4, tmax)
	x, y, ok := line.GetIntersection(x1, y1, x2, y2, lstart.X, lstart.Y, lend.X, lend.Y)
	if !ok {
		// return 0, 0, false
		return nil
	}
	tmid := (tmin + tmax) / 2.0
	if ci := findLineIntersection(c1, c2, c3, c4, lstart, lend, tmin, tmid, epsilon, maxIterations-1); ci != nil {
		return ci
	}
	if ci := findLineIntersection(c1, c2, c3, c4, lstart, lend, tmin, tmid, epsilon, maxIterations-1); ci != nil {
		return ci
	}
	return &point.InterpolationPoint{
		T:     tmid, // How to choose the interpolation point?
		Point: point.NewPoint(x, y),
	}
}

// // Recursively finds intersections of a bezier curve by dividing it into smaller lines.
// // The smaller the distance between t values the more accurate the point will be.
// func findLineIntersection(c1, c2, c3, c4, lstart, lend point.Point, tmin, tmax, epsilon float64, maxIterations int) (float64, float64, bool) {
// 	x1, y1 := Polynomial(c1, c2, c3, c4, tmin)
// 	x2, y2 := Polynomial(c1, c2, c3, c4, tmax)
// 	x, y, ok := line.GetIntersection(x1, y1, x2, y2, lstart.X, lstart.Y, lend.X, lend.Y)
// 	if !ok {
// 		return 0, 0, false
// 	}
// 	tmid := (tmin + tmax) / 2.0
// 	if x, y, ok := findLineIntersection(c1, c2, c3, c4, lstart, lend, tmin, tmid, epsilon, maxIterations-1); ok {
// 		return x, y, ok
// 	}
// 	if x, y, ok := findLineIntersection(c1, c2, c3, c4, lstart, lend, tmin, tmid, epsilon, maxIterations-1); ok {
// 		return x, y, ok
// 	}
// 	return x, y, ok
// }

// // Finds line intersections on a bezier curve. This always finds the same point even if there are multiple and tmin/tmax dont
// // seem to affect the result. I haven't figured out why that is the case yet.
// func findLineIntersection(c1, c2, c3, c4, lstart, lend point.Point, tmin, tmax, epsilon float64, maxIterations int) (float64, float64, bool) {
// 	for i := 0; i < maxIterations; i++ {
// 		t := (tmin + tmax) / 2.
// 		x, y := Polynomial(c1, c2, c3, c4, t)
// 		// Check if point is on line using cross product
// 		lineDir := lend.SubtractPoint(lstart)
// 		vx := x - lstart.X
// 		vy := y - lstart.Y
// 		cross := vx*lineDir.X - vy*lineDir.Y
// 		if math.Abs(cross) < epsilon {
// 			return x, y, true
// 		}
// 		if cross > 0. {
// 			tmax = t
// 		} else {
// 			tmin = t
// 		}
// 	}
// 	return 0, 0, false
// }

// // Uses newtons method to find line intersections on a bezier curve. This fails when the line is vertical.
// func findLineIntersection(c1, c2, c3, c4, lstart, lend point.Point, tStart, tend, epsilon float64, maxIterations int) (float64, float64, bool) {
// 	t := tStart
// 	lne := line.NewLine(lstart.X, lstart.Y, lend.X, lend.Y)
// 	for i := 0; i < maxIterations; i++ {
// 		x, y := Polynomial(c1, c2, c3, c4, t)

// 		if math.Abs(y-lne.PointSlopeForm(x)) < epsilon {
// 			// log.Println(x, y, lne.PointSlopeForm(x))
// 			return x, y, true
// 		}

// 		dx, dy := DerivativeCubicBezier(c1, c2, c3, c4, t)
// 		if dy-lne.PointSlopeForm(dx) == 0. {
// 			log.Println(dx, dy, lne.PointSlopeForm(dx))
// 			break // Avoid division by zero
// 		}

// 		t -= (y - lne.PointSlopeForm(x)) / (dy - lne.M()*dx)
// 		if math.IsNaN(t) {
// 			break
// 		}

// 		if t < 0 {
// 			t = 0.
// 		} else if t > 1 {
// 			t = 1.
// 		}
// 	}
// 	return 0, 0, false
// }

func LineIntersections(c1, c2, c3, c4, lstart, lend point.Point) ([]float64, bool) {
	t := 0.5
	x, y := Polynomial(c1, c2, c3, c4, t)
	intersections := make([]float64, 0)
	x1, y1, i1 := line.GetIntersection(c1.X, c1.Y, x, y, lstart.X, lstart.Y, lend.X, lend.Y)
	if i1 {
		intersections = append(intersections, x1, y1)
	}

	x2, y2, i2 := line.GetIntersection(c4.X, c4.Y, x, y, lstart.X, lstart.Y, lend.X, lend.Y)
	if i2 {
		intersections = append(intersections, x2, y2)
	}

	return intersections, len(intersections) > 0
}

func BezierIntersections(b1, b2 []point.Point, depth int) []point.Point {
	log.Println(depth)
	if depth == 0 {
		return []point.Point{}
	}
	epsilon := 1e-6
	bnd1 := Bounds(b1[0], b1[1], b1[2], b1[3])
	bnd2 := Bounds(b2[0], b2[1], b2[2], b2[3])
	if !bnd1.Overlaps(bnd2) {
		return []point.Point{}
	}

	if isSmallEnough(bnd1, epsilon) && isSmallEnough(bnd2, epsilon) {
		midPoint := point.NewPoint(
			(bnd1.Left+bnd1.Right+bnd2.Left+bnd2.Right)/4.,
			(bnd1.Top+bnd1.Bottom+bnd2.Top+bnd2.Bottom)/4,
		)
		return []point.Point{midPoint}
	}

	b1Left, b1Right := Subdivide(b1, 0.5)
	b2Left, b2Right := Subdivide(b2, 0.5)

	intersections := BezierIntersections(b1Left, b2Left, depth-1)
	intersections = append(intersections, BezierIntersections(b1Right, b2Right, depth-1)...)
	intersections = append(intersections, BezierIntersections(b1Left, b2Right, depth-1)...)
	intersections = append(intersections, BezierIntersections(b1Right, b2Left, depth-1)...)

	return removeDuplicates(intersections, epsilon)
}

func removeDuplicates(points []point.Point, epsilon float64) []point.Point {
	var result []point.Point
	for _, p := range points {
		if !contains(result, p, epsilon) {
			result = append(result, p)
		}
	}
	return result
}

func contains(points []point.Point, p point.Point, epsilon float64) bool {
	for _, existing := range points {
		if math.Abs(existing.X-p.X) < epsilon && math.Abs(existing.Y-p.Y) < epsilon {
			return true
		}
	}
	return false
}

func Subdivide(b []point.Point, t float64) ([]point.Point, []point.Point) {
	p01 := interpolate(b[0], b[1], t)
	p12 := interpolate(b[1], b[2], t)
	p23 := interpolate(b[2], b[3], t)
	p012 := interpolate(p01, p12, t)
	p123 := interpolate(p12, p23, t)
	p0123 := interpolate(p012, p123, t)

	return []point.Point{b[0], p01, p012, p0123},
		[]point.Point{p0123, p123, p23, b[3]}
}

func interpolate(p1, p2 point.Point, t float64) point.Point {
	return point.NewPoint(
		p1.X*(1-t)+p2.X*t,
		p1.Y*(1-t)+p2.Y*t,
	)
}

func isSmallEnough(b bounds.Bounds, epsilon float64) bool {
	// return math.Abs(max.X-min.X) < epsilon && math.Abs(max.Y-min.Y) < epsilon
	return math.Abs(b.Right-b.Left) < epsilon && math.Abs(b.Bottom-b.Top) < epsilon
}

func Bounds(c1, c2, c3, c4 point.Point) bounds.Bounds {
	left := math.Min(math.Min(math.Min(c1.X, c2.X), c3.X), c4.X)
	top := math.Min(math.Min(math.Min(c1.Y, c2.Y), c3.Y), c4.Y)
	right := math.Max(math.Max(math.Max(c1.X, c2.X), c3.X), c4.X)
	bottom := math.Max(math.Max(math.Max(c1.Y, c2.Y), c3.Y), c4.Y)
	return bounds.NewBounds(top, right, bottom, left)
}

// Length approximates the length of a cubic Bézier curve using subdivision
func Length(c1, c2, c3, c4 point.Point) float64 {
	epsilon := 1e-6
	chord := c1.Distance(c4)
	control := c1.Distance(c2) + c2.Distance(c3) + c3.Distance(c4)
	if control-chord <= epsilon {
		return (chord + control) / 2
	}
	left, right := Subdivide([]point.Point{c1, c2, c3, c4}, 0.5)
	return Length(left[0], left[1], left[2], left[3]) + Length(right[0], right[1], right[2], right[3])
}
