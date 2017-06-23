package build

import "strconv"

// Tensor returns the tensor product of g1 and g2:
// a graph whose vertices correspond to ordered pairs (v1, v2),
// where v1 and v2 are vertices in g1 and g2, respectively.
// The vertices (v1, v2) and (w1, w2) are connected by an edge whenever
// both of the edges {v1, w1} and {v2, w2} exist in the original graphs.
//
// In the new graph, vertex (v1, v2) gets index n⋅v1 + v2, where n = g2.Order(),
// and index i corresponds to the vertice (i/n, i%n).
func (g1 *Virtual) Tensor(g2 *Virtual) *Virtual {
	m, n := g1.Order(), g2.Order()
	switch {
	case m < 0 || n < 0:
		return nil
	case m == 0 || n == 0:
		return null
	case m*n/m != n:
		panic("too large m=" + strconv.Itoa(m) + " n=" + strconv.Itoa(n))
	}

	g := generic0(m*n, func(v, w int) (edge bool) {
		v1, v2 := v/n, v%n
		w1, w2 := w/n, w%n
		return g1.Edge(v1, w1) && g2.Edge(v2, w2)
	})

	g.degree = func(v int) (deg int) {
		v1, v2 := v/n, v%n
		return g1.degree(v1) * g2.degree(v2)
	}

	g.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		v1, v2 := v/n, v%n
		a1, a2 := a/n, a%n
		return g1.visit(v1, a1, func(w1 int, c int64) (skip bool) {
			if w1 == a1 {
				return g2.visit(v2, a2, func(w2 int, c int64) (skip bool) {
					return do(n*w1+w2, 0)
				})
			}
			return g2.visit(v2, 0, func(w2 int, c int64) (skip bool) {
				return do(n*w1+w2, 0)
			})
		})
	}
	return g
}
