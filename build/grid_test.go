package build

import "testing"

func TestGrid(t *testing.T) {
	if mess, diff := diff(Grid(0, 0).String(), "0 []"); diff {
		t.Errorf("Grid %s", mess)
	}

	if mess, diff := diff(Grid(0, 1).String(), "0 []"); diff {
		t.Errorf("Grid %s", mess)
	}

	if mess, diff := diff(Grid(1, 0).String(), "0 []"); diff {
		t.Errorf("Grid %s", mess)
	}

	if mess, diff := diff(Grid(1, 1).String(), "1 []"); diff {
		t.Errorf("Grid %s", mess)
	}

	exp := "2 [{0 1}]"
	if mess, diff := diff(Grid(1, 2).String(), exp); diff {
		t.Errorf("Grid %s", mess)
	}

	exp = "2 [{0 1}]"
	if mess, diff := diff(Grid(2, 1).String(), exp); diff {
		t.Errorf("Grid %s", mess)
	}

	exp = "4 [{0 1} {0 2} {1 3} {2 3}]"
	if mess, diff := diff(Grid(2, 2).String(), exp); diff {
		t.Errorf("Grid %s", mess)
	}

	for m := 0; m < 5; m++ {
		for n := 0; n < 5; n++ {
			Consistent("Grid", t, Grid(m, n))
		}
	}

	g := Grid(2, (1 << uint(bitsPerWord-3))) // maximum possible size
	if mess, diff := diff(g.Edge(0, 1), true); diff {
		t.Errorf("Grid %s", mess)
	}
}
