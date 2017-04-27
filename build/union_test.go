package build

import (
	"fmt"
	"testing"
)

func TestUnion(t *testing.T) {
	res := Empty(0).Union(Empty(0))
	exp := "0 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Union %s", mess)
	}
	Consistent("Union", t, res)

	res = Empty(0).AddCost(4).Union(Grid(1, 2).AddCost(5))
	exp = "2 [{0 1}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Union %s", mess)
	}
	Consistent("Union", t, res)

	res = Grid(1, 2).AddCost(4).Union(Grid(1, 2).AddCost(5))
	exp = "2 [{0 1}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Union %s", mess)
	}
	Consistent("Union", t, res)

	res = Grid(2, 2).Union(Kn(3).AddCost(8))
	exp = "4 [{0 1} {0 2} {1 2} {1 3} {2 3}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Union %s", mess)
	}
	Consistent("Union", t, res)

	res = Kn(3).Union(Grid(2, 2))
	exp = "4 [{0 1} {0 2} {1 2} {1 3} {2 3}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Union %s", mess)
	}
	Consistent("Union", t, res)

	for m := 0; m < 5; m++ {
		for n := 0; n < 5; n++ {
			res = Kn(m).Union(Kn(n))
			mess := fmt.Sprintf("Kn(%d).Union(Kn(%d))", m, n)
			Consistent(mess, t, res)

			res = Grid(m, n).Union(Kn(n))
			mess = fmt.Sprintf("Grid(%d,%d).Union(Kn(%d))", m, n, n)
			Consistent(mess, t, res)
			res = Kn(n).Union(Grid(m, n))
			mess = fmt.Sprintf("Kn(%d).Union(Grid(%d,%d))", n, m, n)
			Consistent(mess, t, res)

			km := Kn(m).Keep(func(v, w int) bool {
				return v != m-1 && w != m-1
			})
			res = km.Union(Kn(n))
			mess = fmt.Sprintf("Kn(%d)-{m-1}.Union(Kn(%d))", m, n)
			Consistent(mess, t, res)
		}
	}
}
