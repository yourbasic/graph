package graph

import (
	"math/rand"
	"testing"
)

func TestTopSort(t *testing.T) {
	g := New(0)
	order, ok := TopSort(g)
	if mess, diff := diff(order, []int{}); diff {
		t.Errorf("TopSort %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("TopSort %s", mess)
	}

	g = New(4)
	g.Add(3, 2)
	g.Add(2, 1)
	g.Add(1, 0)
	expOrder := []int{3, 2, 1, 0}
	order, ok = TopSort(g)
	if mess, diff := diff(order, expOrder); diff {
		t.Errorf("TopSort %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("TopSort %s", mess)
	}

	g = New(5)
	g.Add(0, 1)
	g.Add(1, 2)
	g.Add(1, 3)
	g.Add(2, 4)
	g.Add(3, 4)
	order, ok = TopSort(g)
	expOrder1 := []int{0, 1, 2, 3, 4}
	expOrder2 := []int{0, 1, 3, 2, 4}
	mess, diff1 := diff(order, expOrder1)
	_, diff2 := diff(order, expOrder2)
	if diff1 && diff2 {
		t.Errorf("TopSort %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("TopSort %s", mess)
	}
}

func TestAcyclic(t *testing.T) {
	g := New(0)
	if mess, diff := diff(Acyclic(g), true); diff {
		t.Errorf("Acyclic %s", mess)
	}

	g = New(1)
	if mess, diff := diff(Acyclic(g), true); diff {
		t.Errorf("Acyclic %s", mess)
	}
	g.Add(0, 0)
	if mess, diff := diff(Acyclic(g), false); diff {
		t.Errorf("Acyclic %s", mess)
	}

	g = New(4)
	g.Add(0, 1)
	g.Add(2, 1)
	if mess, diff := diff(Acyclic(g), true); diff {
		t.Errorf("Acyclic %s", mess)
	}
	g.Add(0, 2)
	if mess, diff := diff(Acyclic(g), true); diff {
		t.Errorf("Acyclic %s", mess)
	}
	g.Add(3, 0)
	if mess, diff := diff(Acyclic(g), true); diff {
		t.Errorf("Acyclic %s", mess)
	}
	g.Add(1, 3)
	if mess, diff := diff(Acyclic(g), false); diff {
		t.Errorf("Acyclic %s", mess)
	}
}

func BenchmarkAcyclic(b *testing.B) {
	n := 1000
	b.StopTimer()
	g := New(n)
	for i := 0; i < 2*n; i++ {
		v, w := rand.Intn(n), rand.Intn(n)
		if v < w {
			g.AddBoth(v, w)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = Acyclic(g)
	}
}

func BenchmarkTopSort(b *testing.B) {
	n := 1000
	b.StopTimer()
	g := New(n)
	for i := 0; i < 2*n; i++ {
		v, w := rand.Intn(n), rand.Intn(n)
		if v < w {
			g.AddBoth(v, w)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TopSort(g)
	}
}
