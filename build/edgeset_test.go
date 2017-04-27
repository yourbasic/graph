package build

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	res := Empty(0).Add(AllEdges())
	exp := "0 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Add %s", mess)
	}
	Consistent("Add", t, res)

	res = Grid(2, 2).Add(EdgeSet{
		Cost: Cost(4),
	})
	exp = "4 [{0 1} {0 2} {0 3}:4 {1 2}:4 {1 3} {2 3}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Add %s", mess)
	}
	Consistent("Add", t, res)

	for m := 0; m < 5; m++ {
		for n := 0; n < 5; n++ {
			res = Kn(m).Add(AllEdges())
			mess := fmt.Sprintf("Kn(%d).Add(AllEdges())", m)
			Consistent(mess, t, res)

			res = Kn(m).Add(NoEdges())
			mess = fmt.Sprintf("Kn(%d).Add(NoEdges())", m)
			Consistent(mess, t, res)

			res = Grid(m, n).Add(EdgeSet{
				From: Vertex(m),
				To:   Vertex(n),
				Cost: Cost(3),
			})
			mess = fmt.Sprintf("Grid(%d,%d).Add(Edge(%d, %d, 3))", m, n, m, n)
			Consistent(mess, t, res)
		}
	}
}

func TestDelete(t *testing.T) {
	res := Kn(4).Delete(EdgeSet{From: Vertex(1)})
	exp := "4 [{0 2} {0 3} {2 3}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Delete %s", mess)
	}
	Consistent("Delete1", t, res)

	res = Kn(0).Delete(EdgeSet{From: Vertex(1)})
	exp = "0 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Delete %s", mess)
	}
	Consistent("Delete2", t, res)

	res = Kn(4).AddCost(8).Delete(EdgeSet{From: Vertex(1)})
	exp = "4 [{0 2}:8 {0 3}:8 {2 3}:8]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Delete %s", mess)
	}
	Consistent("Delete3", t, res)

	res = Kn(4).Delete(AllEdges())
	exp = "4 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Delete %s", mess)
	}
	Consistent("Delete4", t, res)

	res = Grid(2, 2).Delete(Edge(2, 3))
	exp = "4 [{0 1} {0 2} {1 3}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Delete %s", mess)
	}
	Consistent("Delete5", t, res)

	for m := 0; m < 5; m++ {
		for n := 0; n < 5; n++ {
			res = Kn(m).Delete(AllEdges())
			mess := fmt.Sprintf("Kn(%d).Delete(AllEdges())", m)
			Consistent(mess, t, res)

			res = Kn(m).Delete(NoEdges())
			mess = fmt.Sprintf("Kn(%d).Delete(NoEdges())", m)
			Consistent(mess, t, res)

			res = Grid(m, n).Delete(EdgeSet{
				From: Vertex(m),
				To:   Vertex(n),
				Cost: Cost(3),
			})
			mess = fmt.Sprintf("Grid(%d,%d).Delete(Edge(%d, %d, 3))", m, n, m, n)
			Consistent(mess, t, res)
		}
	}
}

func TestNewEdges(t *testing.T) {
	res := newEdges(0, AllEdges())
	exp := "0 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("newEdges %s", mess)
	}
	Consistent("newEdges1", t, res)

	res = newEdges(3, Edge(1, 2))
	exp = "3 [{1 2}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("newEdges %s", mess)
	}
	Consistent("newEdges2", t, res)

	res = newEdges(3, Edge(1, 1))
	exp = "3 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("newEdges %s", mess)
	}
	Consistent("newEdges3", t, res)

	res = newEdges(4, EdgeSet{
		From: Vertex(1),
		To:   Vertex(2),
		Cost: Cost(3),
	})
	exp = "4 [{1 2}:3]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("newEdges %s", mess)
	}
	Consistent("newEdges4", t, res)

	res = newEdges(3, EdgeSet{})
	exp = "3 [{0 1} {0 2} {1 2}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("newEdges %s", mess)
	}
	Consistent("newEdges5", t, res)

	res = newEdges(4, EdgeSet{
		From: Range(0, 3),
		To:   Range(1, 4),
		Keep: func(v, w int) bool { return v <= w },
	})
	exp = "4 [(0 1) (0 2) (0 3) (1 2) (1 3) (2 3)]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("newEdges %s", mess)
	}
	Consistent("newEdges6", t, res)

	for m := 0; m < 8; m++ {
		for n := 0; n < 8; n++ {
			res = newEdges(m+n, EdgeSet{
				From: Range(m-n, m+n-3),
				To:   Range(n-m+4, m+n+2),
				Keep: func(v, w int) bool { return true },
				Cost: func(v, w int) int64 { return int64(10*v + w) },
			})
			mess := "newEdges(" + strconv.Itoa(m) + "," + strconv.Itoa(n) + ")"
			Consistent(mess, t, res)
		}
	}
}
