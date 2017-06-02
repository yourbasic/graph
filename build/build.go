// Package build offers a tool for building virtual graphs.
//
// Virtual graphs
//
// In a virtual graph no vertices or edges are stored in memory,
// they are instead computed as needed. New virtual graphs are constructed
// by composing and filtering a set of standard graphs, or by writing
// functions that describe the edges of a graph.
// Multigraphs and graphs with self-loops are not suppported.
//
// Non-virtual graphs can be imported, and used as building blocks,
// by the Specific function. Virtual graphs don't need to be “exported‬”;
// they implement the Iterator interface and hence can be used directly
// by any algorithm in the graph package.
//
// Performance tips
//
// When possible, try to use predefined building blocks rather than
// filter functions. In particular, note that graphs built by the Generic
// function must visit all potenential neighbors during iteration.
//
// If space is readily available, you may use the Specific function
// to turn on caching for any component. This gives constant time
// performance for all basic operations on that component.
//
// Tutorial
//
// The Euclid and Maxflow examples show how to build graphs from
// standard components using composition and filtering. They also
// demonstrate how to apply a cost function to a virtual graph.
//
package build

import (
	"github.com/yourbasic/graph"
)

// Virtual represents a virtual graph.
// In a virtual graph no vertices or edges are stored in memory,
// they are instead computed as needed. New virtual graphs are constructed
// by composing and filtering a set of standard graphs, or by writing
// functions that describe the edges of a graph.
type Virtual struct {
	// The `order` field is, in fact, a constant function.
	// It returns the number of vertices in the graph.
	order int

	// The `edge` and `cost` functions define a weighted graph without self-loops.
	//
	//  • edge(v, w) returns true whenever (v, w) belongs to the graph;
	//    the value is disregarded when v == w.
	//
	//  • cost(v, w) returns the cost of (v, w);
	//    the value is disregarded when edge(v, w) is false.
	//
	edge func(v, w int) bool
	cost func(v, w int) int64

	// The `degree` and `visit` functions can be used to improve performance.
	// They MUST BE CONSISTENT with edge and cost. If not implemented,
	// the `generic` or `generic0` implementation is used instead.
	// The `Consistent` test function should be used to check compliance.
	//
	//  • degree(v) returns the outdegree of vertex v.
	//
	//  • visit(v) visits all neighbors w of v for which w ≥ a in
	//    NUMERICAL ORDER calling do(w, c) for edge (v, w) of cost c.
	//    If a call to do returns true, visit MUST ABORT the iteration
	//    and return true; if successful it should return false.
	//    Precondition: a ≥ 0.
	//
	degree func(v int) int
	visit  func(v int, a int, do func(w int, c int64) (skip bool)) (aborted bool)
}

// FilterFunc is a function that tells if there is a directed edge from v to w.
// The nil value represents an edge function that always returns true.
type FilterFunc func(v, w int) bool

// CostFunc is a function that computes the cost of an edge from v to w.
// The nil value represents a cost function that always returns 0.
type CostFunc func(v, w int) int64

// Cost returns a CostFunc that always returns n.
func Cost(n int64) CostFunc {
	return func(int, int) int64 { return n }
}

func neverEdge(int, int) bool  { return false }
func alwaysEdge(v, w int) bool { return v != w }

func zero(int, int) int64 { return 0 }

func degreeZero(int) int { return 0 }
func degreeOne(int) int  { return 1 }

func noNeighbors(int, int, func(w int, c int64) bool) bool { return false }

const bitsPerWord = 32 << uint(^uint(0)>>63)

func min(m, n int) int {
	if m > n {
		return n
	}
	return m
}

func max(m, n int) int {
	if m < n {
		return n
	}
	return m
}

// null is the null graph; a graph with no vertices.
var null = new(Virtual)

// singleton returns a graph with one vertex.
func singleton() *Virtual {
	return &Virtual{
		order:  1,
		edge:   neverEdge,
		cost:   zero,
		degree: degreeZero,
		visit:  noNeighbors,
	}
}

// edge returns a graph with two edges (0, 1) and (1, 0).
func edge() *Virtual {
	g := &Virtual{
		order:  2,
		cost:   zero,
		edge:   alwaysEdge,
		degree: degreeOne,
	}
	g.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		w := 1 - v
		if w < a {
			return
		}
		return do(w, 0)
	}
	return g
}

