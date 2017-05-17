package graph

import (
	"math/rand"
	"testing"
)

func SetUpImm() (g0, g1, g1c, g5, g5c *Immutable) {
	g0 = Sort(New(0))

	g := New(1)
	g.Add(0, 0)
	g1 = Sort(g)

	g.AddCost(0, 0, 1)
	g1c = Sort(g)

	g = New(5)
	g.Add(0, 1)
	g.Add(2, 3)
	g5 = Sort(g)

	g.AddCost(2, 3, 1)
	g5c = Sort(g)
	return
}

func TestSort(t *testing.T) {
	g0, g1, g1c, g5, g5c := SetUpImm()

	res := g0
	exp := "0 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Sort: %s", mess)
	}
	Consistent("Sort g0", t, res)

	res = g1
	exp = "1 [(0 0)]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Sort: %s", mess)
	}
	Consistent("Sort g1", t, res)

	res = g1c
	exp = "1 [(0 0):1]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Sort: %s", mess)
	}
	Consistent("Sort g1c", t, res)

	res = g5
	exp = "5 [(0 1) (2 3)]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Sort: %s", mess)
	}
	Consistent("Sort g5", t, res)

	res = g5c
	exp = "5 [(0 1) (2 3):1]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Sort: %s", mess)
	}
	Consistent("Sort g5c", t, res)

	res = Sort(g5c)
	exp = "5 [(0 1) (2 3):1]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Sort: %s", mess)
	}
	Consistent("Sort Sort(g5c)", t, res)

	n := 10
	g := New(n)
	for i := 0; i < 2*n; i++ {
		g.AddBothCost(rand.Intn(n), rand.Intn(n), rand.Int63())
	}
	Consistent("Sort rand", t, Sort(g))
	Consistent("Sort Sort(rand)", t, Sort(Sort(g)))
}

func TestTranspose(t *testing.T) {
	g0, g1, g1c, g5, g5c := SetUpImm()

	res := Transpose(g0)
	exp := "0 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Transpose: %s", mess)
	}
	Consistent("Transpose g0", t, res)

	res = Transpose(g1)
	exp = "1 [(0 0)]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Transpose g1: %s", mess)
	}
	Consistent("Transpose", t, res)

	res = Transpose(g1c)
	exp = "1 [(0 0):1]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Transpose g1c: %s", mess)
	}
	Consistent("Transpose", t, res)

	res = Transpose(g5)
	exp = "5 [(1 0) (3 2)]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Transpose: %s", mess)
	}
	Consistent("Transpose g5", t, res)

	res = Transpose(g5c)
	exp = "5 [(1 0) (3 2):1]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Transpose: %s", mess)
	}
	Consistent("Transpose g5c", t, res)

	n := 10
	g := New(n)
	for i := 0; i < 2*n; i++ {
		g.AddBothCost(rand.Intn(n), rand.Intn(n), rand.Int63())
	}
	Consistent("Transpose rand", t, Transpose(g))
}

func TestOrderImm(t *testing.T) {
	g0, g1, g1c, g5, g5c := SetUpImm()
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
	if mess, diff := diff(g1c.Order(), 1); diff {
		t.Errorf("g1.%s %s", s, mess)
	}
	if mess, diff := diff(g5c.Order(), 5); diff {
		t.Errorf("g5.%s %s", s, mess)
	}
}

func TestEdgeImm(t *testing.T) {
	_, g1, g1c, g5, _ := SetUpImm()

	if mess, diff := diff(g1.Edge(-1, 0), false); diff {
		t.Errorf("g1.Edge(-1, 0) %s", mess)
	}
	if mess, diff := diff(g1.Edge(1, 0), false); diff {
		t.Errorf("g1.Edge(1, 0) %s", mess)
	}
	if mess, diff := diff(g1.Edge(0, 0), true); diff {
		t.Errorf("g1.Edge(0, 0) %s", mess)
	}
	if mess, diff := diff(g1c.Edge(0, 0), true); diff {
		t.Errorf("g1c.Edge(0, 0) %s", mess)
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

func TestDegreeImm(t *testing.T) {
	_, g1, g1c, g5, g5c := SetUpImm()

	if mess, diff := diff(g1.Degree(0), 1); diff {
		t.Errorf("g1.Degree(0) %s", mess)
	}
	if mess, diff := diff(g1c.Degree(0), 1); diff {
		t.Errorf("g1c.Degree(0) %s", mess)
	}
	if mess, diff := diff(g5.Degree(0), 1); diff {
		t.Errorf("g5.Degree(0) %s", mess)
	}
	if mess, diff := diff(g5c.Degree(0), 1); diff {
		t.Errorf("g5c.Degree(0) %s", mess)
	}
	if mess, diff := diff(g5.Degree(1), 0); diff {
		t.Errorf("g5.Degree(1) %s", mess)
	}
	if mess, diff := diff(g5c.Degree(1), 0); diff {
		t.Errorf("g5c.Degree(1) %s", mess)
	}
}
