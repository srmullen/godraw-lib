package polygon

import (
	"math"

	"github.com/srmullen/godraw-lib/geometry/d2"
	"github.com/srmullen/godraw-lib/geometry/d2/line"
	"github.com/srmullen/godraw-lib/geometry/d2/path"
	"github.com/srmullen/godraw-lib/geometry/d2/point"

	"github.com/engelsjk/polygol"
)

type Points []point.Point

func (p Points) Data() []float64 {
	var data []float64
	for _, point := range p {
		data = append(data, point.X, point.Y)
	}
	return data
}

// Polygon is a closed path that does not have any curves
type Polygon struct {
	*path.Path
}

func NewPolygon(coords d2.CoordData) *Polygon {
	return &Polygon{
		Path: path.NewPath(coords.Data(), true),
	}
}

func NewStar(x, y, radius, innerRadius float64, points int) *Polygon {
	if points < 3 {
		panic("Cannot create polygon with less than 3 sides")
	}
	points *= 2
	coords := Points{}
	for i := 0; i < points; i++ {
		if i%2 == 0 {
			coords = append(coords, point.NewPointFromAngle(float64(i)*2*math.Pi/float64(points), radius).Add(x, y))
		} else {
			coords = append(coords, point.NewPointFromAngle(float64(i)*2*math.Pi/float64(points), innerRadius).Add(x, y))
		}
	}
	return &Polygon{
		Path: path.NewPath(coords.Data(), true),
	}
}

func isLess(xs []float64, i, j int) bool {
	// return segment < next && segment < prev
	return ((xs[i%len(xs)] < xs[j%len(xs)]) || (xs[i%len(xs)] == xs[j%len(xs)] && i < j))
}

// l is some other line of the form ax + by = c
// FIXME: need to be able to pass in the line
func isLessGen(xs, ys []float64, i, j int) bool {
	// How to pass in the line varaibles?
	a := 1.0
	b := 1.0
	// c := 1.0

	di := a*xs[i] + b*ys[i]
	dj := a*xs[j%len(xs)] + b*ys[j%len(ys)]
	return ((di < dj) || (di == dj && i < j))
}

// FIXME: doesn't work. need to fix isLessGen
func (p *Polygon) IsMonotone() bool {
	xs := p.Xs()
	localMins := 0
	for i := 1; i < len(xs); i++ {
		if isLess(xs, i, i-1) && isLess(xs, i, i+1) {
			localMins += 1
		}
	}
	return localMins == 1
}

// https://cs.stackexchange.com/questions/1577/how-do-i-test-if-a-polygon-is-monotone-with-respect-to-a-line
// Monotone with resect to the x-axis
func (p *Polygon) IsXMonotone() bool {
	xs := p.Xs()
	localMins := 0
	for i := 1; i < len(xs); i++ {
		if isLess(xs, i, i-1) && isLess(xs, i, i+1) {
			localMins += 1
		}
	}
	return localMins == 1
}

// Monotone with resect to the y-axis
// Not working - isLess doesn't work for this
func (p *Polygon) IsYMonotone() bool {
	ys := p.Ys()
	localMins := 0
	for i := 1; i < len(ys); i++ {
		if isLess(ys, i, i-1) && isLess(ys, i, i+1) {
			localMins += 1
		}
	}
	return localMins == 1
}

func NewNgon(sides int, x, y, radius float64) *Polygon {
	if sides < 3 {
		panic("Cannot create polygon with less than 3 sides")
	}
	points := Points{}
	for i := 0; i < sides; i++ {
		points = append(points, point.NewPointFromAngle(float64(i)*2*math.Pi/float64(sides), radius).Add(x, y))
	}
	return &Polygon{
		Path: path.NewPath(points.Data(), true),
	}
}

func (p *Polygon) GetIntersections(other *Polygon) Points {
	intersections := Points{}
	// If bounds don't intersect, then there are no intersections
	if !p.GetBounds().Overlaps(other.GetBounds()) {
		return intersections
	}
	for i := 0; i < len(p.Segments); i++ {
		from := p.Segments[i].Point
		to := p.Segments[(i+1)%len(p.Segments)].Point
		for j := 0; j < len(other.Segments); j++ {
			jFrom := other.Segments[j].Point
			jTo := other.Segments[(j+1)%len(other.Segments)].Point
			// intersection := segment.Point.Intersection(other.Segments[j].Point)
			// intersection := line.NewLine(from.X, from.Y, to.X, to.Y).Intersection(line.NewLine(jFrom.X, jFrom.Y, jTo.X, jTo.Y))
			if x, y, ok := line.GetIntersection(from.X, from.Y, to.X, to.Y, jFrom.X, jFrom.Y, jTo.X, jTo.Y); ok {
				intersections = append(intersections, point.NewPoint(x, y))
			}
		}
	}
	return intersections
}

