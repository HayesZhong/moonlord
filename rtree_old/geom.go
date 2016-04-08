package rtree_old

import (
	"fmt"
	"math"
	"strings"
)

// DimError represents a failure due to mismatched dimensions.
type DimError struct {
	Expected int
	Actual   int
}

func (err DimError) Error() string {
	return "rtreego: dimension mismatch"
}

// DistError is an improper distance measurement.  It implements the error
// and is generated when a distance-related assertion fails.
type DistError float64

func (err DistError) Error() string {
	return "rtreego: improper distance"
}

// Point represents a point in n-dimensional Euclidean space.
type Point struct {
	data []float64
	Obj  interface{}
}

func NewPoint(dim int, obj interface{}) *Point {
	return &Point{
		data: make([]float64, dim),
		Obj:  obj,
	}
}

func NewPointWithData(data []float64, obj interface{}) *Point {
	return &Point{
		data: data,
		Obj:  obj,
	}
}

func (p Point) SetObj(obj *interface{}) {
	p.Obj = obj
}

func (p Point) set(i int, value float64) {
	p.data[i] = value
}

func (p Point) indexValue(i int) float64 {
	return p.data[i]
}
func (p Point) length() int {
	return len(p.data)
}

// Dist computes the Euclidean distance between two points p and q.
func (p Point) dist(q Point) float64 {
	if p.length() != q.length() {
		panic(DimError{p.length(), q.length()})
	}
	sum := 0.0
	for i := 0; i < p.length(); i++ {
		dx := p.data[i] - q.data[i]
		sum += dx * dx
	}
	return math.Sqrt(sum)
}

// minDist computes the square of the distance from a point to a rectangle.
// If the point is contained in the rectangle then the distance is zero.
//
// Implemented per Definition 2 of "Nearest Neighbor Queries" by
// N. Roussopoulos, S. Kelley and F. Vincent, ACM SIGMOD, pages 71-79, 1995.
func (p Point) minDist(r *Rect) float64 {
	if p.length() != r.p.length() {
		panic(DimError{p.length(), r.p.length()})
	}

	sum := 0.0
	var pi float64
	for i := 0; i < p.length(); i++ {
		pi = p.data[i]
		if pi < r.p.data[i] {
			d := pi - r.p.data[i]
			sum += d * d
		} else if pi > r.q.data[i] {
			d := pi - r.q.data[i]
			sum += d * d
		} else {
			sum += 0
		}
	}
	return sum
}

// minMaxDist computes the minimum of the maximum distances from p to points
// on r.  If r is the bounding box of some geometric objects, then there is
// at least one object contained in r within minMaxDist(p, r) of p.
//
// Implemented per Definition 4 of "Nearest Neighbor Queries" by
// N. Roussopoulos, S. Kelley and F. Vincent, ACM SIGMOD, pages 71-79, 1995.
func (p Point) minMaxDist(r *Rect) float64 {
	if p.length() != r.p.length() {
		panic(DimError{p.length(), r.p.length()})
	}

	// by definition, MinMaxDist(p, r) =
	// min{1<=k<=n}(|pk - rmk|^2 + sum{1<=i<=n, i != k}(|pi - rMi|^2))
	// where rmk and rMk are defined as follows:

	rm := func(k int) float64 {
		if p.indexValue(k) <= (r.p.indexValue(k)+r.q.indexValue(k))/2 {
			return r.p.indexValue(k)
		}
		return r.q.indexValue(k)
	}

	rM := func(k int) float64 {
		if p.indexValue(k) >= (r.p.indexValue(k)+r.q.indexValue(k))/2 {
			return r.p.indexValue(k)
		}
		return r.q.indexValue(k)
	}

	// This formula can be computed in linear time by precomputing
	// S = sum{1<=i<=n}(|pi - rMi|^2).

	S := 0.0
	for i := range p.data {
		d := p.data[i] - rM(i)
		S += d * d
	}

	// Compute MinMaxDist using the precomputed S.
	min := math.MaxFloat64
	for k := range p.data {
		d1 := p.indexValue(k) - rM(k)
		d2 := p.indexValue(k) - rm(k)
		d := S - d1*d1 + d2*d2
		if d < min {
			min = d
		}
	}

	return min
}

// Rect represents a subset of n-dimensional Euclidean space of the form
// [a1, b1] x [a2, b2] x ... x [an, bn], where ai < bi for all 1 <= i <= n.
type Rect struct {
	p, q *Point // Enforced by NewRect: p.data[i] <= q.data[i] for all i.
}

func (r *Rect) GetP() *Point {
	return r.p
}
func (r *Rect) GetQ() *Point {
	return r.q
}

func (r *Rect) Bounds() *Rect {
	return r
}

// The coordinate of the point of the rectangle at i
func (r *Rect) PointCoord(i int) float64 {
	return r.p.data[i]
}

// The coordinate of the lengths of the rectangle at i
func (r *Rect) LengthsCoord(i int) float64 {
	return r.q.data[i] - r.p.data[i]
}

// Returns true if the two rectangles are equal
func (r *Rect) Equal(other *Rect) bool {
	for i := 0; i < r.p.length(); i++ {
		if r.p.data[i] != other.p.data[i] {
			return false
		}
	}
	for i := 0; i < r.q.length(); i++ {
		if r.q.data[i] != other.q.data[i] {
			return false
		}
	}
	return true
}

func (r *Rect) String() string {
	s := make([]string, r.p.length())
	for i := 0; i < r.p.length(); i++ {
		a := r.p.data[i]
		b := r.q.data[i]
		s[i] = fmt.Sprintf("[%.2f, %.2f]", a, b)
	}
	return strings.Join(s, "x")
}

