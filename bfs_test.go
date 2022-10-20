package graph

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestBFS(t *testing.T) {
	g := New(10)
	for _, e := range []struct {
		v, w int
	}{
		{0, 1}, {0, 4}, {0, 7}, {0, 9},
		{4, 2}, {7, 5}, {7, 8},
		{2, 3}, {5, 6},
		{3, 6}, {8, 9}, {4, 4},
	} {
		g.AddBoth(e.v, e.w)
	}
	exp := "0147925836"
	res := "0"
	BFS(Sort(g), 0, func(v, w int, c int64) {
		res += strconv.Itoa(w)
	})
	if mess, diff := diff(res, exp); diff {
		t.Errorf("BFS: %s", mess)
	}
}

func BenchmarkBFS(b *testing.B) {
	n := 1000
	g := New(n)
	for i := 0; i < 5*n; i++ {
		g.AddBoth(rand.Intn(n), rand.Intn(n))
	}

	fn := func(v, w int, c int64) {}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BFS(g, 0, fn)
	}
}
