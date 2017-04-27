package build

import "testing"

func TestTensor(t *testing.T) {
	res := Empty(4).Tensor(Grid(1, 2))
	exp := "8 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Tensor %s", mess)
	}
	Consistent("Tensor", t, res)

	res = Kn(1).Tensor(Grid(1, 2))
	exp = "2 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Tensor %s", mess)
	}
	Consistent("Tensor", t, res)

	exp = "4 [{0 3} {1 2}]"
	res = Grid(1, 2).Tensor(Grid(1, 2))
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Tensor %s", mess)
	}
	Consistent("Tensor", t, res)

	exp = "6 [{0 3} {1 2} {2 5} {3 4}]"
	res = Grid(1, 3).Tensor(Grid(1, 2))
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Tensor %s", mess)
	}
	Consistent("Tensor", t, res)

	for m := 0; m < 4; m++ {
		for n := 0; n < 4; n++ {
			Consistent("Tensor", t, Kn(m).Tensor(Kn(n)))
			Consistent("Tensor", t, Kn(m).Tensor(Grid(m, m+n)))
			Consistent("Tensor", t, Grid(m, m+n).Tensor(Kn(m)))
		}
	}
}
