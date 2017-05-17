package graph

import (
	"math/rand"
	"testing"
)

func TestStrongComponents(t *testing.T) {
	g := New(0)
	if mess, diff := diff(StrongComponents(g), [][]int{}); diff {
		t.Errorf("StronglyConnected %s", mess)
	}

	g = New(10)
	g.Add(0, 1)
	g.Add(1, 2)
	g.Add(2, 0)
	g.Add(3, 1)
	g.Add(3, 2)
	g.Add(3, 5)
	g.Add(4, 2)
	g.Add(4, 6)
	g.Add(5, 3)
	g.Add(5, 4)
	g.Add(6, 4)
	g.Add(7, 5)
	g.Add(7, 6)
	g.Add(7, 7)
	g.Add(8, 8)
	exp := [][]int{{2, 1, 0}, {6, 4}, {5, 3}, {7}, {8}, {9}}
	if mess, diff := diff(StrongComponents(g), exp); diff {
		t.Errorf("StronglyConnected %s", mess)
	}
}

func BenchmarkStrongComponents(b *testing.B) {
	n := 1000
	b.StopTimer()
	g := New(n)
	for i := 0; i < n; i++ {
		g.Add(rand.Intn(n), rand.Intn(n))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = StrongComponents(g)
	}
}
