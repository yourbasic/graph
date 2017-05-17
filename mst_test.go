package graph

import (
	"math/rand"
	"testing"
)

func TestMST(t *testing.T) {
	g := New(0)
	if mess, diff := diff(MST(g), []int{}); diff {
		t.Errorf("MST: %s", mess)
	}

	g = New(10)
	g.AddBothCost(0, 1, 4)
	g.AddBothCost(0, 7, 8)
	g.AddBothCost(1, 2, 8)
	g.AddBothCost(1, 7, 11)
	g.AddBothCost(2, 3, 7)
	g.AddBothCost(2, 8, 2)
	g.AddBothCost(2, 5, 4)
	g.AddBothCost(3, 4, 9)
	g.AddBothCost(3, 5, 14)
	g.AddBothCost(4, 5, 10)
	g.AddBothCost(5, 6, 2)
	g.AddBothCost(6, 7, 1)
	g.AddBothCost(6, 8, 6)
	g.AddBothCost(7, 8, 7)
	exp := []int{-1, 0, 5, 2, 3, 6, 7, 0, 2, -1}
	if mess, diff := diff(MST(g), exp); diff {
		t.Errorf("MST: %s", mess)
	}
}

func BenchmarkMST(b *testing.B) {
	n := 1000
	b.StopTimer()
	g := New(n)
	for i := 0; i < 2*n; i++ {
		g.AddCost(rand.Intn(n), rand.Intn(n), int64(rand.Int()))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = MST(g)
	}
}
