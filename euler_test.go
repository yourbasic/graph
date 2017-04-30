package graph

import (
	"testing"
)

func TestEulerDirected(t *testing.T) {
	g := New(0)
	walk, ok := EulerDirected(g)
	if mess, diff := diff(walk, []int{}); diff {
		t.Errorf("EulerDirected: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("EulerDirected: %s", mess)
	}

	g = New(4)
	walk, ok = EulerDirected(g)
	if mess, diff := diff(walk, []int{}); diff {
		t.Errorf("EulerDirected: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("EulerDirected: %s", mess)
	}

	g.Add(0, 0)
	walk, ok = EulerDirected(g)
	if mess, diff := diff(walk, []int{0, 0}); diff {
		t.Errorf("EulerDirected: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("EulerDirected: %s", mess)
	}

	g.Add(0, 1)
	walk, ok = EulerDirected(g)
	if mess, diff := diff(walk, []int{0, 0, 1}); diff {
		t.Errorf("EulerDirected: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("EulerDirected: %s", mess)
	}

	g.Add(2, 3)
	walk, ok = EulerDirected(g)
	if mess, diff := diff(walk, []int{}); diff {
		t.Errorf("EulerDirected: %s", mess)
	}
	if mess, diff := diff(ok, false); diff {
		t.Errorf("EulerDirected: %s", mess)
	}

	g.Delete(2, 3)
	g.Add(2, 1)
	walk, ok = EulerDirected(g)
	if mess, diff := diff(walk, []int{}); diff {
		t.Errorf("EulerDirected: %s", mess)
	}
	if mess, diff := diff(ok, false); diff {
		t.Errorf("EulerDirected: %s", mess)
	}

	g.Delete(2, 1)
	g.Add(2, 2)
	walk, ok = EulerDirected(g)
	if mess, diff := diff(walk, []int{}); diff {
		t.Errorf("EulerDirected: %s", mess)
	}
	if mess, diff := diff(ok, false); diff {
		t.Errorf("EulerDirected: %s", mess)
	}

	g.Add(1, 2)
	walk, ok = EulerDirected(g)
	if mess, diff := diff(walk, []int{0, 0, 1, 2, 2}); diff {
		t.Errorf("EulerDirected: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("EulerDirected: %s", mess)
	}
}

func TestEulerUndirected(t *testing.T) {
	g := New(0)
	walk, ok := EulerUndirected(g)
	if mess, diff := diff(walk, []int{}); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}

	g = New(7)
	walk, ok = EulerUndirected(g)
	if mess, diff := diff(walk, []int{}); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}

	g.AddBoth(0, 0)
	walk, ok = EulerUndirected(g)
	if mess, diff := diff(walk, []int{0, 0}); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}

	g.AddBoth(0, 1)
	walk, ok = EulerUndirected(g)
	if mess, diff := diff(walk, []int{0, 0, 1}); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}

	g.AddBoth(2, 3)
	walk, ok = EulerUndirected(g)
	if mess, diff := diff(walk, []int{}); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}
	if mess, diff := diff(ok, false); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}

	g.AddBoth(1, 2)
	walk, ok = EulerUndirected(g)
	if mess, diff := diff(walk, []int{0, 0, 1, 2, 3}); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}

	g.AddBoth(2, 2)
	walk, ok = EulerUndirected(g)
	if mess, diff := diff(walk, []int{0, 0, 1, 2, 2, 3}); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}
	if mess, diff := diff(ok, true); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}

	g.AddBoth(4, 5)
	g.AddBoth(5, 6)
	g.AddBoth(4, 6)
	walk, ok = EulerUndirected(g)
	if mess, diff := diff(walk, []int{}); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}
	if mess, diff := diff(ok, false); diff {
		t.Errorf("EulerUndirected: %s", mess)
	}
}
