package build

import (
	"fmt"
	"github.com/yourbasic/graph"
	"reflect"
	"testing"
)

func diff(res, exp interface{}) (message string, diff bool) {
	if !reflect.DeepEqual(res, exp) {
		message = fmt.Sprintf("%v; want %v", res, exp)
		diff = true
	}
	return
}

// Check internal consistency.
func Consistent(mess string, t *testing.T, g *Virtual) {
	if g == nil {
		return
	}
	n := g.order
	// Check that degree and visit are consistent with edge and cost.
	for v := 0; v < n; v++ {
		deg := 0
		visited := make([]bool, n)
		g.visit(v, 0, func(w int, c int64) (skip bool) {
			visited[w] = true
			deg++
			if !g.edge(v, w) {
				t.Errorf("%s: visit(%d): (%d,%d); edge(%d,%d): false\n", mess, v, w, c, v, w)
			}
			if cost := g.cost(v, w); c != cost {
				t.Errorf("%s: visit(%d): (%d,%d); cost(%d,%d): %v\n", mess, v, w, c, v, w, cost)
			}
			return
		})
		d := g.degree(v)
		if deg != d {
			t.Errorf("%s: visit(%d) found %v neighbors; degree(%d): %v\n", mess, v, deg, v, d)
		}
		for w, ok := range visited {
			if ok {
				continue
			}
			if g.edge(v, w) {
				t.Errorf("%s: visit(%d,%d): --; edge(%d,%d): true\n", mess, v, w, v, w)
			}
		}
		CheckStart(mess, t, g, v, visited)
		CheckAbort(mess, t, g, v)
	}
	return
}

// Check visit from all different starting points.
func CheckStart(mess string, t *testing.T, g *Virtual, v int, visited []bool) {
	n := g.order
	check := make([]bool, len(visited)) // Check that we get the same neighbors.
	for a := 0; a <= n+1; a++ {
		prev := -1
		for v := a; v < n; v++ {
			check[v] = visited[v]
		}
		g.visit(v, a, func(w int, _ int64) (skip bool) {
			if w < 0 {
				t.Errorf("%s: visit(%d, %d): returned negative w=%d", mess, v, a, w)
			}
			if w >= n {
				t.Errorf("%s: visit(%d, %d): returned w=%d bigger than n=%d", mess, v, a, w, n)
			}
			if w < a {
				t.Errorf("%s: visit(%d, %d): returned small w=%d", mess, v, a, w)
			} else if 0 <= w && w < n && !check[w] {
				t.Errorf("%s, visit(%d, %d) didn't return %d when a=0.\n", mess, v, a, w)
			}
			if v == w {
				t.Errorf("%s: visit(%d, %d): self-loop detected\n", mess, v, a)
			}
			if prev > w {
				t.Errorf("%s, visit(%d, %d) returned %d before %d.\n", mess, v, a, prev, w)
			}
			check[w] = false
			prev = w
			return
		})
		for w := a; w < n; w++ {
			if check[w] {
				t.Errorf("%s, visit(%d, %d) did return %d when a=0.\n", mess, v, a, w)
			}
		}
	}
}

// Check visit for failed abort and wrong return value.
func CheckAbort(mess string, t *testing.T, g *Virtual, v int) {
	for n := 1; n <= g.degree(v); n++ {
		count := 0
		aborted := g.visit(v, 0, func(_ int, _ int64) bool {
			count++
			return count == n // break after n iterations
		})
		if count != n {
			t.Errorf("%s: visit(%d) didnt't abort after %d iteration(s).\n", mess, v, n)
		}
		if !aborted {
			t.Errorf("%s: visit(%d) didnt't return true after %d iteration(s).\n", mess, v, n)
		}
		aborted = g.visit(v, 0, func(_ int, _ int64) bool {
			return false
		})
		if aborted {
			t.Errorf("%s: visit(%d) returned true after completed iteration.\n", mess, v)
		}

	}
}

func TestNull(t *testing.T) {
	g := null
	if mess, diff := diff(g.String(), "0 []"); diff {
		t.Errorf("null %s", mess)
	}
	if mess, diff := diff(g.Order(), 0); diff {
		t.Errorf("null %s", mess)
	}
	Consistent("null", t, g)
}

func TestSingleton(t *testing.T) {
	if mess, diff := diff(singleton().String(), "1 []"); diff {
		t.Errorf("singleton %s", mess)
	}
	Consistent("singleton", t, singleton())

	if mess, diff := diff(singleton().String(), "1 []"); diff {
		t.Errorf("singleton %s", mess)
	}
	Consistent("singleton", t, singleton())
}

func TestEdge(t *testing.T) {
	e := edge()
	if mess, diff := diff(e.String(), "2 [{0 1}]"); diff {
		t.Errorf("edge %s", mess)
	}
	Consistent("edge", t, e)
}

