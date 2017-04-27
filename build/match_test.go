package build

import "testing"

func TestMatch(t *testing.T) {
	g := Grid(2, 2)
	b := EdgeSet{
		Cost: func(v, w int) int64 { return int64(10*v + w) },
	}
	res := g.Match(Empty(2), b)
	exp := "6 [{0 1} {0 2} (0 4):4 {1 3} (1 5):15 {2 3} (4 0):40 (5 1):51]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Match %s", mess)
	}
	Consistent("Match1", t, res)

	c5 := Cycle(5)
	p5 := Cycle(5).Complement()
	Petersen := c5.Match(p5, AllEdges())
	exp = "10 [{0 1} {0 4} {0 5} {1 2} {1 6} {2 3} {2 7} {3 4} {3 8} {4 9} {5 7} {5 8} {6 8} {6 9} {7 9}]"
	if mess, diff := diff(Petersen.String(), exp); diff {
		t.Errorf("Match %s", mess)
	}
	Consistent("Match2", t, res)

	b = EdgeSet{
		From: Vertex(1),
		To:   Vertex(3),
		Keep: func(v, w int) bool { return v < w },
		Cost: Cost(8),
	}
	cost := func(v, w int) int64 { return int64(10*v + w) }
	g = Grid(1, 2).AddCostFunc(cost)
	res = g.Match(g, b)
	exp = "4 [(0 1):1 (1 0):10 (1 3):8 (2 3):1 (3 2):10]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Match %s", mess)
	}
	Consistent("Match3", t, res)

	b = EdgeSet{
		From: Range(1, 3).Or(Range(4, 6)),
		To:   Range(5, 7).Or(Range(9, 11)),
		Keep: func(v, w int) bool { return v < w },
		Cost: cost,
	}
	g = Empty(6)
	res = g.Match(g, b)
	exp = "12 [(1 6):16 (2 9):29 (4 10):50]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Match %s", mess)
	}
	Consistent("Match4", t, res)

	for m := 0; m < 5; m++ {
		for n := 0; n < 5; n++ {
			b := EdgeSet{Range(1, m), Range(m-n, m+n), nil, nil}
			res = Kn(m).AddCostFunc(cost).Match(Grid(m, m+n).AddCostFunc(cost), b)
			Consistent("Match", t, res)
			res = Grid(1, m).AddCostFunc(cost).Match(Kn(m).AddCostFunc(cost), b)
			Consistent("Match", t, res)
		}
	}
}
