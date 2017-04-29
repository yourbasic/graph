package build

import "testing"

func TestSubgraph(t *testing.T) {
	cost := func(v, w int) int64 { return int64(10*v + w) }

	res := Kn(6).Subgraph(Range(3, 6))
	exp := "3 [{0 1} {0 2} {1 2}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Subgraph %s", mess)
	}
	Consistent("Subgraph1", t, res)

	res = Kn(2).Subgraph(Range(-1, 10))
	exp = "2 [{0 1}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Subgraph %s", mess)
	}
	Consistent("Subgraph2", t, res)

	res = Kn(6).AddCostFunc(cost).Subgraph(Range(3, 6))
	exp = "3 [(0 1):34 (0 2):35 (1 0):43 (1 2):45 (2 0):53 (2 1):54]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Subgraph %s", mess)
	}
	Consistent("Subgraph3", t, res)

	res = Kn(6).AddCostFunc(cost).Subgraph(Vertex(1).Or(Vertex(3)))
	exp = "2 [(0 1):13 (1 0):31]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Subgraph %s", mess)
	}
	Consistent("Subgraph4", t, res)

	res = Grid(3, 3).Subgraph(Range(0, 4).Or(Range(5, 9)))
	exp = "8 [{0 1} {0 3} {1 2} {2 4} {3 5} {4 7} {5 6} {6 7}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Subgraph %s", mess)
	}
	Consistent("Subgraph5", t, res)

	for m := 0; m < 6; m++ {
		for n := 0; n < 6; n++ {
			set := Range(1, m).Or(Range(2, n)).Or(Range(3, m+n))
			res = Kn(m + n).AddCostFunc(cost).Subgraph(set)
			Consistent("Subgraph", t, res)
			res = Grid(m, n).AddCostFunc(cost).Subgraph(set)
			Consistent("Subgraph", t, res)
		}
	}
}
