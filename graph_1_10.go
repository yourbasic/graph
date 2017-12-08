// +build go1.10

// Package graph contains generic implementations of basic graph algorithms.
//
// Generic graph algorithms
//
// The algorithms in this library can be applied to any graph data
// structure implementing the two Iterator methods: Order, which returns
// the number of vertices, and Visit, which iterates over the neighbors
// of a vertex.
//
// All algorithms operate on directed graphs with a fixed number
// of vertices, labeled from 0 to n-1, and edges with integer cost.
// An undirected edge {v, w} of cost c is represented by the two
// directed edges (v, w) and (w, v), both of cost c.
// A self-loop, an edge connecting a vertex to itself,
// is both directed and undirected.
//
// Graph data structures
//
// The type Mutable represents a directed graph with a fixed number
// of vertices and weighted edges that can be added or removed.
// The implementation uses hash maps to associate each vertex
// in the graph with its adjacent vertices. This gives constant
// time performance for all basic operations.
//
// The type Immutable is a compact representation of an immutable graph.
// The implementation uses lists to associate each vertex in the graph
// with its adjacent vertices. This makes for fast and predictable
// iteration: the Visit method produces its elements by reading
// from a fixed sorted precomputed list. This type supports multigraphs.
//
// Virtual graphs
//
// The subpackage graph/build offers a tool for building virtual graphs.
// In a virtual graph no vertices or edges are stored in memory,
// they are instead computed as needed. New virtual graphs are constructed
// by composing and filtering a set of standard graphs, or by writing
// functions that describe the edges of a graph.
//
// Tutorial
//
// The Basics example shows how to build  a plain graph and how to
// efficiently use the Visit iterator, the key abstraction of this package.
//
// The DFS example contains a full implementation of depth-first search.
//
package graph

import (
	"fmt"
	"sort"
	"strings"
)

// Iterator describes a weighted graph; an Iterator can be used
// to describe both ordinary graphs and multigraphs.
type Iterator interface {
	// Order returns the number of vertices in a graph.
	Order() int

	// Visit calls the do function for each neighbor w of vertex v,
	// with c equal to the cost of the edge from v to w.
	//
	//  • If do returns true, Visit returns immediately, skipping
	//    any remaining neighbors, and returns true.
	//
	//  • The calls to the do function may occur in any order,
	//    and the order may vary.
	//
	Visit(v int, do func(w int, c int64) (skip bool)) (aborted bool)
}

// The maximum and minum value of an edge cost.
const (
	Max int64 = 1<<63 - 1
	Min int64 = -1 << 63
)

type edge struct {
	v, w int
	c    int64
}

// String returns a description of g with two elements:
// the number of vertices, followed by a sorted list of all edges.
func String(g Iterator) string {
	n := g.Order()
	// This may be a multigraph, so we look for duplicates by counting.
	count := make(map[edge]int)
	for v := 0; v < n; v++ {
		g.Visit(v, func(w int, c int64) (skip bool) {
			count[edge{v, w, c}]++
			return
		})
	}
	edges := make([]edge, 0, len(count))
	for e := range count {
		edges = append(edges, e)
	}
	// Sort lexicographically on (v, w, c).
	sort.Slice(edges, func(i, j int) bool {
		v := edges[i].v == edges[j].v
		w := edges[i].w == edges[j].w
		switch {
		case v && w:
			return edges[i].c < edges[j].c
		case v:
			return edges[i].w < edges[j].w
		default:
			return edges[i].v < edges[j].v
		}
	})
	// Build the string.
	buf := new(strings.Builder)
	fmt.Fprintf(buf, "%d [", n)
	for i, e := range edges {
		c := count[e]
		if i != 0 && c > 0 {
			buf.WriteByte(' ')
		}
		if e.v < e.w {
			// Collect edges in opposite directions into an undirected edge.
			back := edge{e.w, e.v, e.c}
			m := min(c, count[back])
			count[back] -= m
			writeEdge(buf, e, m, true)
			if m > 0 && c-m > 0 {
				buf.WriteByte(' ')
			}
			writeEdge(buf, e, c-m, false)
		} else {
			writeEdge(buf, e, c, false)
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

func writeEdge(buf *strings.Builder, e edge, count int, bi bool) {
	if count <= 0 {
		return
	}
	if count > 1 {
		fmt.Fprintf(buf, "%d×", count)
	}
	if bi {
		buf.WriteByte('{')
	} else {
		buf.WriteByte('(')
	}
	fmt.Fprintf(buf, "%d %d", e.v, e.w)
	if bi {
		buf.WriteByte('}')
	} else {
		buf.WriteByte(')')
	}
	if e.c != 0 {
		buf.WriteByte(':')
		switch e.c {
		case Max:
			buf.WriteString("max")
		case Min:
			buf.WriteString("min")
		default:
			fmt.Fprintf(buf, "%d", e.c)
		}
	}
}

// Stats holds basic data about an Iterator.
type Stats struct {
	Size     int // Number of unique edges.
	Multi    int // Number of duplicate edges.
	Weighted int // Number of edges with non-zero cost.
	Loops    int // Number of self-loops.
	Isolated int // Number of vertices with outdegree zero.
}

// Check collects data about an Iterator.
func Check(g Iterator) Stats {
	if g, ok := g.(*Immutable); ok {
		return g.stats
	}
	_, mutable := g.(*Mutable)

	n := g.Order()
	degree := make([]int, n)
	type edge struct{ v, w int }
	edges := make(map[edge]bool)
	var stats Stats
	for v := 0; v < n; v++ {
		g.Visit(v, func(w int, c int64) (skip bool) {
			if w < 0 || w >= n {
				panic(fmt.Sprintf("vertex out of range: %d", w))
			}
			if v == w {
				stats.Loops++
			}
			if c != 0 {
				stats.Weighted++
			}
			degree[v]++
			if mutable { // A Mutable is never a multigraph.
				stats.Size++
				return
			}
			if edges[edge{v, w}] {
				stats.Multi++
			} else {
				stats.Size++
			}
			edges[edge{v, w}] = true
			return
		})
	}
	for _, deg := range degree {
		if deg == 0 {
			stats.Isolated++
		}
	}
	return stats
}

// Equal tells if g and h have the same number of vertices,
// and the same edges with the same costs.
func Equal(g, h Iterator) bool {
	if g.Order() != h.Order() {
		return false
	}
	edges := make(map[edge]int)
	for v := 0; v < g.Order(); v++ {
		g.Visit(v, func(w int, c int64) (skip bool) {
			edges[edge{v, w, c}]++
			return
		})
		if h.Visit(v, func(w int, c int64) (skip bool) {
			if edges[edge{v, w, c}] == 0 {
				return true
			}
			edges[edge{v, w, c}]--
			return
		}) {
			return false
		}
		for _, n := range edges {
			if n > 0 {
				return false
			}
		}
	}
	return true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
