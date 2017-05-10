package build

// Subgraph returns a subgraph of g that consists of the vertices in s
// and all edges whose endpoints are in s.
func (g *Virtual) Subgraph(s VertexSet) *Virtual {
	n := g.Order()
	s = s.And(Range(0, n))
	m := s.size()
	switch {
	case m == 0:
		return null
	case m == 1:
		return singleton()
	case m == n:
		return g
	}

	cost := func(v, w int) int64 {
		return g.cost(s.get(v), s.get(w))
	}
	res := generic(m, cost, func(v, w int) bool {
		return g.edge(s.get(v), s.get(w))
	})
	res.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		if a >= m {
			return
		}
		v0 := s.get(v)
		s0 := s.And(Range(s.get(a), n))
		for _, in := range s0.set {
			if more := false; g.visit(v0, in.a, func(w0 int, c int64) (skip bool) {
				switch {
				case w0 >= in.b:
					more, skip = true, true
					return
				case do(s.rank(w0), c):
					return true
				default:
					return
				}
			}) && !more {
				return true
			}
		}
		return
	}
	return res
}
