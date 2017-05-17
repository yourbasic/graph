package graph

import (
	"math/rand"
	"testing"
)

func TestComponents(t *testing.T) {
	g := New(0)
	if mess, diff := diff(Components(g), [][]int{}); diff {
		t.Errorf("Components %s", mess)
	}
	if mess, diff := diff(Connected(g), false); diff {
		t.Errorf("Connected %s", mess)
	}

	g = New(1)
	if mess, diff := diff(Components(g), [][]int{{0}}); diff {
		t.Errorf("Components %s", mess)
	}
	if mess, diff := diff(Connected(g), true); diff {
		t.Errorf("Connected %s", mess)
	}

	g.Add(0, 0)
	if mess, diff := diff(Components(g), [][]int{{0}}); diff {
		t.Errorf("Components %s", mess)
	}
	if mess, diff := diff(Connected(g), true); diff {
		t.Errorf("Connected %s", mess)
	}

	g = New(4)
	g.Add(0, 1)
	g.Add(2, 1)
	if mess, diff := diff(Components(g), [][]int{{0, 1, 2}, {3}}); diff {
		t.Errorf("Components %s", mess)
	}
	if mess, diff := diff(Connected(g), false); diff {
		t.Errorf("Connected %s", mess)
	}

	g.AddBoth(0, 1)
	g.AddBoth(1, 2)
	g.AddBoth(2, 3)
	g.AddBoth(0, 3)
	if mess, diff := diff(Components(g), [][]int{{0, 1, 2, 3}}); diff {
		t.Errorf("Components %s", mess)
	}
	if mess, diff := diff(Connected(g), true); diff {
		t.Errorf("Connected %s", mess)
	}
}

func BenchmarkConnected(b *testing.B) {
	n := 1000
	b.StopTimer()
	g := New(n)
	for i := 0; i < n; i++ {
		g.AddBoth(rand.Intn(n), rand.Intn(n))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = Connected(g)
	}
}

func BenchmarkComponents(b *testing.B) {
	n := 1000
	b.StopTimer()
	g := New(n)
	for i := 0; i < n; i++ {
		g.AddBoth(rand.Intn(n), rand.Intn(n))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = Components(g)
	}
}
