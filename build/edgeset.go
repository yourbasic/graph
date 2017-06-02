package build

// EdgeSet describes a set of edges; an edge (v, w), v ≠ w, belongs to the set
// if Keep(v, w) is true and (v, w) belongs to either From × To or To × From.
// The zero value of an edge set is the universe, the set containing all edges.
type EdgeSet struct {
	From, To VertexSet
	Keep     FilterFunc
	Cost     CostFunc
}

// AllEdges returns the universe, the set containing all edges.
// The edge cost is zero.
func AllEdges() EdgeSet {
	return EdgeSet{}
}

// NoEdges returns a set that includes no edges.
func NoEdges() EdgeSet {
	return EdgeSet{
		From: Range(0, 0),
		To:   Range(0, 0),
		Keep: neverEdge,
	}
}

// Edge returns a set consisting of a single edge {v, w}, v ≠ w, of zero cost.
func Edge(v, w int) EdgeSet {
	if v < 0 || w < 0 || v == w {
		return NoEdges()
	}
	return EdgeSet{
		From: Vertex(v),
		To:   Vertex(w),
	}
}

// Contains tells if the set contains the edge {v, w}.
func (e EdgeSet) Contains(v, w int) bool {
	switch {
	case e.Keep != nil && !e.Keep(v, w):
		return false
	case e.From.Contains(v) && e.To.Contains(w):
		return true
	case e.To.Contains(v) && e.From.Contains(w):
		return true
	default:
		return false
	}
}

// Add returns a graph containing all edges in g plus all edges in e.
// Any edges belonging to both g and e will retain their cost from g.
func (g *Virtual) Add(e EdgeSet) *Virtual {
	return g.union(newEdges(g.Order(), e), true)
}

// Delete returns a graph containing all edges in g except those also found in e.
func (g *Virtual) Delete(e EdgeSet) *Virtual {
	return g.Keep(func(v, w int) bool {
		return !e.Contains(v, w)
	})
}

// newEdges returns a virtual graph with n vertices and all edges
// belonging to the edge set.
func newEdges(n int, e EdgeSet) *Virtual {
	switch {
	case n < 0:
		return nil
	case n == 0:
		return null
	case n == 1:
		return singleton()
	}

	var noCost bool
	if e.Cost == nil {
		noCost = true
		e.Cost = zero
	}
	var noFilter bool
	if e.Keep == nil {
		noFilter = true
		e.Keep = alwaysEdge
	}

	from := e.From.And(Range(0, n))
	to := e.To.And(Range(0, n))
	if from.size() == 0 || to.size() == 0 {
		return Empty(n)
	}
	res := generic(n, e.Cost, func(v, w int) (edge bool) {
		return e.Contains(v, w)
	})

	intersect := from.And(to)
	union := from.Or(to)
	if noFilter {
		res.degree = func(v int) (deg int) {
			switch {
			case intersect.Contains(v):
				return union.size() - 1
			case from.Contains(v):
				return to.size()
			case to.Contains(v):
				return from.size()
			default:
				return
			}
		}
	}

	visit := func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		switch {
		case intersect.Contains(v):
			for _, in := range union.And(Range(a, n)).set {
				for w := in.a; w < in.b; w++ {
					if v != w && e.Keep(v, w) && do(w, e.Cost(v, w)) {
						return true
					}
				}
			}
			return
		case from.Contains(v):
			for _, in := range to.And(Range(a, n)).set {
				for w := in.a; w < in.b; w++ {
					if e.Keep(v, w) && do(w, e.Cost(v, w)) {
						return true
					}
				}
			}
			return
		case to.Contains(v):
			for _, in := range from.And(Range(a, n)).set {
				for w := in.a; w < in.b; w++ {
					if e.Keep(v, w) && do(w, e.Cost(v, w)) {
						return true
					}
				}
			}
			return
		default:
			return
		}
	}

	visit0 := func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		switch {
		case intersect.Contains(v):
			for _, in := range union.And(Range(a, n)).set {
				for w := in.a; w < in.b; w++ {
					if v != w && e.Keep(v, w) && do(w, 0) {
						return true
					}
				}
			}
			return
		case from.Contains(v):
			for _, in := range to.And(Range(a, n)).set {
				for w := in.a; w < in.b; w++ {
					if e.Keep(v, w) && do(w, 0) {
						return true
					}
				}
			}
			return
		case to.Contains(v):
			for _, in := range from.And(Range(a, n)).set {
				for w := in.a; w < in.b; w++ {
					if e.Keep(v, w) && do(w, 0) {
						return true
					}
				}
			}
			return
		default:
			return
		}
	}

	if noCost {
		res.visit = visit0
	} else {
		res.visit = visit
	}
	return res
}