func TestLine(t *testing.T) {
	if mess, diff := diff(line(-1), (*Virtual)(nil)); diff {
		t.Errorf("line %s", mess)
	}
	if mess, diff := diff(line(0).String(), "0 []"); diff {
		t.Errorf("line %s", mess)
	}
	if mess, diff := diff(line(1).String(), "1 []"); diff {
		t.Errorf("line %s", mess)
	}
	if mess, diff := diff(line(2).String(), "2 [{0 1}]"); diff {
		t.Errorf("line %s", mess)
	}
	if mess, diff := diff(line(3).String(), "3 [{0 1} {1 2}]"); diff {
		t.Errorf("line %s", mess)
	}
	for n := 0; n < 5; n++ {
		Consistent("line", t, line(n))
	}
}

func TestGeneric(t *testing.T) {
	g := generic0(-1, alwaysEdge)
	if mess, diff := diff(g, (*Virtual)(nil)); diff {
		t.Errorf("generic0 %s", mess)
	}
	Consistent("generic0", t, g)

	g = generic0(0, alwaysEdge)
	if mess, diff := diff(g.String(), "0 []"); diff {
		t.Errorf("generic0 %s", mess)
	}
	Consistent("generic0", t, g)

	g = generic0(1, alwaysEdge)
	if mess, diff := diff(g.String(), "1 []"); diff {
		t.Errorf("generic0 %s", mess)
	}
	Consistent("generic0", t, g)

	g = generic(-1, zero, alwaysEdge)
	if mess, diff := diff(g, (*Virtual)(nil)); diff {
		t.Errorf("generic %s", mess)
	}
	Consistent("generic", t, g)

	g = generic(0, zero, alwaysEdge)
	if mess, diff := diff(g.String(), "0 []"); diff {
		t.Errorf("generic %s", mess)
	}
	Consistent("generic", t, g)

	g = generic(1, zero, alwaysEdge)
	if mess, diff := diff(g.String(), "1 []"); diff {
		t.Errorf("generic %s", mess)
	}
	Consistent("generic", t, g)

	g = generic(2, zero, alwaysEdge)
	if mess, diff := diff(g.String(), "2 [{0 1}]"); diff {
		t.Errorf("generic %s", mess)
	}
	Consistent("generic", t, g)

	g = Generic(-1, nil)
	if mess, diff := diff(g, (*Virtual)(nil)); diff {
		t.Errorf("Generic %s", mess)
	}
	Consistent("Generic", t, g)

	g = Generic(0, nil)
	if mess, diff := diff(g.String(), "0 []"); diff {
		t.Errorf("Generic %s", mess)
	}
	Consistent("Generic", t, g)

	g = Generic(1, nil)
	if mess, diff := diff(g.String(), "1 []"); diff {
		t.Errorf("Generic %s", mess)
	}
	Consistent("Generic", t, g)

	g = Generic(2, nil)
	if mess, diff := diff(g.String(), "2 [{0 1}]"); diff {
		t.Errorf("Generic %s", mess)
	}
	Consistent("Generic", t, g)

	g = Generic(0, func(v, w int) bool { return v != 0 })
	if mess, diff := diff(g.String(), "0 []"); diff {
		t.Errorf("Generic %s", mess)
	}
	Consistent("Generic", t, g)

	g = Generic(1, func(v, w int) bool { return v != 0 })
	if mess, diff := diff(g.String(), "1 []"); diff {
		t.Errorf("Generic %s", mess)
	}
	Consistent("Generic", t, g)

	g = Generic(2, func(v, w int) bool { return v != 0 })
	if mess, diff := diff(g.String(), "2 [(1 0)]"); diff {
		t.Errorf("Generic %s", mess)
	}
	Consistent("Generic", t, g)

	g3 := Generic(3, func(v, w int) bool { return v != 0 })
	if mess, diff := diff(g3.String(), "3 [(1 0) {1 2} (2 0)]"); diff {
		t.Errorf("Generic %s", mess)
	}
	Consistent("Generic", t, g3)
}

func TestSpecific(t *testing.T) {
	if mess, diff := diff(graph.Equal(Kn(4), Specific(Kn(4))), true); diff {
		t.Errorf("Specific %s", mess)
	}
	Consistent("Specific", t, Specific(Kn(4)))

	cost := func(v, w int) int64 { return int64(10*v + w) }
	res := Specific(Grid(2, 2).AddCostFunc(cost))
	exp := Grid(2, 2).AddCostFunc(cost)
	if mess, diff := diff(graph.Equal(res, exp), true); diff {
		t.Errorf("Specific %s", mess)
	}
	Consistent("Specific", t, Specific(Kn(4)))

	if mess, diff := diff(res.cost(0, 0), int64(0)); diff {
		t.Errorf("Specific cost %s", mess)
	}
	if mess, diff := diff(res.cost(0, 1), int64(1)); diff {
		t.Errorf("Specific cost %s", mess)
	}
	if mess, diff := diff(res.cost(1, 0), int64(10)); diff {
		t.Errorf("Specific cost %s", mess)
	}
}

func TestEmpty(t *testing.T) {
	if mess, diff := diff(Empty(-1), (*Virtual)(nil)); diff {
		t.Errorf("Empty %s", mess)
	}
	if mess, diff := diff(Empty(0).String(), "0 []"); diff {
		t.Errorf("Empty %s", mess)
	}
	if mess, diff := diff(Empty(1).String(), "1 []"); diff {
		t.Errorf("Empty %s", mess)
	}
	if mess, diff := diff(Empty(2).String(), "2 []"); diff {
		t.Errorf("Empty %s", mess)
	}
	for n := 0; n < 10; n++ {
		Consistent("Empty", t, Empty(n))
	}
}

