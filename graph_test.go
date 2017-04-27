package graph

import (
	"fmt"
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
func Consistent(mess string, t *testing.T, g Iterator) {
	if g == nil {
		return
	}
	gm, mutable := g.(*Mutable)
	gi, immutable := g.(*Immutable)
	n := g.Order()
	// Check that degree and visit are consistent with edge and cost.
	for v := 0; v < n; v++ {
		deg := 0
		visited := make([]bool, n)
		g.Visit(v, func(w int, c int64) (skip bool) {
			visited[w] = true
			deg++
			if mutable && !gm.Edge(v, w) || immutable && !gi.Edge(v, w) {
				t.Errorf("%s: visit(%d): (%d,%d); edge(%d,%d): false\n", mess, v, w, c, v, w)
			}
			if mutable && gm.Cost(v, w) != c {
				t.Errorf("%s: visit(%d): (%d,%d); cost(%d,%d): %v\n", mess, v, w, c, v, w, gm.Cost(v, w))
			}
			return
		})
		CheckAbort(mess, t, g, v, deg)
		switch {
		case mutable:
			if deg != gm.Degree(v) {
				t.Errorf("%s: visit(%d) found %v neighbors; degree(%d): %v\n", mess, v, deg, v, gm.Degree(v))
			}
		case immutable:
			if deg != gi.Degree(v) {
				t.Errorf("%s: visit(%d) found %v neighbors; degree(%d): %v\n", mess, v, deg, v, gi.Degree(v))
			}
		}
		for w, ok := range visited {
			if ok {
				continue
			}
			if mutable && gm.Edge(v, w) || immutable && gi.Edge(v, w) {
				t.Errorf("%s: visit(%d,%d): --; edge(%d,%d): true\n", mess, v, w, v, w)
			}
		}
		if immutable {
			CheckStart(mess, t, gi, v, visited)
		}
	}
	return
}

// Check VisitFrom from all different starting points.
func CheckStart(mess string, t *testing.T, g *Immutable, v int, visited []bool) {
	n := g.Order()
	check := make([]bool, len(visited)) // Check that we get the same neighbors.
	for a := 0; a <= n+1; a++ {
		prev := -1
		for v := a; v < n; v++ {
			check[v] = visited[v]
		}
		g.VisitFrom(v, a, func(w int, _ int64) (skip bool) {
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

// Check Visit for failed abort and wrong return value.
func CheckAbort(mess string, t *testing.T, g Iterator, v int, deg int) {
	for n := 1; n <= deg; n++ {
		count := 0
		aborted := g.Visit(v, func(_ int, _ int64) bool {
			count++
			return count == n // break after n iterations
		})
		if count != n {
			t.Errorf("%s: visit(%d) didnt't abort after %d iteration(s).\n", mess, v, n)
		}
		if !aborted {
			t.Errorf("%s: visit(%d) didnt't return true after %d iteration(s).\n", mess, v, n)
		}
		aborted = g.Visit(v, func(_ int, _ int64) bool {
			return false
		})
		if aborted {
			t.Errorf("%s: visit(%d) returned true after completed iteration.\n", mess, v)
		}

	}
}

type Multi struct{}

func (m Multi) Order() int {
	return 3
}

func (m Multi) Visit(v int, action func(w int, c int64) (skip bool)) (aborted bool) {
	if v == 1 {
		for i := 0; i < 4; i++ {
			if action(0, 5) {
				return true
			}
		}
	}
	if v == 0 {
		for i := 0; i < 3; i++ {
			if action(0, 0) {
				return true
			}
		}
		for i := 0; i < 2; i++ {
			if action(0, 5) {
				return true
			}
		}
		for i := 0; i < 5; i++ {
			if action(1, 5) {
				return true
			}
		}
		for i := 0; i < 1; i++ {
			if action(1, 7) {
				return true
			}
		}
	}
	return
}

func TestString(t *testing.T) {
	res := String(Multi{})
	exp := "3 [3×(0 0) 2×(0 0):5 4×{0 1}:5 (0 1):5 (0 1):7]"
	if mess, diff := diff(res, exp); diff {
		t.Errorf("String: %s", mess)
	}
}

func TestCheck(t *testing.T) {
	g := New(0)
	res := Check(g)
	exp := Stats{0, 0, 0, 0, 0}
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check: %s", mess)
	}
	res = Check(Sort(g))
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check immutable: %s", mess)
	}

	g = New(1)
	res = Check(g)
	exp = Stats{0, 0, 0, 0, 1}
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check: %s", mess)
	}
	res = Check(Sort(g))
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check immutable: %s", mess)
	}

	g = New(1)
	g.AddCost(0, 0, 1)
	res = Check(g)
	exp = Stats{1, 0, 1, 1, 0}
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check %s", mess)
	}
	res = Check(Sort(g))
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check immutable: %s", mess)
	}

	g = New(3)
	g.AddCost(0, 1, 1)
	res = Check(g)
	exp = Stats{1, 0, 1, 0, 2}
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check: %s", mess)
	}
	res = Check(Sort(g))
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check immutable: %s", mess)
	}

	g.AddCost(1, 0, 2)
	res = Check(g)
	exp = Stats{2, 0, 2, 0, 1}
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check: %s", mess)
	}
	res = Check(Sort(g))
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check immutable: %s", mess)
	}

	g.AddCost(1, 0, 1)
	res = Check(g)
	exp = Stats{2, 0, 2, 0, 1}
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check: %s", mess)
	}
	res = Check(Sort(g))
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check immutable: %s", mess)
	}

	g.Add(2, 2)
	res = Check(g)
	exp = Stats{3, 0, 2, 1, 0}
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check: %s", mess)
	}
	res = Check(Sort(g))
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check immutable: %s", mess)
	}

	g.Add(1, 2)
	res = Check(g)
	exp = Stats{4, 0, 2, 1, 0}
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check: %s", mess)
	}
	res = Check(Sort(g))
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check immutable: %s", mess)
	}

	g.AddBoth(0, 1)
	res = Check(g)
	exp = Stats{4, 0, 0, 1, 0}
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check: %s", mess)
	}
	res = Check(Sort(g))
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check immutable: %s", mess)
	}

	res = Check(Multi{})
	exp = Stats{3, 12, 12, 5, 1}
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check: %s", mess)
	}
	res = Check(Sort(Multi{}))
	if mess, diff := diff(res, exp); diff {
		t.Errorf("Check immutable: %s", mess)
	}
}

