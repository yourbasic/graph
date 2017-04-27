package build

import "testing"

func TestCartesian(t *testing.T) {
	res := Empty(2).Cartesian(Grid(1, 2))
	exp := "4 [{0 1} {2 3}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Cartesian %s", mess)
	}
	Consistent("Cartesian1", t, res)

	res = Kn(1).Cartesian(Grid(1, 2))
	exp = "2 [{0 1}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Cartesian %s", mess)
	}
	Consistent("Cartesian2", t, res)

	res = Grid(1, 2).Cartesian(Grid(1, 2))
	exp = "4 [{0 1} {0 2} {1 3} {2 3}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Cartesian %s", mess)
	}
	Consistent("Cartesian3", t, res)

	res = Grid(1, 3).Cartesian(Grid(1, 3))
	exp = "9 [{0 1} {0 3} {1 2} {1 4} {2 5} {3 4} {3 6} {4 5} {4 7} {5 8} {6 7} {7 8}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Cartesian %s", mess)
	}
	Consistent("Cartesian4", t, res)

	res = Kn(2).Cartesian(Grid(1, 2))
	exp = "4 [{0 1} {0 2} {1 3} {2 3}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Cartesian %s", mess)
	}
	Consistent("Cartesian5", t, res)

	for m := 0; m < 4; m++ {
		for n := 0; n < 4; n++ {
			Consistent("Cartesian6", t, Kn(m).Cartesian(Cycle(n)))
			Consistent("Cartesian7", t, Kn(m).Cartesian(Grid(m, m+n)))
			Consistent("Cartesian8", t, Grid(m, m+n).Cartesian(Kn(m)))
		}
	}
}
