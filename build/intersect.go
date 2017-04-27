package build

// Intersect returns the graph intersection g1 ∩ g2,
// which consists of the intersection of the two vertex sets
// and the intersection of the two edge sets of g1 and g2.
// The edges of the new graph will have zero cost.
func (g1 *Virtual) Intersect(g2 *Virtual) *Virtual {
	n := min(g1.order, g2.order)
	switch n {
	case 0:
		return null
	case 1:
		return singleton()
	}

	var res *Virtual
	if g1.order == g2.order {
		res = generic0(n, func(v, w int) bool {
			return g1.edge(v, w) && g2.edge(v, w)
		})
	} else {
		res = generic0(n, func(v, w int) bool {
			return v < n && w < n && g1.edge(v, w) && g2.edge(v, w)
		})
	}

	res.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		w, n := findBoth(v, a, g1, g2)
		for n != 0 {
			for u := w; u < w+n; u++ {
				if do(u, 0) {
					return true
				}
			}
			w, n = findBoth(v, w+n, g1, g2)
		}
		return
	}
	return res
}

// findBoth returns the smallest n consecutive neighbors w, w+1, ..., w+n-1
// of v in g1 ∩ g2 for which w >= a.
func findBoth(v int, a int, g1, g2 *Virtual) (w int, n int) {
	w1, n1 := g1.find(v, a)
	w2, n2 := g2.find(v, a)
	for n1 != 0 && n2 != 0 {
		switch {
		case w1+n1 <= w2:
			w1, n1 = g1.find(v, w2)
		case w2+n2 <= w1:
			w2, n2 = g2.find(v, w1)
		case w1 < w2:
			return w2, min(w1+n1-w2, n2)
		default:
			return w1, min(w2+n2-w1, n1)
		}
	}
	return
}

// find returns the smallest n  consecutive neighbors w, w+1, ..., w+n-1
// of v for which w >= a.
func (g *Virtual) find(v int, a int) (w int, n int) {
	prev := -1
	g.visit(v, a, func(w0 int, c0 int64) (skip bool) {
		switch prev {
		case -1:
			w, n = w0, 1
			prev = w0
			return
		case w0 - 1:
			n += 1
			prev = w0
			return
		}
		return true
	})
	return
}
