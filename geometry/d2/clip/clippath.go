package clip

import (
	"github.com/srmullen/godraw-lib/util"

	"github.com/srmullen/godraw-lib/geometry/d2/bounds"
	"github.com/srmullen/godraw-lib/geometry/d2/path"
)

// Given a Bounds and list of Paths, returns a list of Paths
// that do not extend outside of the bounds.
// Returns an array of *path.Path because path could exit and then
// reenter the bounds.
func ClipPath(bound bounds.Bounds, pth *path.Path) []*path.Path {
	ret := make([]*path.Path, 0)

	pb := pth.GetBounds()
	// Check if the path is entirly contained within the bounds.
	if bound.ContainsBounds(pb) {
		ret = append(ret, pth)
	} else if bound.Overlaps(*pb) {
		// Probably need to handle closed and unclosed paths differently

		// Handle open paths
		// Create a path that segements will be added to
		np := path.NewOpenPath([]float64{})

		previousInBound := false
		// Handle first segment
		seg := pth.Segments[0]

		if bound.Contains(seg.X, seg.Y) {
			ns := path.NewSegment(seg.X, seg.Y)
			np.Segments = append(np.Segments, ns)
			previousInBound = true
		}

		// # Segment creation rules (without curves)
		// 1. current and previous in bound
		// 		= add current to path
		// 2. current in bound, previous not in bound
		// 		= add bound intersection, add current
		// 3. current not in bound, previous not in bound
		// 		= check for intersections and add them to the path if they exist
		// 4. current not in bound, previous in bound
		// 		= add bound intersection to path

		// # Segment creation rules for curves.  [] = exclusive, () = inclusive
		// Unlike straight lines, curves can start and end in the bounds but still have intersections
		// 1. prev out, current out, 0 intersections
		// 		= do not add current to path
		// 2. prev in, current in, 0 intersections.
		// 		= add current to path
		// 3. prev in, current out, 1 intersection
		// 		= Add [prev .. first intersection)
		// 			Start a new path
		// 4. prev out, current in, 1 intersection
		// 		= Add (intersection .. current)
		// 5. prev in, current in, 2 intersections
		// 		= Add [prev .. first intersection)
		// 			Start new path.
		// 			Add (second intersection .. current)
		// 6. prev out, current out, 2 intersections
		// 		= Add (first intersection .. second intersection)
		// 			Start new path.
		// 7. prev in, current out, 3 intersection
		// 		= [prev .. first intersection)
		// 			Start new path.
		// 			(second intersection .. third intersection)
		// 			Start new path
		// 8. prev out, current in, 3 intersection
		// 		= (first intersection .. second intersection)
		// 			Start new path
		// 			(third intersection .. current)

		nsegs := len(pth.Segments)
		if pth.Closed {
			nsegs++
		}

		// TODO: need a way to calculate the best resolution value or set on curve.
		resolution := 20

		startNewPath := func() {
			ret = append(ret, np)
			np = path.NewOpenPath([]float64{})
		}

		interpolateTo := func(t float64, from, to path.Segment) {
			ts := util.Linspace(0, t, resolution)
			for i := 1; i < len(ts)-1; i++ {
				x, y := from.Curve.Interpolate(from.Point, to.Point, ts[i])
				np.Segments = append(np.Segments, path.NewSegment(x, y))
			}
		}

		interpolateFrom := func(t float64, from, to path.Segment) {
			ts := util.Linspace(t, 1, resolution)
			// start at 0 because inclusive
			for i := 0; i < len(ts)-1; i++ {
				x, y := from.Curve.Interpolate(from.Point, to.Point, ts[i])
				np.Segments = append(np.Segments, path.NewSegment(x, y))
			}
		}

		interpolateBetween := func(tfrom, tto float64, from, to path.Segment) {
			ts := util.Linspace(tfrom, tto, resolution)
			// start at 0 because inclusive
			for i := 0; i < len(ts); i++ {
				x, y := from.Curve.Interpolate(from.Point, to.Point, ts[i])
				np.Segments = append(np.Segments, path.NewSegment(x, y))
			}
		}

		// Find intersections for each segment
		for i := 1; i < nsegs; i++ {
			seg := pth.Segments[util.Mod(i, len(pth.Segments))]
			prev := pth.Segments[util.Mod(i-1, len(pth.Segments))]
			if prev.Curve != nil {
				// find curve intersections
				intersections := prev.BoundIntersections(seg.Point.X, seg.Point.Y, bound.Top, bound.Right, bound.Bottom, bound.Left)
				curInBound := bound.Contains(seg.X, seg.Y)
				if !previousInBound && !curInBound && len(intersections) == 0 {
					// 1
					// Do not add current to path
				} else if previousInBound && curInBound && len(intersections) == 0 {
					// 2
					// TODO: If seg is a curve, what should be done with it?.
					// Could add prev.Curve to the last segment of np.Segments, or generate all points between prev and cur segments.
					np.Segments = append(np.Segments, seg)
				} else if previousInBound && !curInBound && len(intersections) == 1 {
					// 3 = Add [prev .. first intersection)
					// 	 = Start a new path
					intersection := intersections[0]
					interpolateTo(intersection.T, prev, seg)
					next := path.NewSegment(intersection.X, intersection.Y)
					np.Segments = append(np.Segments, next)
					startNewPath()
				} else if !previousInBound && curInBound && len(intersections) == 1 {
					// 4 = Add (intersection .. current)
					intersection := intersections[0]
					interpolateFrom(intersection.T, prev, seg)
					np.Segments = append(np.Segments, path.NewSegment(seg.X, seg.Y))
				} else if previousInBound && curInBound && len(intersections) == 2 {
					// 5 = Add [prev .. first intersection)
					// 	 = Start new path.
					//	 = Add (second intersection .. current)
					i1 := intersections[0]
					interpolateTo(i1.T, prev, seg)
					startNewPath()
					i2 := intersections[1]
					np.Segments = append(np.Segments, path.NewSegment(i2.X, i2.Y))
					interpolateFrom(i2.T, prev, seg)
					// add current
					np.Segments = append(np.Segments, path.NewSegment(seg.X, seg.Y))
				} else if !previousInBound && !curInBound && len(intersections) == 2 {
					// 6 = Add (first intersection .. second intersection)
					// 			Start new path.
					interpolateBetween(intersections[0].T, intersections[1].T, prev, seg)
					startNewPath()
				} else if previousInBound && !curInBound && len(intersections) == 3 {
					// 7. prev in, current out, 3 intersection
					// 		= [prev .. first intersection)
					// 			Start new path.
					// 			(second intersection .. third intersection)
					// 			Start new path
					i1, i2, i3 := intersections[0], intersections[1], intersections[2]
					interpolateTo(i1.T, prev, seg)
					startNewPath()
					interpolateBetween(i2.T, i3.T, prev, seg)
					startNewPath()
				} else if !previousInBound && curInBound && len(intersections) == 3 {
					// 8. prev out, current in, 3 intersection
					// 		= (first intersection .. second intersection)
					// 			Start new path
					// 			(third intersection .. current)
					i1, i2, i3 := intersections[0], intersections[1], intersections[2]
					interpolateBetween(i1.T, i2.T, prev, seg)
					startNewPath()
					interpolateFrom(i3.T, prev, seg)
					np.Segments = append(np.Segments, path.NewSegment(seg.X, seg.Y))
				} else {
					panic("Clip Bounds: Unhandled segment")
				}
				previousInBound = curInBound
			} else {
				// Stright line between prev and seg
				if bound.Contains(seg.X, seg.Y) {
					if !previousInBound {
						intersections := prev.BoundIntersections(seg.Point.X, seg.Point.Y, bound.Top, bound.Right, bound.Bottom, bound.Left)
						if len(intersections) == 0 {
							panic("could not find intersection")
						} else if len(intersections) == 1 {
							// Add the intersection as a segment
							np.Segments = append(np.Segments, path.NewSegment(intersections[0].X, intersections[0].Y))
							np.Segments = append(np.Segments, seg)
						} else {
							// Multiple intersections
							// TODO: I think this this scenario is only possible with a curve
						}
					} else {
						np.Segments = append(np.Segments, seg)
					}
					previousInBound = true
				} else { // Out of bounds
					if previousInBound {
						prev := pth.Segments[util.Mod(i-1, len(pth.Segments))]
						// Find the intersection.
						intersections := prev.BoundIntersections(seg.Point.X, seg.Point.Y, bound.Top, bound.Right, bound.Bottom, bound.Left)
						if len(intersections) == 0 {
							panic("could not find intersection")
						}
						np.Segments = append(np.Segments, path.NewSegment(intersections[0].X, intersections[0].Y))
						// End the current path and start a new one
						ret = append(ret, np)
						np = path.NewOpenPath([]float64{})
					} else {
						// Previous and Current are not in bounds. There could still be intersections between them
						prev := pth.Segments[util.Mod(i-1, len(pth.Segments))]
						// Find the intersection.
						intersections := prev.BoundIntersections(seg.Point.X, seg.Point.Y, bound.Top, bound.Right, bound.Bottom, bound.Left)
						if len(intersections) > 0 {
							// Could there just be one intersection if it hits the corner?
							for _, p := range intersections {
								np.Segments = append(np.Segments, path.NewSegment(p.X, p.Y))
							}
						}
					}
					previousInBound = false
				}
			}
		}
		if len(np.Segments) > 0 {
			ret = append(ret, np)
		}
	} else {
		// Path is outside bounds. Don't include it in the returned values.
	}
	return ret
}
