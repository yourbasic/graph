package build

// Join joins g1 to g2 by adding edges from vertices in g1 to
// vertices in g2. Only the edges in the bridge are added.
// The vertices of g2 are renumbered before the operation:
// vertex v ∊ g2 becomes v + g1.Order() in the new graph.
func (g1 *Virtual) Join(g2 *Virtual, bridge EdgeSet) *Virtual {
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
	joinCost := func(v, w int) int64 {
		switch {
		case v < t && w < t:
			return g1.cost(v, w)
		case v >= t && w >= t:
			return g2.cost(v-t, w-t)
		default:
			return bridge.Cost(v, w)
		}
	}

	var noFilter bool
	if bridge.Keep == nil {
		noFilter = true
		bridge.Keep = alwaysEdge
	}
	s1 := bridge.From.And(Range(0, t))
	s2 := bridge.To.And(Range(t, g2.order+t))

	res := generic(n, joinCost, func(v, w int) (edge bool) {
		switch {
		case v < t && w < t:
			return g1.edge(v, w)
		case v >= t && w >= t:
			return g2.edge(v-t, w-t)
		case !bridge.Keep(v, w):
			return false
		default:
			return s1.Contains(v) && s2.Contains(w) ||
				s1.Contains(w) && s2.Contains(v)
		}
	})

	if noFilter {
		res.degree = func(v int) (deg int) {
			switch {
			case v < t:
				deg = g1.degree(v)
				if s1.Contains(v) {
					deg += s2.size()
				}
				return
			default:
				deg = g2.degree(v - t)
				if s2.Contains(v) {
					deg += s1.size()
				}
				return
			}
		}
	}

	res.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		if v < t {
			if g1.visit(v, a, do) {
				return true
			}
			if !s1.Contains(v) {
				return
			}
			for _, in := range s2.And(Range(a, n)).set {
				for w := in.a; w < in.b; w++ {
					if bridge.Keep(v, w) && do(w, bridge.Cost(v, w)) {
						return true
					}
				}
			}
			return
		}
		if s2.Contains(v) {
			for _, in := range s1.And(Range(a, n)).set {
				for w := in.a; w < in.b; w++ {
					if bridge.Keep(v, w) && do(w, bridge.Cost(v, w)) {
						return true
					}
				}
			}
		}
		return g2.visit(v-t, max(0, a-t), func(w int, c int64) bool {
			return do(w+t, c)
		})
	}
	return res
}
