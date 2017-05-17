package build_test

import (
	"fmt"
	"github.com/yourbasic/graph"
	"github.com/yourbasic/graph/build"
	"math"
)

// Find a shortest path going back and forth between
// two sets of points in the plane.
func Example_euclid() {
	type Point struct{ x, y int }

	// Euclidean distance.
	Euclid := func(p, q Point) float64 {
		xd := p.x - q.x
		yd := p.y - q.y
		return math.Sqrt(float64(xd*xd + yd*yd))
	}

	// 0     3
	// 1     4
	// 2     5
	points := []Point{
		{0, 0}, {0, 1}, {0, 2},
		{4, 0}, {4, 1}, {4, 2},
	}

	// Build a complete bipartite graph, connecting the three
	// left-hand points to the three points on the right,
	// and then apply a cost function to the edges of the graph.
	g := build.Kmn(3, 3).AddCostFunc(func(v, w int) int64 {
		// Distance to three decimal places.
		return int64(1000 * Euclid(points[v], points[w]))
	})

	// Find a shortest path from 0 to 2.
	path, dist := graph.ShortestPath(g, 0, 2)
	fmt.Println("path:", path, "length:", float64(dist)/1000)
	// Output:
	// path: [0 4 2] length: 8.246
}

// Find a maximum flow in a virtual grid graph.
func Example_maxflow() {
	// Build an undirected n×n grid with a silly edge cost.
	n := 10
	grid := build.Grid(n, n).AddCostFunc(func(v, w int) int64 {
		return int64((v + w) * (w - v))
	})

	// Keep only the forward-pointing edges.
	grid = grid.Keep(func(v, w int) bool { return v < w })

	// Join the top row, 0..n-1, of the grid to a source S.
	// The vertex in S will get index n*n in the new graph.
	// The new edges should point from S to the grid.
	// Give the new edges maximum capacity.
	S := build.Empty(1)
	grid_S := grid.Join(S, build.EdgeSet{
		From: build.Range(0, n),
		To:   build.Vertex(n * n),
		Keep: func(v, w int) bool { return v > w },
		Cost: build.Cost(graph.Max),
	})

	// Join the bottom row, n*(n-1)..n*n-1, of the grid to a sink T.
	// The vertex in T will get index n*n+1 in the new graph.
	// The new edges should point from the grid to T.
	// Give the new edges maximum capacity.
	T := build.Empty(1)
	grid_S_T := grid_S.Join(T, build.EdgeSet{
		From: build.Range(n*(n-1), n*n),
		To:   build.Vertex(n*n + 1),
		Keep: func(v, w int) bool { return v < w },
		Cost: build.Cost(graph.Max),
	})

	// Find the maximum flow from S to T.
	flow, _ := graph.MaxFlow(grid_S_T, n*n, n*n+1)
	fmt.Println(flow)
	// Output: 1900
}

// Filter a graph with a function.
func ExampleVirtual_Keep() {
	// The complete graph with four vertices.
	Kn4 := build.Kn(4)
	fmt.Println(Kn4)

	// Remove all edges incident with vertex 0.
	Kn4MinusZero := Kn4.Keep(func(v, w int) bool { return v != 0 && w != 0 })
	fmt.Println(Kn4MinusZero)
	// Output:
	// 4 [{0 1} {0 2} {0 3} {1 2} {1 3} {2 3}]
	// 4 [{1 2} {1 3} {2 3}]
}

// Compute the difference of two graphs.
func ExampleVirtual_Intersect() {
	// The complete graph with four vertices.
	Kn4 := build.Kn(4)
	fmt.Println(Kn4)

	// The circle graph with four vertices.
	Cn4 := build.Cycle(4)
	fmt.Println(Cn4)

	// Remove all edges belonging to Cn4 from Kn4.
	Kn4DiffCn4 := Kn4.Intersect(Cn4.Complement())
	fmt.Println(Kn4DiffCn4)
	// Output:
	// 4 [{0 1} {0 2} {0 3} {1 2} {1 3} {2 3}]
	// 4 [{0 1} {0 3} {1 2} {2 3}]
	// 4 [{0 2} {1 3}]
}

