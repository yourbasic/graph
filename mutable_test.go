package graph

import (
	"math/rand"
	"testing"
)

func setup() (g0, g1, g5 *Mutable) {
	g0 = New(0)

	g1 = New(1)
	g1.Add(0, 0)

	g5 = New(5)
	g5.Add(0, 1)
	g5.AddCost(2, 3, 1)
	return
}

func TestNew(t *testing.T) {
	g0, g1, g5 := setup()

	if mess, diff := diff(String(g0), "0 []"); diff {
		t.Errorf("New g0 %s", mess)
	}
	Consistent("New g0", t, g0)

	if mess, diff := diff(String(g1), "1 [(0 0)]"); diff {
		t.Errorf("New g1 %s", mess)
	}
	Consistent("New g1", t, g1)

	if mess, diff := diff(String(g5), "5 [(0 1) (2 3):1]"); diff {
		t.Errorf("New g5 %s", mess)
	}
	Consistent("New g5", t, g5)

	n := 10
	g := New(n)
	for i := 0; i < 2*n; i++ {
		g.AddBothCost(rand.Intn(n), rand.Intn(n), rand.Int63())
	}
	Consistent("Mutable rand", t, g)
}

func TestCopy(t *testing.T) {
	g0, g1, g5 := setup()

	res := Copy(g0)
	if mess, diff := diff(g0, res); diff {
		t.Errorf("Copy mutable g0 %s", mess)
	}
	Consistent("Copy mutable g0", t, res)

	res = Copy(g1)
	if mess, diff := diff(g1, res); diff {
		t.Errorf("Copy mutable g1 %s", mess)
	}
	Consistent("Copy mutable g1", t, res)

	res = Copy(g5)
	if mess, diff := diff(g5, res); diff {
		t.Errorf("Copy mutable g5 %s", mess)
	}
	Consistent("Copy mutable g5", t, res)

	res = Copy(Sort(g0))
	if mess, diff := diff(g0, res); diff {
		t.Errorf("Copy immutable g0 %s", mess)
	}
	Consistent("Copy immutable g0", t, res)

	res = Copy(Sort(g1))
	if mess, diff := diff(g1, res); diff {
		t.Errorf("Copy immutable g1 %s", mess)
	}
	Consistent("Copy immutable g1", t, res)

	res = Copy(Sort(g5))
	if mess, diff := diff(g5, res); diff {
		t.Errorf("Copy immutable g5 %s", mess)
	}
	Consistent("Copy immutable g5", t, res)

	res = Copy(Multi{})
	exp := "3 [(0 0):5 (0 1):7 (1 0):5]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Copy iterator Multi %s", mess)
	}
	Consistent("Copy iterator Multi", t, res)

	n := 10
	g := New(n)
	for i := 0; i < 2*n; i++ {
		g.AddBothCost(rand.Intn(n), rand.Intn(n), rand.Int63())
	}
	Consistent("Copy mutable rand", t, Copy(g))
	Consistent("Copy immutable rand", t, Copy(Sort(g)))
	Consistent("Copy iterator rand", t, Copy(Iterator(g)))
}

func TestOrder(t *testing.T) {
	g0, g1, g5 := setup()
	s := "Order()"

	if mess, diff := diff(g0.Order(), 0); diff {
		t.Errorf("g0.%s %s", s, mess)
	}
	if mess, diff := diff(g1.Order(), 1); diff {
		t.Errorf("g1.%s %s", s, mess)
	}
	if mess, diff := diff(g5.Order(), 5); diff {
		t.Errorf("g5.%s %s", s, mess)
	}
}

func TestEdge(t *testing.T) {
	_, g1, g5 := setup()

	if mess, diff := diff(g1.Edge(-1, 0), false); diff {
		t.Errorf("g1.Edge(-1, 0) %s", mess)
	}
	if mess, diff := diff(g1.Edge(1, 0), false); diff {
		t.Errorf("g1.Edge(1, 0) %s", mess)
	}
	if mess, diff := diff(g1.Edge(0, 0), true); diff {
		t.Errorf("g1.Edge(0, 0) %s", mess)
	}
	if mess, diff := diff(g5.Edge(0, 1), true); diff {
		t.Errorf("g5.Edge(0, 1) %s", mess)
	}
	if mess, diff := diff(g5.Edge(1, 0), false); diff {
		t.Errorf("g5.Edge(1, 0) %s", mess)
	}
	if mess, diff := diff(g5.Edge(2, 3), true); diff {
		t.Errorf("g5.Edge(2, 3) %s", mess)
	}
	if mess, diff := diff(g5.Edge(3, 2), false); diff {
		t.Errorf("g5.Edge(3, 2) %s", mess)
	}
}

