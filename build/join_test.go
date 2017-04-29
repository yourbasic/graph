package build

import "testing"

func TestJoin(t *testing.T) {
	cost := func(v, w int) int64 { return int64(10*v + w) }
	b := EdgeSet{
		From: Range(0, 4),
		To:   Range(0, 4),
		Keep: func(v, w int) bool { return true },
		Cost: Cost(8),
	}
	g := Grid(2, 2)
	res := g.Join(Empty(0), b)
	exp := "4 [{0 1} {0 2} {1 3} {2 3}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Join %s", mess)
	}
	Consistent("Join", t, res)

	res = Empty(0).Join(g, b)
	exp = "4 [{0 1} {0 2} {1 3} {2 3}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Join %s", mess)
	}
	Consistent("Join", t, res)

	b = EdgeSet{
		From: Vertex(1),
		To:   Vertex(3),
		Keep: func(v, w int) bool { return v < w },
		Cost: Cost(8),
	}
	g = Grid(1, 2).AddCostFunc(cost)
	res = g.Join(g, b)
	exp = "4 [(0 1):1 (1 0):10 (1 3):8 (2 3):1 (3 2):10]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Join %s", mess)
	}
	Consistent("Join", t, res)

	res = Empty(0).Join(Empty(0), AllEdges())
	exp = Empty(0).String()
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Join %s", mess)
	}
	Consistent("Join", t, res)

	res = Empty(2).Join(Empty(3), AllEdges())
	exp = Kmn(2, 3).String()
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Join %s", mess)
	}
	Consistent("Join", t, res)

	directed := EdgeSet{
		Keep: func(v, w int) bool { return v < w },
	}
	res = Empty(2).Join(Empty(3), directed)
	exp = Kmn(2, 3).Keep(func(v, w int) bool { return v < w }).String()
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Join %s", mess)
	}
	Consistent("Join", t, res)

	directed.Cost = Cost(3)
	res = Kn(2).AddCost(1).Join(Kn(2).AddCost(2), directed)
	exp = "4 [{0 1}:1 (0 2):3 (0 3):3 (1 2):3 (1 3):3 {2 3}:2]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Join %s", mess)
	}
	Consistent("Join", t, res)

	full := AllEdges()
	full.Cost = Cost(3)
	res = Kn(1).AddCost(1).Join(Kn(1).AddCost(2), full)
	exp = "2 [{0 1}:3]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Join %s", mess)
	}
	Consistent("Join", t, res)

	directed.Cost = Cost(9)
	g = Kn(1).AddCost(8).Join(Grid(1, 2).AddCostFunc(cost), directed)
	exp = "3 [(0 1):9 (0 2):9 (1 2):1 (2 1):10]"
	if mess, diff := diff(g.String(), exp); diff {
		t.Errorf("Join %s", mess)
	}
	Consistent("Join", t, g)

	res = Empty(2).Join(Empty(3), EdgeSet{Range(0, 1), Range(3, 4), nil, nil})
	exp = "5 [{0 3}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Join %s", mess)
	}
	Consistent("Join", t, res)

	b = EdgeSet{Range(0, 1), Range(2, 3), nil, nil}
	res = Kn(2).AddCost(1).Join(Kn(2).AddCost(2), b)
	exp = "4 [{0 1}:1 {0 2} {2 3}:2]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Join %s", mess)
	}
	Consistent("Join", t, res)

	for m := 0; m < 4; m++ {
		for n := 0; n < 4; n++ {
			b := EdgeSet{Range(m, n), Range(m-n, m+n), nil, nil}
			res = Kn(m).AddCostFunc(cost).Join(Grid(m, m+n).AddCostFunc(cost), b)
			Consistent("Join", t, res)
			res = Grid(m, m+n).AddCostFunc(cost).Join(Kn(m).AddCostFunc(cost), b)
			Consistent("Join", t, res)
		}
	}
}
