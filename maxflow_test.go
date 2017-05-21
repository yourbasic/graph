package graph

import (
	"math/rand"
	"testing"
)

func TestMaxFlow(t *testing.T) {
	g := New(1)
	g.AddCost(0, 0, 8)
	flow, res := MaxFlow(g, 0, 0)
	if mess, diff := diff(flow, Max); diff {
		t.Errorf("MaxFlow(0, 0) %s", mess)
	}
	if mess, diff := diff(String(res), "1 []"); diff {
		t.Errorf("MaxFlow(0, 0) %s", mess)
	}

	g = New(2)
	g.AddCost(0, 1, 5)
	flow, res = MaxFlow(g, 0, 1)
	if mess, diff := diff(flow, int64(5)); diff {
		t.Errorf("MaxFlow(0, 0) %s", mess)
	}
	if mess, diff := diff(String(res), "2 [(0 1):5]"); diff {
		t.Errorf("MaxFlow(0, 1) %s", mess)
	}

	g = New(6)
	for _, e := range []struct {
		v, w int
		c    int64
	}{
		{0, 1, 16}, {0, 2, 13}, {1, 2, 10}, {2, 1, 4},
		{1, 3, 12}, {2, 4, 14}, {3, 2, 9}, {4, 3, 7},
		{3, 5, 20}, {4, 5, 4},
	} {
		g.AddCost(e.v, e.w, e.c)
	}
	_, res = MaxFlow(g, 0, 5)
	exp := "6 [(0 1):12 (0 2):11 (1 3):12 (2 4):11 (3 5):19 (4 3):7 (4 5):4]"
	if mess, diff := diff(String(res), exp); diff {
		t.Errorf("MaxFlow(0, 5) %s", mess)
	}
	_, res = MaxFlow(g, 0, 1)
	exp = "6 [(0 1):16 (0 2):4 (2 1):4]"
	if mess, diff := diff(String(res), exp); diff {
		t.Errorf("MaxFlow(0, 1) %s", mess)
	}
	_, res = MaxFlow(g, 0, 2)
	exp = "6 [(0 1):16 (0 2):13 (1 2):10 (1 3):6 (3 2):6]"
	if mess, diff := diff(String(res), exp); diff {
		t.Errorf("MaxFlow(0, 2) %s", mess)
	}
	_, res = MaxFlow(g, 0, 3)
	exp = "6 [(0 1):12 (0 2):7 (1 3):12 (2 4):7 (4 3):7]"
	if mess, diff := diff(String(res), exp); diff {
		t.Errorf("MaxFlow(0, 3) %s", mess)
	}
	_, res = MaxFlow(g, 0, 4)
	exp = "6 [(0 1):1 (0 2):13 (1 2):1 (2 4):14]"
	if mess, diff := diff(String(res), exp); diff {
		t.Errorf("MaxFlow(0, 4) %s", mess)
	}
	_, res = MaxFlow(g, 3, 1)
	exp = "6 [(2 1):4 (3 2):4]"
	if mess, diff := diff(String(res), exp); diff {
		t.Errorf("MaxFlow(3, 1) %s", mess)
	}
}

func BenchmarkMaxFlow(b *testing.B) {
	n := 50
	b.StopTimer()
	g := New(n)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			g.AddCost(i, j, int64(rand.Int()))
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MaxFlow(g, 0, n-1)
	}
}