func TestDegree(t *testing.T) {
	_, g1, g5 := setup()

	if mess, diff := diff(g1.Degree(0), 1); diff {
		t.Errorf("g1.Degree(0) %s", mess)
	}
	if mess, diff := diff(g5.Degree(0), 1); diff {
		t.Errorf("g5.Degree(0) %s", mess)
	}
	if mess, diff := diff(g5.Degree(1), 0); diff {
		t.Errorf("g5.Degree(1) %s", mess)
	}
}

func TestCost(t *testing.T) {
	_, g1, g5 := setup()

	cost := g1.Cost(-1, 0)
	if mess, diff := diff(cost, int64(0)); diff {
		t.Errorf("g1.Cost(0, 0) %s", mess)
	}

	cost = g1.Cost(1, 0)
	if mess, diff := diff(cost, int64(0)); diff {
		t.Errorf("g1.Cost(0, 0) %s", mess)
	}

	cost = g1.Cost(0, 0)
	if mess, diff := diff(cost, int64(0)); diff {
		t.Errorf("g1.Cost(0, 0) %s", mess)
	}

	cost = g5.Cost(0, 1)
	if mess, diff := diff(cost, int64(0)); diff {
		t.Errorf("g5.Cost(0, 1) %s", mess)
	}

	cost = g5.Cost(2, 3)
	if mess, diff := diff(cost, int64(1)); diff {
		t.Errorf("g5.Cost(2, 3) %s", mess)
	}

	cost = g5.Cost(3, 2)
	if mess, diff := diff(cost, int64(0)); diff {
		t.Errorf("g5.Cost(3, 2) %s", mess)
	}

	cost = g5.Cost(1, 2)
	if mess, diff := diff(cost, int64(0)); diff {
		t.Errorf("g5.Cost(1, 2) %s", mess)
	}

	for i := 0; i < 2; i++ {
		g1.Delete(0, 0)
		g5.Delete(1, 0)
		g5.Delete(2, 3)

		cost := g1.Cost(0, 0)
		if mess, diff := diff(cost, int64(0)); diff {
			t.Errorf("g1.Cost(0, 0) %s", mess)
		}

		cost = g5.Cost(0, 1)
		if mess, diff := diff(cost, int64(0)); diff {
			t.Errorf("g5.Cost(0, 1) %s", mess)
		}

		cost = g5.Cost(2, 3)
		if mess, diff := diff(cost, int64(0)); diff {
			t.Errorf("g5.Cost(2, 3) %s", mess)
		}

		cost = g5.Cost(3, 2)
		if mess, diff := diff(cost, int64(0)); diff {
			t.Errorf("g5.Cost(3, 2) %s", mess)
		}

		cost = g5.Cost(1, 2)
		if mess, diff := diff(cost, int64(0)); diff {
			t.Errorf("diff.Cost(1, 2) %s", mess)
		}
	}
}

func TestAddNewCost(t *testing.T) {
	_, g1, g5 := setup()

	g1.AddCost(0, 0, 8)
	cost := g1.Cost(0, 0)
	if mess, diff := diff(cost, int64(8)); diff {
		t.Errorf("g1.Cost(0, 0) %s", mess)
	}

	g5.Add(2, 3)
	cost = g5.Cost(2, 3)
	if mess, diff := diff(cost, int64(0)); diff {
		t.Errorf("g5.Cost(2, 3) %s", mess)
	}

	g5.Add(3, 2)
	cost = g5.Cost(3, 2)
	if mess, diff := diff(cost, int64(0)); diff {
		t.Errorf("g5.Cost(3, 2) %s", mess)
	}
}

func TestAddBothNewCost(t *testing.T) {
	_, g1, g5 := setup()

	g1.AddBothCost(0, 0, 8)
	cost := g1.Cost(0, 0)
	if mess, diff := diff(cost, int64(8)); diff {
		t.Errorf("g1.Cost(0, 0, 8) %s", mess)
	}

	g5.AddBoth(2, 3)
	cost = g5.Cost(2, 3)
	if mess, diff := diff(cost, int64(0)); diff {
		t.Errorf("g5.Cost(2, 3) %s", mess)
	}

	g5.AddBoth(3, 2)
	cost = g5.Cost(3, 2)
	if mess, diff := diff(cost, int64(0)); diff {
		t.Errorf("g5.Cost(3, 2) %s", mess)
	}
}
