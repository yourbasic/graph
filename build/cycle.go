package build

// Cycle returns a virtual cycle graph with n vertices and
// the edges {0, 1}, {1, 2}, {2, 3},... , {n-1, 0}.
func Cycle(n int) *Virtual {
	switch {
	case n < 0:
		return nil
	case n == 0:
		return null
	case n == 1:
		return singleton()
	case n == 2:
		return edge()
	}

	g := generic0(n, func(v, w int) (edge bool) {
		switch v - w {
		case 1 - n, -1, 1, n - 1:
			edge = true
		}
		return
	})

	// Precondition : n ≥ 3.
	g.degree = func(v int) int { return 2 }

	// Precondition : n ≥ 3.
	g.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		var w [2]int
		switch v {
		case 0:
			w = [2]int{1, n - 1}
		case n - 1:
			w = [2]int{0, n - 2}
		default:
			w = [2]int{v - 1, v + 1}
		}
		for _, w := range w {
			if w >= a && do(w, 0) {
				return true
			}
		}
		return
	}
	return g
}
