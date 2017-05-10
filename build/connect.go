package build

// Connect connects g1 and g2 by making v1 in g1 a common vertex
// with 0 in g2. The vertices in g2 are renumbered: 0 becomes v1
// in the new graph; and w, with w > 0, becomes w + g1.Order() - 1.
//
func (g1 *Virtual) Connect(v1 int, g2 *Virtual) *Virtual {
	if g1.order == 0 || g2.order == 0 || v1 < 0 || v1 >= g1.order {
		return nil
	}

	n := g1.order + g2.order - 1
	t := g1.order - 1 // transpose
	if n == 1 {
		return singleton()
	}

	newCost := func(v, w int) int64 {
		switch {
		case v <= t && w <= t:
			return g1.cost(v, w)
		case v > t && w > t:
			return g2.cost(v-t, w-t)
		case v == v1:
			return g2.cost(0, w-t)
		case w == v1:
			return g2.cost(v-t, 0)
		default:
			return 0
		}
	}

	res := generic(n, newCost, func(v, w int) bool {
		switch {
		case v <= t && w <= t:
			return g1.edge(v, w)
		case v > t && w > t:
			return g2.edge(v-t, w-t)
		case v == v1:
			return g2.edge(0, w-t)
		case w == v1:
			return g2.edge(v-t, 0)
		default:
			return false
		}
	})

	res.degree = func(v int) (deg int) {
		switch {
		case v == v1:
			return g1.degree(v) + g2.degree(0)
		case v <= t:
			return g1.degree(v)
		default:
			return g2.degree(v - t)
		}
	}

	res.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		switch {
		case v > t:
			return g2.visit(v-t, max(0, a-t), func(w int, c int64) (skip bool) {
				if w == 0 {
					return v1 >= a && do(v1, c)
				}
				return do(w+t, c)
			})
		case g1.visit(v, a, do):
			return true
		case v == v1:
			return g2.visit(0, max(0, a-t), func(w int, c int64) (skip bool) {
				return do(w+t, c)
			})
		default:
			return
		}
	}
	return res
}
