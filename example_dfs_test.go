package graph_test

import (
	"fmt"
	"github.com/yourbasic/graph"
)

const (
	White = iota
	Gray
	Black
)

// The package doesn't support vertex labeling. However,
// since vertices are always numbered 0..n-1, it's easy
// to add this type of data on the side. This implementation
// of depth-first search uses separate slices to keep track of
// vertex colors, predecessors and discovery times.
type DFSData struct {
	Time     int
	Color    []int
	Prev     []int
	Discover []int
	Finish   []int
}

func DFS(g graph.Iterator) DFSData {
	n := g.Order() // Order returns the number of vertices.
	d := DFSData{
		Time:     0,
		Color:    make([]int, n),
		Prev:     make([]int, n),
		Discover: make([]int, n),
		Finish:   make([]int, n),
	}
	for v := 0; v < n; v++ {
		d.Color[v] = White
		d.Prev[v] = -1
	}
	for v := 0; v < n; v++ {
		if d.Color[v] == White {
			d.dfsVisit(g, v)
		}
	}
	return d
}

func (d *DFSData) dfsVisit(g graph.Iterator, v int) {
	d.Color[v] = Gray
	d.Time++
	d.Discover[v] = d.Time
	// Visit calls a function for each neighbor w of v,
	// with c equal to the cost of the edge (v, w).
	// The iteration is aborted if the function returns true.
	g.Visit(v, func(w int, c int64) (skip bool) {
		if d.Color[w] == White {
			d.Prev[w] = v
			d.dfsVisit(g, w)
		}
		return
	})
	d.Color[v] = Black
	d.Time++
	d.Finish[v] = d.Time
}

// Show how to use this package by implementing a complete depth-first search.
func Example_dFS() {
	// Build a small directed graph:
	//
	//   0 ---> 1 <--> 2     3
	//
	g := graph.New(4)
	g.Add(0, 1)
	g.AddBoth(1, 2)

	fmt.Println(g)
	fmt.Println(DFS(g))
	// Output:
	// 4 [(0 1) {1 2}]
	// {8 [2 2 2 2] [-1 0 1 -1] [1 2 3 7] [6 5 4 8]}
}
