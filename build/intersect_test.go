package build

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	res := Empty(0).AddCost(4).Intersect(Grid(1, 2).AddCost(5))
	exp := "0 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Intersect %s", mess)
	}
	Consistent("Intersect1", t, res)

	res = Grid(1, 2).AddCost(4).Intersect(Grid(1, 2).AddCost(5))
	exp = "2 [{0 1}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Intersect %s", mess)
	}
	Consistent("Intersect2", t, res)

	res = Grid(2, 2).Intersect(Kn(3).AddCost(8))
	exp = "3 [{0 1} {0 2}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Intersect %s", mess)
	}
	Consistent("Intersect3", t, res)

	res = Kn(3).Intersect(Grid(2, 2))
	exp = "3 [{0 1} {0 2}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Intersect %s", mess)
	}
	Consistent("Intersect4", t, res)

	for m := 0; m < 5; m++ {
		for n := 0; n < 5; n++ {
			res = Kn(m).Intersect(Kn(n))
			mess := fmt.Sprintf("Kn(%d).Intersect(Kn(%d))", m, n)
			Consistent(mess, t, res)

			res = Grid(m, n).Intersect(Kn(n))
			mess = fmt.Sprintf("Grid(%d,%d).Intersect(Kn(%d))", m, n, n)
			Consistent(mess, t, res)
			res = Kn(n).Intersect(Grid(m, n))
			mess = fmt.Sprintf("Kn(%d).Intersect(Grid(%d,%d))", n, m, n)
			Consistent(mess, t, res)

			km := Kn(m).Keep(func(v, w int) bool {
				return v != m-1 && w != m-1
			})
			res = km.Intersect(Kn(n))
			Consistent("Intersect", t, res)
		}
	}
}
