package build

// Match connects g1 to g2 by matching vertices in g1 with vertices in g2.
// Only vertices belonging to the bridge are included,
// and the vertices are matched in numerical order.
// The vertices of g2 are renumbered before the matching:
// vertex v ∊ g2 becomes v + g1.Order() in the new graph.
func (g1 *Virtual) Match(g2 *Virtual, bridge EdgeSet) *Virtual {
	n := g1.order + g2.order
	t := g1.order // transpose
	switch {
	case n == 0:
		return null
	case g1.order == 0:
		return g2
	case g2.order == 0:
		return g1
	}

	if bridge.Cost == nil {
		bridge.Cost = zero
	}
	matchCost := func(v, w int) int64 {
		switch {
		case v < t && w < t:
			return g1.cost(v, w)
		case v >= t && w >= t:
			return g2.cost(v-t, w-t)
		default:
			return bridge.Cost(v, w)
		}
	}

	if bridge.Keep == nil {
		bridge.Keep = alwaysEdge
	}
	s1 := bridge.From.And(Range(0, t))
	s2 := bridge.To.And(Range(t, g2.order+t))

	res := generic(n, matchCost, func(v, w int) (edge bool) {
		switch {
		case v < t && w < t:
			return g1.edge(v, w)
		case v >= t && w >= t:
			return g2.edge(v-t, w-t)
		case !bridge.Keep(v, w):
			return false
		default:
			s1v := s1.rank(v)
			s1w := s1.rank(w)
			s2v := s2.rank(v)
			s2w := s2.rank(w)
			return s1v != -1 && s1v == s2w || s1w != -1 && s1w == s2v
		}
	})

	res.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		if v < t {
			if g1.visit(v, a, do) {
				return true
			}
			if w := s2.get(s1.rank(v)); w >= a && bridge.Keep(v, w) {
				if do(w, bridge.Cost(v, w)) {
					return true
				}
			}
			return
		}
		if w := s1.get(s2.rank(v)); w >= a && bridge.Keep(v, w) {
			if do(w, bridge.Cost(v, w)) {
				return true
			}
		}
		return g2.visit(v-t, max(0, a-t), func(w int, c int64) bool {
			return do(w+t, c)
		})
	}
	return res
}
