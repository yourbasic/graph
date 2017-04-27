package build

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	cost := func(v, w int) int64 { return int64(10*v + w) }
	res := Grid(1, 2).AddCostFunc(cost).Connect(1, Grid(1, 2).AddCostFunc(cost))
	exp := "3 [(0 1):1 (1 0):10 (1 2):1 (2 1):10]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Connect %s", mess)
	}
	Consistent("Connect1", t, res)

	res = Grid(1, 2).AddCostFunc(cost).Connect(0, Grid(1, 2).AddCostFunc(cost))
	exp = "3 [(0 1):1 (0 2):1 (1 0):10 (2 0):10]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Connect %s", mess)
	}
	Consistent("Connect2", t, res)

	res = Kn(1).Connect(0, Grid(1, 2).AddCostFunc(cost))
	exp = "2 [(0 1):1 (1 0):10]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Connect %s", mess)
	}
	Consistent("Connect3", t, res)

	res = Grid(1, 2).AddCostFunc(cost).Connect(1, Kn(1))
	exp = "2 [(0 1):1 (1 0):10]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Connect %s", mess)
	}
	Consistent("Connect4", t, res)

	res = Grid(1, 2).AddCostFunc(cost).Connect(0, Kn(1))
	exp = "2 [(0 1):1 (1 0):10]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Connect %s", mess)
	}
	Consistent("Connect5", t, res)

	res = Kn(0).Connect(0, Kn(1))
	if res != nil {
		t.Errorf("Connect should produce nil connecting a null graph.")
	}
	Consistent("Connect6", t, res)

	res = Kn(1).Connect(0, Kn(0))
	if res != nil {
		t.Errorf("Connect should produce nil connecting a null graph.")
	}
	Consistent("Connect7", t, res)

	res = Cycle(3).Connect(0, Cycle(3)).Connect(0, Cycle(3))
	exp = "7 [{0 1} {0 2} {0 3} {0 4} {0 5} {0 6} {1 2} {3 4} {5 6}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Connect %s", mess)
	}
	Consistent("Connect8", t, res)

	for m := 0; m < 5; m++ {
		for n := 0; n < 5; n++ {
			res = Kn(m).Connect(m-1, Kn(n))
			mess := fmt.Sprintf("Kn(%d).Connect(Kn(%d))", m, n)
			Consistent(mess, t, res)

			res = Grid(m, n).AddCostFunc(cost).Connect(m-1, Kn(n))
			mess = fmt.Sprintf("Grid(%d,%d).Connect(Kn(%d))", m, n, n)
			Consistent(mess, t, res)
			res = Kn(n).Connect(n-1, Grid(m, n).AddCostFunc(cost))
			mess = fmt.Sprintf("Kn(%d).Connect(Grid(%d,%d))", n, m, n)
			Consistent(mess, t, res)

			km := Kn(m).Keep(func(v, w int) bool {
				return v != m-1 && w != m-1
			}).AddCostFunc(cost)
			res = km.Connect(0, Kn(n).AddCostFunc(cost))
			Consistent(mess, t, res)
		}
	}
}
