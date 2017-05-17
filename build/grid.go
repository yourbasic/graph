package build

import "strconv"

// Grid returns a virtual graph whose vertices correspond to integer
// points in the plane: y-coordinates being in the range 0..m-1,
// and x-coordinates in the range 0..n-1. Two vertices of a grid
// are adjacent whenever the corresponding points are at distance 1.
//
// Point (x, y) gets index nx + y, and index i corresponds to the point (i/n, i%n).
func Grid(m, n int) *Virtual {
	switch {
	case m < 0 || n < 0:
		return nil
	case m == 0 || n == 0:
		return null
	case m == 1 && n == 1:
		return singleton()
	case m == 1 && n == 2 || m == 2 && n == 1:
		return edge()
	case m == 1 && n > 2:
		return line(n)
	case m > 2 && n == 1:
		return line(m)
	case m*n/m != n:
		panic("too large m=" + strconv.Itoa(m) + " n=" + strconv.Itoa(n))
	}

	g := generic0(m*n, func(v, w int) (edge bool) {
		xdiff, ydiff := v/n-w/n, v%n-w%n
		switch {
		case xdiff == 0:
			return ydiff == -1 || ydiff == 1
		case ydiff == 0:
			return xdiff == -1 || xdiff == 1
		}
		return
	})

	g.degree = func(v int) (deg int) {
		vn := v % n
		if v >= n { // up
			deg++
		}
		if vn != 0 { // left
			deg++
		}
		if vn != n-1 { // right
			deg++
		}
		if v < g.order-n { // down
			deg++
		}
		return
	}

	g.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		if v < a-n {
			return
		}
		if w := v - n; w >= a && do(w, 0) { // up
			return true
		}
		vn := v % n
		if vn != 0 { // left
			if w := v - 1; w >= a && do(w, 0) {
				return true
			}
		}
		if vn != n-1 { // right
			if w := v + 1; w >= a && do(w, g.cost(v, w)) {
				return true
			}
		}
		if v < g.order-n { // down
			return do(v+n, 0)
		}
		return
	}
	return g
}