// line(n) returns the graph {0, 1}, {1, 2}, {2, 3},... , {n-2, n-1}.
func line(n int) *Virtual {
	switch {
	case n < 0:
		return nil
	case n == 0:
		return null
	case n == 1:
		return singleton()
	case n == 2:
		return edge()
	}
	g := generic0(n, func(v, w int) (edge bool) {
		switch v - w {
		case -1, 1:
			edge = true
		}
		return
	})
	g.degree = func(v int) int {
		switch v {
		case 0, n - 1:
			return 1
		default:
			return 2
		}
	}
	g.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		if w := v - 1; w >= a && do(w, 0) {
			return true
		}
		if w := v + 1; w >= a && w < n && do(w, 0) {
			return true
		}
		return
	}
	return g
}

// generic returns a standard implementation; cost and edge can't be nil.
func generic(n int, cost CostFunc, edge func(v, w int) bool) *Virtual {
	switch {
	case n < 0:
		return nil
	case n == 0:
		return null
	case n == 1:
		return singleton()
	}
	g := &Virtual{
		order: n,
		edge:  func(v, w int) bool { return v != w && edge(v, w) },
		cost:  cost,
	}
	g.degree = func(v int) (deg int) {
		g.visit(v, 0, func(int, int64) (skip bool) {
			deg++
			return
		})
		return
	}
	g.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		for w := a; w < n; w++ {
			if g.edge(v, w) && do(w, cost(v, w)) {
				return true
			}
		}
		return
	}
	return g
}

// generic0 returns a standard implementation; edge can't be nil.
func generic0(n int, edge func(v, w int) bool) *Virtual {
	switch {
	case n < 0:
		return nil
	case n == 0:
		return null
	case n == 1:
		return singleton()
	}
	g := &Virtual{
		order: n,
		edge:  func(v, w int) bool { return v != w && edge(v, w) },
		cost:  zero,
	}
	g.degree = func(v int) (deg int) {
		g.visit(v, 0, func(int, int64) (skip bool) {
			deg++
			return
		})
		return
	}
	g.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		for w := a; w < n; w++ {
			if g.edge(v, w) && do(w, 0) {
				return true
			}
		}
		return
	}
	return g
}

// Generic returns a virtual graph with n vertices; its edge set consists of
// all edges (v, w), v ≠ w, for which edge(v, w) returns true.
func Generic(n int, edge FilterFunc) *Virtual {
	switch {
	case n < 0:
		return nil
	case n == 0:
		return null
	case n == 1:
		return singleton()
	case edge == nil:
		return Kn(n)
	}
	return generic0(n, edge)
}

// Specific returns a cached copy of g with constant time performance for
// all basic operations. It uses space proportional to the size of the graph.
//
// This function does not accept multigraphs and graphs with self-loops.
func Specific(g graph.Iterator) *Virtual {
	h := graph.Sort(g)
	stats := graph.Check(h)
	if stats.Multi != 0 || stats.Loops != 0 {
		panic("Virtual doesn't support multiple edges or self-loops")
	}
	res := &Virtual{
		order:  h.Order(),
		edge:   h.Edge,
		visit:  h.VisitFrom,
		degree: h.Degree,
	}
	if stats.Weighted == 0 {
		res.cost = zero
		return res
	}
	res.cost = func(v, w int) (cost int64) {
		if !h.Edge(v, w) {
			return 0
		}
		h.VisitFrom(v, w, func(w int, c int64) (skip bool) {
			cost = c
			return true
		})
		return
	}
	return res
}

// Empty returns a virtual graph with n vertices and no edges.
func Empty(n int) *Virtual {
	switch {
	case n < 0:
		return nil
	case n == 0:
		return null
	case n == 1:
		return singleton()
	}
	return &Virtual{
		order:  n,
		edge:   neverEdge,
		cost:   zero,
		degree: degreeZero,
		visit:  noNeighbors,
	}
}

// Kn returns a complete simple graph with n vertices.
func Kn(n int) *Virtual {
	switch {
	case n < 0:
		return nil
	case n == 0:
		return null
	case n == 1:
		return singleton()
	}
	g := &Virtual{
		order:  n,
		edge:   alwaysEdge,
		cost:   zero,
		degree: func(v int) int { return n - 1 },
	}
	g.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		for w := a; w < g.order; w++ {
			if v != w && do(w, 0) {
				return true
			}
		}
		return
	}
	return g
}

