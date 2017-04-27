package build

// Union returns the graph union g1 ∪ g2, which
// consists of the union of the two vertex sets
// and the union of the two edge sets of g1 and g2.
// The edges of the new graph will have zero cost.
func (g1 *Virtual) Union(g2 *Virtual) *Virtual {
	return g1.union(g2, false)
}

// If cost is true keep costs and use g1's cost for edges that belong to both g1 and g2.
func (g1 *Virtual) union(g2 *Virtual, cost bool) *Virtual {
	switch {
	case g1.order == 0:
		if cost {
			return g2
		}
		return g2.AddCost(0)
	case g2.order == 0:
		if cost {
			return g1
		}
		return g1.AddCost(0)
	}

	newCost := zero
	if cost {
		newCost = func(v, w int) int64 {
			if g2.edge(v, w) && !g1.edge(v, w) {
				return g2.cost(v, w)
			}
			return g1.cost(v, w)
		}
	}
	var res *Virtual
	switch {
	case g1.order == g2.order:
		res = generic(g1.order, newCost, func(v, w int) bool {
			return g1.edge(v, w) || g2.edge(v, w)
		})
	case g1.order < g2.order:
		res = generic(g2.order, newCost, func(v, w int) bool {
			return v < g1.order && w < g1.order && g1.edge(v, w) || g2.edge(v, w)
		})
	default:
		res = generic(g1.order, newCost, func(v, w int) bool {
			return v < g2.order && w < g2.order && g2.edge(v, w) || g1.edge(v, w)
		})
	}

	res.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		next := 0
		if v < g1.order && g1.visit(v, a, func(w int, c int64) (skip bool) {
			// First all neighbors from g2 that are less than w...
			if next != w && v < g2.order && a <= w {
				if more := false; g2.visit(v, max(next, a), func(w0 int, c0 int64) (skip bool) {
					if w0 >= w {
						more, skip = true, true
						return
					}
					if cost {
						return do(w0, c0)
					}
					return do(w0, 0)
				}) && !more {
					return true
				}
			}
			// ...then w.
			switch {
			case cost && do(w, c):
				return true
			case !cost && do(w, 0):
				return true
			}
			next = w + 1
			return
		}) {
			return true
		}
		// When done with g1, produce any leftovers from g2.
		return v < g2.order && g2.visit(v, max(next, a), func(w int, c int64) (skip bool) {
			if cost {
				return do(w, c)
			}
			return do(w, 0)
		})
	}
	return res
}