func TestEqual(t *testing.T) {
	g := New(0)
	if mess, diff := diff(Equal(g, g), true); diff {
		t.Errorf("Equal: %s", mess)
	}

	h := New(1)
	if mess, diff := diff(Equal(h, h), true); diff {
		t.Errorf("Equal: %s", mess)
	}
	if mess, diff := diff(Equal(g, h), false); diff {
		t.Errorf("Equal: %s", mess)
	}
	if mess, diff := diff(Equal(h, g), false); diff {
		t.Errorf("Equal: %s", mess)
	}

	g = New(1)
	if mess, diff := diff(Equal(g, h), true); diff {
		t.Errorf("Equal: %s", mess)
	}
	if mess, diff := diff(Equal(h, g), true); diff {
		t.Errorf("Equal: %s", mess)
	}

	g.Add(0, 0)
	if mess, diff := diff(Equal(g, h), false); diff {
		t.Errorf("Equal: %s", mess)
	}
	if mess, diff := diff(Equal(h, g), false); diff {
		t.Errorf("Equal: %s", mess)
	}

	h.Add(0, 0)
	if mess, diff := diff(Equal(g, h), true); diff {
		t.Errorf("Equal: %s", mess)
	}
	if mess, diff := diff(Equal(h, g), true); diff {
		t.Errorf("Equal: %s", mess)
	}

	g.AddCost(0, 0, 1)
	if mess, diff := diff(Equal(g, h), false); diff {
		t.Errorf("Equal: %s", mess)
	}
	if mess, diff := diff(Equal(h, g), false); diff {
		t.Errorf("Equal: %s", mess)
	}

	g = New(3)
	h = New(3)
	g.AddBoth(0, 1)
	g.AddCost(0, 1, 1)
	g.AddBoth(1, 2)
	g.Add(0, 2)
	h.AddBoth(0, 1)
	h.AddCost(0, 1, 1)
	h.AddBoth(1, 2)
	h.Add(0, 2)
	if mess, diff := diff(Equal(g, h), true); diff {
		t.Errorf("Equal: %s", mess)
	}
	if mess, diff := diff(Equal(h, g), true); diff {
		t.Errorf("Equal: %s", mess)
	}

	g.Add(0, 0)
	if mess, diff := diff(Equal(g, h), false); diff {
		t.Errorf("Equal: %s", mess)
	}
	if mess, diff := diff(Equal(h, g), false); diff {
		t.Errorf("Equal: %s", mess)
	}

	g.Delete(0, 0)
	g.Delete(0, 2)
	if mess, diff := diff(Equal(g, h), false); diff {
		t.Errorf("Equal: %s", mess)
	}
	if mess, diff := diff(Equal(h, g), false); diff {
		t.Errorf("Equal: %s", mess)
	}
}