//hayes
func NewRectWithTwoPoint(p *Point, q *Point) (r *Rect) {
	return &Rect{
		p: p,
		q: q,
	}
}

// NewRect constructs and returns a pointer to a Rect given a corner point and
// the lengths of each dimension.  The point p should be the most-negative point
// on the rectangle (in every dimension) and every length should be positive.
func NewRect(p *Point, lengths []float64) (r *Rect) {
	r = new(Rect)
	r.p = p
	r.q = NewPoint(p.length(), nil)
	for i := 0; i < p.length(); i++ {
		r.q.set(i, p.data[i]+lengths[i])
	}
	return
}

// size computes the measure of a rectangle (the product of its side lengths).
func (r *Rect) size() float64 {
	size := 1.0
	for i := 0; i < r.p.length(); i++ {
		a := r.p.data[i]
		b := r.q.data[i]
		size *= b - a
	}
	return size

}

// margin computes the sum of the edge lengths of a rectangle.
func (r *Rect) margin() float64 {
	// The number of edges in an n-dimensional rectangle is n * 2^(n-1)
	// (http://en.wikipedia.org/wiki/Hypercube_graph).  Thus the number
	// of edges of length (ai - bi), where the rectangle is determined
	// by p = (a1, a2, ..., an) and q = (b1, b2, ..., bn), is 2^(n-1).
	//
	// The margin of the rectangle, then, is given by the formula
	// 2^(n-1) * [(b1 - a1) + (b2 - a2) + ... + (bn - an)].
	dim := r.p.length()
	sum := 0.0
	for i := 0; i < r.p.length(); i++ {
		a := r.p.data[i]
		b := r.q.data[i]
		sum += b - a
	}
	return math.Pow(2, float64(dim-1)) * sum
}

// containsPoint tests whether p is located inside or on the boundary of r.
func (r *Rect) containsPoint(p Point) bool {
	if p.length() != r.p.length() {
		panic(DimError{r.p.length(), p.length()})
	}

	for i, a := range p.data {
		// p is contained in (or on) r if and only if p <= a <= q for
		// every dimension.
		if a < r.p.data[i] || a > r.q.data[i] {
			return false
		}
	}

	return true
}

// containsRect tests whether r2 is is located inside r1.
func (r1 *Rect) containsRect(r2 *Rect) bool {
	if r1.p.length() != r2.p.length() {
		panic(DimError{r1.p.length(), r2.p.length()})
	}

	for i := 0; i < r1.q.length(); i++ {
		a1, b1, a2, b2 := r1.p.data[i], r1.q.data[i], r2.p.data[i], r2.q.data[i]
		// enforced by constructor: a1 <= b1 and a2 <= b2.
		// so containment holds if and only if a1 <= a2 <= b2 <= b1
		// for every dimension.
		if a1 > a2 || b2 > b1 {
			return false
		}
	}

	return true
}

// intersect computes the intersection of two rectangles.  If no intersection
// exists, the intersection is nil.
func intersect(r1, r2 *Rect) *Rect {
	dim := r1.p.length()
	if r2.p.length() != dim {
		panic(DimError{dim, r2.p.length()})
	}

	// There are four cases of overlap:
	//
	//     1.  a1------------b1
	//              a2------------b2
	//              p--------q
	//
	//     2.       a1------------b1
	//         a2------------b2
	//              p--------q
	//
	//     3.  a1-----------------b1
	//              a2-------b2
	//              p--------q
	//
	//     4.       a1-------b1
	//         a2-----------------b2
	//              p--------q
	//
	// Thus there are only two cases of non-overlap:
	//
	//     1. a1------b1
	//                    a2------b2
	//
	//     2.             a1------b1
	//        a2------b2
	//
	// Enforced by constructor: a1 <= b1 and a2 <= b2.  So we can just
	// check the endpoints.

	p := NewPoint(dim, nil)
	q := NewPoint(dim, nil)

	for i := range p.data {
		a1, b1, a2, b2 := r1.p.data[i], r1.q.data[i], r2.p.data[i], r2.q.data[i]
		if b2 <= a1 || b1 <= a2 {
			return nil
		}
		p.data[i] = math.Max(a1, a2)
		q.data[i] = math.Min(b1, b2)
	}
	return &Rect{p, q}
}

// ToRect constructs a rectangle containing p with side lengths 2*tol.
func (p Point) ToRect(tol float64) *Rect {
	dim := p.length()
	a, b := NewPoint(dim, nil), NewPoint(dim, nil)
	for i := range p.data {
		a.data[i] = p.data[i] - tol
		b.data[i] = p.data[i] + tol
	}
	return &Rect{a, b}
}

// boundingBox constructs the smallest rectangle containing both r1 and r2.
func boundingBox(r1, r2 *Rect) (bb *Rect) {
	bb = new(Rect)
	dim := r1.p.length()
	bb.p = NewPoint(dim, nil)
	bb.q = NewPoint(dim, nil)
	if r2.p.length() != dim {
		panic(DimError{dim, r2.p.length()})
	}
	for i := 0; i < dim; i++ {
		if r1.p.data[i] <= r2.p.data[i] {
			bb.p.set(i, r1.p.data[i])
		} else {
			bb.p.set(i, r2.p.data[i])
		}
		if r1.q.data[i] <= r2.q.data[i] {
			bb.q.set(i, r2.q.data[i])
		} else {
			bb.q.set(i, r1.q.data[i])
		}
	}
	return
}

// boundingBoxN constructs the smallest rectangle containing all of r...
func boundingBoxN(rects ...*Rect) (bb *Rect) {
	if len(rects) == 1 {
		bb = rects[0]
		return
	}
	bb = boundingBox(rects[0], rects[1])
	for _, rect := range rects[2:] {
		bb = boundingBox(bb, rect)
	}
	return
}