// Complement returns the complement graph of g.
// This graph has the same vertices as g,
// but its edge set consists of the edges not present in g.
// The edges of the complement graph will have zero cost.
func (g *Virtual) Complement() *Virtual {
	n := g.order
	switch n {
	case 0:
		return null
	case 1:
		return singleton()
	}
	res := generic0(n, func(v, w int) (edge bool) {
		return v != w && !g.edge(v, w)
	})
	res.degree = func(v int) int { return n - 1 - g.degree(v) }
	res.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		prev := a
		if g.visit(v, a, func(w0 int, _ int64) (skip bool) {
			for w := prev; w < w0; w++ {
				if v != w && do(w, 0) {
					return true
				}
			}
			prev = w0 + 1
			return
		}) {
			return true
		}
		for w := prev; w < n; w++ {
			if v != w && do(w, 0) {
				return true
			}
		}
		return
	}
	return res
}

// Keep returns a graph containing all edges (v, w) of g for which edge(v, w) is true.
func (g *Virtual) Keep(edge FilterFunc) *Virtual {
	n := g.order
	switch {
	case n == 0:
		return null
	case n == 1:
		return singleton()
	case edge == nil:
		return g
	}
	res := generic(g.order, g.cost, func(v, w int) bool {
		return edge(v, w) && g.edge(v, w)
	})
	res.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		return g.visit(v, a, func(w int, c int64) bool {
			return edge(v, w) && do(w, c)
		})
	}
	return res
}

// AddCost returns a copy of g with a new cost assigned to all edges.
func (g *Virtual) AddCost(c int64) *Virtual {
	res := *g
	res.cost = Cost(c)
	res.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		return g.visit(v, a, func(w int, _ int64) bool {
			return do(w, c)
		})
	}
	return &res
}

// AddCostFunc returns a copy of g with a new cost function assigned.
func (g *Virtual) AddCostFunc(c CostFunc) *Virtual {
	if c == nil {
		h := g.AddCost(0)
		return h
	}
	res := *g
	res.cost = c
	res.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		return g.visit(v, a, func(w int, _ int64) bool {
			return do(w, c(v, w))
		})
	}
	return &res
}

// Order returns the number of vertices in the graph.
func (g *Virtual) Order() int {
	return g.order
}

// Degree returns the number of outward directed edges from v.
func (g *Virtual) Degree(v int) int {
	if v < 0 || v >= g.order {
		panic("vertex out of range")
	}
	return g.degree(v)
}

// Edge tells if there is an edge from v to w.
func (g *Virtual) Edge(v, w int) bool {
	if v < 0 || v >= g.order || w < 0 || w >= g.order {
		return false
	}
	return g.edge(v, w)
}

// Cost returns the cost of an edge from v to w, or 0 if no such edge exists.
func (g *Virtual) Cost(v, w int) int64 {
	if v < 0 || v >= g.order || w < 0 || w >= g.order {
		return 0
	}
	if g.edge(v, w) {
		return g.cost(v, w)
	}
	return 0
}

// Visit calls the do function for each neighbor w of v,
// with c equal to the cost of the edge from v to w.
// The neighbors are visited in increasing numerical order.
// If do returns true, Visit returns immediately,
// skipping any remaining neighbors, and returns true.
func (g *Virtual) Visit(v int, do func(w int, c int64) bool) bool {
	if v < 0 || v >= g.order {
		panic("vertex out of range")
	}
	return g.visit(v, 0, do)
}

// VisitFrom calls the do function starting from the first neighbor w
// for which w ≥ a, with c equal to the cost of the edge from v to w.
// The neighbors are then visited in increasing numerical order.
// If do returns true, VisitFrom returns immediately,
// skipping any remaining neighbors, and returns true.
func (g *Virtual) VisitFrom(v int, a int, do func(w int, c int64) bool) bool {
	n := g.order
	switch {
	case v < 0 || v >= n:
		panic("vertex out of range")
	case a < 0:
		a = 0
	case a > n:
		a = n
	}
	return g.visit(v, a, do)
}

// String returns a string representation of the graph.
func (g *Virtual) String() string {
	return graph.String(g)
}