func (p *Polygon) LineIntersections(lne *path.Path) Points {
	// var intersections []*point.Point
	intersections := Points{}
	for i := 0; i < len(p.Segments); i++ {
		from := p.Segments[i].Point
		to := p.Segments[(i+1)%len(p.Segments)].Point
		if x, y, ok := line.GetIntersection(from.X, from.Y, to.X, to.Y, lne.Segments[0].X, lne.Segments[0].Y, lne.Segments[1].X, lne.Segments[1].Y); ok {
			intersections = append(intersections, point.NewPoint(x, y))
		}
		// intersection := line.NewLine(from.X, from.Y, to.X, to.Y).Intersection(lne)
		// if intersection != nil {
		// 	intersections = append(intersections, intersection)
		// }
	}
	// Remove duplicates (intersection through the corner of a polygon)
	// ret := []*point.Point{}
	ret := Points{}
	for _, intersection := range intersections {
		found := false
		for _, point := range ret {
			if point.EqualsWithTolerance(intersection, 0.0001) {
				found = true
				break
			}
		}
		if !found {
			ret = append(ret, intersection)
		}
	}
	return ret
}

func (p *Polygon) Translate(x, y float64) *Polygon {
	return &Polygon{
		Path: path.PathTranslate(p.Path, x, y),
	}
}

func (p *Polygon) Scale(scalex, scaley float64) *Polygon {
	return &Polygon{
		Path: p.Path.Scale(scalex, scaley),
	}
}

func (p *Polygon) ToGeom() polygol.Geom {
	var points [][]float64
	for _, segment := range p.Segments {
		points = append(points, []float64{segment.X, segment.Y})
	}
	return [][][][]float64{{points}}
}

func ToGeoms(polygons ...*Polygon) []polygol.Geom {
	geoms := make([]polygol.Geom, len(polygons))
	for i, polygon := range polygons {
		geoms[i] = polygon.ToGeom()
	}
	return geoms
}

func FromGeom(multipolygon polygol.Geom) []*Polygon {
	polygons := make([]*Polygon, len(multipolygon))
	for i, polygon := range multipolygon {
		var points Points
		for _, vertex := range polygon[0] {
			points = append(points, point.Point{X: vertex[0], Y: vertex[1]})
		}
		polygons[i] = NewPolygon(points)
	}
	return polygons
}

func (p *Polygon) Union(others ...*Polygon) []*Polygon {
	geoms := ToGeoms(others...)
	multipolygons, _ := polygol.Union(p.ToGeom(), geoms...)
	return FromGeom(multipolygons)
}

func (p *Polygon) Intersection(others ...*Polygon) []*Polygon {
	geoms := ToGeoms(others...)
	multipolygons, _ := polygol.Intersection(p.ToGeom(), geoms...)
	return FromGeom(multipolygons)
}

func (p *Polygon) Difference(others ...*Polygon) []*Polygon {
	geoms := ToGeoms(others...)
	multipolygons, _ := polygol.Difference(p.ToGeom(), geoms...)
	return FromGeom(multipolygons)
}

func (p *Polygon) XOR(others ...*Polygon) []*Polygon {
	geoms := ToGeoms(others...)
	multipolygons, _ := polygol.XOR(p.ToGeom(), geoms...)
	return FromGeom(multipolygons)
}

func (poly *Polygon) ContainsPoint(x, y float64) bool {
	bound := poly.GetBounds()
	if x < bound.Left || x > bound.Right || y < bound.Top || y > bound.Bottom {
		return false
	}
	// create a horizontal line that extends from the point
	lne := path.NewLine(x, y, bound.Right+21, y)

	// count the number of intersections with the polygon
	intersections := poly.LineIntersections(&lne.Path)

	// if the point is on the border, then the polygon contains the point
	for _, intersection := range intersections {
		// log.Println(intersection.X, intersection.Y, x, y)
		if intersection.X == x && intersection.Y == y {
			return true
		}
	}

	// if the number of intersections is odd, then the point is inside the polygon
	return len(intersections)%2 == 1
}

func GetTriangleCenter(p1, p2, p3 *point.Point) point.Point {
	return point.NewPoint((p1.X+p2.X+p3.X)/3, (p1.Y+p2.Y+p3.Y)/3)
}
