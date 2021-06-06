package graph

import (
	"math/rand"
	"testing"
)

func TestShortestPath(t *testing.T) {
	g := New(6)
	g.AddCost(0, 1, 1)
	g.AddCost(0, 2, 1)
	g.AddCost(0, 3, 3)
	g.AddCost(1, 3, 0)
	g.AddCost(2, 3, 1)
	g.AddCost(2, 5, 8)
	g.AddCost(3, 5, 7)
	g.AddCost(1, 5, -1)
	parent, dist := ShortestPaths(g, 0)
	expParent := []int{-1, 0, 0, 1, -1, 3}
	expDist := []int64{0, 1, 1, 1, -1, 8}
	if mess, diff := diff(parent, expParent); diff {
		t.Errorf("ShortestPaths->parent %s", mess)
	}
	if mess, diff := diff(dist, expDist); diff {
		t.Errorf("ShortestPaths->dist %s", mess)
	}

	path, d := ShortestPath(g, 0, 5)
	if mess, diff := diff(path, []int{0, 1, 3, 5}); diff {
		t.Errorf("ShortestPath->path %s", mess)
	}
	if mess, diff := diff(d, int64(8)); diff {
		t.Errorf("ShortestPath->dist %s", mess)
	}

	path, d = ShortestPath(g, 0, 0)
	if mess, diff := diff(path, []int{0}); diff {
		t.Errorf("ShortestPath->path %s", mess)
	}
	if mess, diff := diff(d, int64(0)); diff {
		t.Errorf("ShortestPath->dist %s", mess)
	}

	path, d = ShortestPath(g, 0, 4)
	if mess, diff := diff(path, []int{}); diff {
		t.Errorf("ShortestPath->path %s", mess)
	}
	if mess, diff := diff(d, int64(-1)); diff {
		t.Errorf("ShortestPath->dist %s", mess)
	}
}

func randomGraph(n int) (*Mutable, int) {
	g := New(n)
	h := n / 2
	var t int
	for i := 0; i < n; i++ {
		g.Add(0, rand.Intn(n))
		if i == h {
			t = rand.Intn(n)
			g.Add(rand.Intn(n), t)
		} else {
			g.Add(rand.Intn(n), rand.Intn(n))
		}
	}
	return g, t
}

// Store benchmark results as global variables to prevent unwanted optimizations.
var path []int
var dist int64

func BenchmarkShortestPath250(b *testing.B) {
	g, t := randomGraph(250)
	b.ResetTimer()
	var p []int
	var d int64
	for i := 0; i < b.N; i++ {
		p, d = ShortestPath(g, 0, t)
	}
	path, dist = p, d
}

func BenchmarkShortestPath500(b *testing.B) {
	g, t := randomGraph(500)
	b.ResetTimer()
	var p []int
	var d int64
	for i := 0; i < b.N; i++ {
		p, d = ShortestPath(g, 0, t)
	}
	path, dist = p, d
}

func BenchmarkShortestPath1000(b *testing.B) {
	g, t := randomGraph(1000)
	b.ResetTimer()
	var p []int
	var d int64
	for i := 0; i < b.N; i++ {
		p, d = ShortestPath(g, 0, t)
	}
	path, dist = p, d
}

var (
	parent    []int
	distances []int64
)

func BenchmarkShortestPaths(b *testing.B) {
	n := 1000
	g := New(n)
	for i := 0; i < n; i++ {
		g.Add(0, rand.Intn(n))
		g.Add(rand.Intn(n), rand.Intn(n))
	}
	b.ResetTimer()
	var p []int
	var d []int64
	for i := 0; i < b.N; i++ {
		p, d = ShortestPaths(g, 0)
	}
	parent, distances = p, d
}
