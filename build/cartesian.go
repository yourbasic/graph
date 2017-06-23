package build

import "strconv"

// Cartesian returns the cartesian product of g1 and g2:
// a graph whose vertices correspond to ordered pairs (v1, v2),
// where v1 and v2 are vertices in g1 and g2, respectively.
// The vertices (v1, v2) and (w1, w2) are connected by an edge if
// v1 = w1 and {v2, w2} ∊ g2 or v2 = w2 and {v1, w1} ∊ g1.
//
// In the new graph, vertex (v1, v2) gets index n⋅v1 + v2, where n = g2.Order(),
// and index i corresponds to the vertice (i/n, i%n).
func (g1 *Virtual) Cartesian(g2 *Virtual) *Virtual {
	m, n := g1.Order(), g2.Order()
	switch {
	case m == 0 || n == 0:
		return null
	case m*n/m != n:
		panic("too large m=" + strconv.Itoa(m) + " n=" + strconv.Itoa(n))
	}
	order := m * n
	g := generic0(order, func(v, w int) (edge bool) {
		v1, v2 := v/n, v%n
		w1, w2 := w/n, w%n
		return v1 == w1 && g2.Edge(v2, w2) || v2 == w2 && g1.Edge(v1, w1)
	})

	g.degree = func(v int) (deg int) {
		return g1.degree(v/n) + g2.degree(v%n)
	}

	g.visit = func(v int, a int, do func(w int, _ int64) bool) (aborted bool) {
		v1, v2 := v/n, v%n
		a1 := a / n
		if a1 < v1 {
			if more := false; g1.visit(v1, a1, func(w1 int, _ int64) (skip bool) {
				if w1 >= v1 {
					more, skip = true, true
					return
				}
				w := n*w1 + v2 // v2 == w2
				return w >= a && do(w, 0)
			}) && !more {
				return true
			}
		}
		if a1 <= v1 {
			a2 := 0
			if a1 == v1 {
				a2 = a % n
			}
			if g2.visit(v2, a2, func(w2 int, _ int64) (skip bool) {
				return do(n*v1+w2, 0) // v1 == w1
			}) {
				return true
			}
		}
		return g1.visit(v1, max(a1, v1+1), func(w1 int, _ int64) (skip bool) {
			w := n*w1 + v2 // v2 == w2
			return w >= a && do(w, 0)
		})
	}
	return g
}