// Build a weighted barbell graph.
func ExampleVirtual_Connect_barbell() {
	// The n-barbell graph consists of two non-overlapping n-vertex cliques
	// together with a single edge that has an endpoint in each clique.
	n := 2
	bar := build.Grid(1, 2).AddCost(20)
	redPlate := build.Kn(n).AddCost(25)

	// Connect one plate to each end of the bar.
	barbell := bar.Connect(0, redPlate).Connect(1, redPlate)
	fmt.Println(barbell)
	// Output: 4 [{0 1}:20 {0 2}:25 {1 3}:25]
}

// Build a wheel graph.
func ExampleVirtual_Join_wheel() {
	// A wheel graph is a graph formed by joining a single vertex
	// to all vertices of a cycle.
	n := 32
	hub := build.Empty(1)
	rim := build.Cycle(n)
	wheel := hub.Join(rim, build.AllEdges())
	fmt.Println("Number of spokes:", wheel.Degree(0))
	// Output:
	// Number of spokes: 32
}

// Build a Petersen graph.
func ExampleVirtual_Match_petersen() {
	// The Petersen graph is most commonly drawn as a pentagon
	// with a pentagram inside, with five spokes.
	pentagon := build.Cycle(5)
	pentagram := pentagon.Complement()
	petersen := pentagon.Match(pentagram, build.AllEdges())
	fmt.Println(petersen)
	// Output:
	// 10 [{0 1} {0 4} {0 5} {1 2} {1 6} {2 3} {2 7} {3 4} {3 8} {4 9} {5 7} {5 8} {6 8} {6 9} {7 9}]
}

// Build a crown graph by multiplying Kₙ with K₂.
func ExampleVirtual_Tensor() {
	// The tensor product G × K₂ is a bipartite graph called the bipartite double cover of G.
	// A crown graph on 2n vertices is an undirected graph with two sets of vertices Uᵢ and Vⱼ
	// and with an edge from Uᵢ to Vⱼ whenever i ≠ j.
	// It can be computed as the bipartite double cover of Kn.
	k4 := build.Kn(4)
	crown8 := k4.Tensor(build.Kn(2))
	fmt.Println(crown8)
	// Output:
	// 8 [{0 3} {0 5} {0 7} {1 2} {1 4} {1 6} {2 5} {2 7} {3 4} {3 6} {4 7} {5 6}]
}

// Build a cube graph by multiplying a single edge with itself.
func ExampleVirtual_Cartesian() {
	edge := build.Grid(1, 2)
	square := edge.Cartesian(edge)
	cube := square.Cartesian(edge)

	fmt.Println(edge)
	fmt.Println(square)
	fmt.Println(cube)

	fmt.Println(graph.Equal(edge, build.Hyper(1)))
	fmt.Println(graph.Equal(square, build.Hyper(2)))
	fmt.Println(graph.Equal(cube, build.Hyper(3)))
	// Output:
	// 2 [{0 1}]
	// 4 [{0 1} {0 2} {1 3} {2 3}]
	// 8 [{0 1} {0 2} {0 4} {1 3} {1 5} {2 3} {2 6} {3 7} {4 5} {4 6} {5 7} {6 7}]
	// true
	// true
	// true
}

// Build a cube graph.
func ExampleVirtual_Match_cube() {
	// Build a cube graph.
	square := build.Grid(2, 2)
	cube := square.Match(square, build.AllEdges())
	fmt.Println(cube)
	fmt.Println(graph.Equal(cube, build.Hyper(3)))
	// Output:
	// 8 [{0 1} {0 2} {0 4} {1 3} {1 5} {2 3} {2 6} {3 7} {4 5} {4 6} {5 7} {6 7}]
	// true
}

// Build a directed graph containing all edges (v, w) for which v is odd and w even.
func ExampleGeneric() {
	// Define a graph by a function.
	g := build.Generic(10, func(v, w int) bool {
		// Include all edges with v odd and w even.
		return v%2 == 1 && w%2 == 0
	})

	// In a topological ordering of this graph,
	// odd numbered vertices must come before even.
	order, acyclic := graph.TopSort(g)
	fmt.Println("Acyclic:", acyclic)
	fmt.Println(order)
	// Output:
	// Acyclic: true
	// [1 3 5 7 9 0 2 4 6 8]
}