func TestKn(t *testing.T) {
	if mess, diff := diff(Kn(-1), (*Virtual)(nil)); diff {
		t.Errorf("Kn %s", mess)
	}
	if mess, diff := diff(Kn(0).String(), "0 []"); diff {
		t.Errorf("Kn %s", mess)
	}
	if mess, diff := diff(Kn(1).String(), "1 []"); diff {
		t.Errorf("Kn %s", mess)
	}
	if mess, diff := diff(Kn(2).String(), "2 [{0 1}]"); diff {
		t.Errorf("Kn %s", mess)
	}
	for n := 0; n < 5; n++ {
		Consistent("Kn", t, Kn(n))
	}
}

func TestComplement(t *testing.T) {
	res := Kn(0).AddCost(1).Complement()
	exp := "0 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Complement %s", mess)
	}
	Consistent("Complement", t, res)

	res = Kn(4).AddCost(1).Complement()
	exp = "4 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Complement %s", mess)
	}
	Consistent("Complement", t, res)

	res = Empty(1).Complement()
	exp = "1 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Complement %s", mess)
	}
	Consistent("Complement", t, res)

	res = Cycle(2).AddCost(1).Complement()
	exp = "2 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Complement %s", mess)
	}
	Consistent("Complement", t, res)

	res = Cycle(4).AddCost(1).Complement()
	exp = "4 [{0 2} {1 3}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Complement %s", mess)
	}
	Consistent("Complement", t, res)

	res = Grid(2, 3).Complement()
	exp = "6 [{0 2} {0 4} {0 5} {1 3} {1 5} {2 3} {2 4} {3 5}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("Complement %s", mess)
	}
	Consistent("Complement", t, res)
}

func TestKeep(t *testing.T) {
	g := Cycle(3).AddCost(1).Keep(func(v, w int) bool {
		return v+1 != w
	})
	if mess, diff := diff(g.String(), "3 [{0 2}:1 (1 0):1 (2 1):1]"); diff {
		t.Errorf("Keep %s", mess)
	}
	Consistent("Keep", t, g)

	g = Cycle(3).AddCost(1).Keep(nil)
	if mess, diff := diff(g.String(), "3 [{0 1}:1 {0 2}:1 {1 2}:1]"); diff {
		t.Errorf("Keep %s", mess)
	}
	Consistent("Keep", t, g)
}

func TestAddCost(t *testing.T) {
	res := Grid(1, 2).AddCost(5)
	exp := "2 [{0 1}:5]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("AddCost %s", mess)
	}
	Consistent("AccCost", t, res)

	res = Empty(0).AddCost(0)
	exp = "0 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("AddCost %s", mess)
	}
	Consistent("AccCost", t, res)
}

func TestAddCostFunc(t *testing.T) {
	res := Grid(1, 2).AddCostFunc(Cost(5))
	exp := "2 [{0 1}:5]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("AddCostFunc %s", mess)
	}
	Consistent("AccCostFunc", t, res)

	if mess, diff := diff(res.Cost(0, 0), int64(0)); diff {
		t.Errorf("Cost %s", mess)
	}
	if mess, diff := diff(res.Cost(0, 1), int64(5)); diff {
		t.Errorf("Cost %s", mess)
	}
	if mess, diff := diff(res.Cost(-1, 0), int64(0)); diff {
		t.Errorf("Cost %s", mess)
	}

	res = Grid(1, 2).AddCostFunc(nil)
	exp = "2 [{0 1}]"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("AddCostFunc %s", mess)
	}
	Consistent("AccCostFunc", t, res)

	res = Empty(0).AddCostFunc(Cost(5))
	exp = "0 []"
	if mess, diff := diff(res.String(), exp); diff {
		t.Errorf("AddCostFunc %s", mess)
	}
	Consistent("AddCostFunc", t, res)
}

func TestVisit(t *testing.T) {
	g := Grid(1, 2).AddCost(5)
	res := g.Visit(0, func(w int, c int64) (skip bool) {
		return w == 1 && c == 5
	})
	if mess, diff := diff(res, true); diff {
		t.Errorf("Visit %s", mess)
	}

	res = g.VisitFrom(0, 0, func(w int, c int64) (skip bool) {
		return w == 1 && c == 5
	})
	if mess, diff := diff(res, true); diff {
		t.Errorf("Visit %s", mess)
	}

	res = g.VisitFrom(0, -1, func(w int, c int64) (skip bool) {
		return w == 1 && c == 5
	})
	if mess, diff := diff(res, true); diff {
		t.Errorf("Visit %s", mess)
	}

	res = g.VisitFrom(0, 2, func(w int, c int64) (skip bool) {
		return w == 1 && c == 5
	})
	if mess, diff := diff(res, false); diff {
		t.Errorf("Visit %s", mess)
	}
}
