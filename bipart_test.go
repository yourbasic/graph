package graph

import (
	"math/rand"
	"testing"
)

func TestBipartition(t *testing.T) {
	g := New(0)
	part, ok := Bipartition(g)
	if mess, diff := diff(part, []int{}); diff {
		t.Errorf("Bipartition: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("Bipartition: %s", mess)
	}

	g = New(1)
	part, ok = Bipartition(g)
	if mess, diff := diff(part, []int{0}); diff {
		t.Errorf("Bipartition: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("Bipartition: %s", mess)
	}

	g.Add(0, 0)
	part, ok = Bipartition(g)
	if mess, diff := diff(part, []int{}); diff {
		t.Errorf("Bipartition: %s", mess)
	}
	if mess, diff := diff(ok, false); diff {
		t.Errorf("Bipartition: %s", mess)
	}

	g = New(2)
	part, ok = Bipartition(g)
	if mess, diff := diff(part, []int{0, 1}); diff {
		t.Errorf("Bipartition: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("Bipartition: %s", mess)
	}

	g.AddBoth(0, 1)
	part, ok = Bipartition(g)
	if mess, diff := diff(part, []int{0}); diff {
		t.Errorf("Bipartition: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("Bipartition: %s", mess)
	}

	g.Add(0, 0)
	part, ok = Bipartition(g)
	if mess, diff := diff(part, []int{}); diff {
		t.Errorf("Bipartition: %s", mess)
	}
	if mess, diff := diff(ok, false); diff {
		t.Errorf("Bipartition: %s", mess)
	}

	g = New(5)
	g.Add(0, 1)
	g.Add(0, 2)
	g.Add(0, 3)
	part, ok = Bipartition(g)
	if mess, diff := diff(part, []int{0, 4}); diff {
		t.Errorf("Bipartition: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("Bipartition: %s", mess)
	}

	g.Add(2, 3)
	part, ok = Bipartition(g)
	if mess, diff := diff(part, []int{}); diff {
		t.Errorf("Bipartition: %s", mess)
	}
	if mess, diff := diff(ok, false); diff {
		t.Errorf("Bipartition: %s", mess)
	}

	g = New(5)
	g.AddBoth(0, 1)
	g.AddBoth(1, 2)
	g.AddBoth(2, 3)
	g.AddBoth(3, 0)
	part, ok = Bipartition(g)
	if mess, diff := diff(part, []int{0, 2, 4}); diff {
		t.Errorf("Bipartition: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("Bipartition: %s", mess)
	}
}

func BenchmarkBipartition(b *testing.B) {
	n := 1000
	g := New(n)
	for i := 0; i < 3*n; i++ {
		g.AddBoth(rand.Intn(n), rand.Intn(n))
	}
	h := Sort(g)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Bipartition(h)
	}
}
