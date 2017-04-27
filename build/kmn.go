package build

import "strconv"

// Kmn returns a virtual complete bipartite graph with m+n vertices.
// The vertices are divided into two subsets U = [0..m) and V = [m..m+n),
// and the edge set consists of all edges {u, v}, where u ∊ U and v ∊ V.
func Kmn(m, n int) *Virtual {
	switch {
	case m < 0 || n < 0:
		return nil
	case m == 0 && n == 0:
		return null
	case m == 0 && n == 1 || m == 1 && n == 0:
		return singleton()
	case m == 1 && n == 1:
		return edge()
	case m+n <= 0:
		panic("too large m=" + strconv.Itoa(m) + " n=" + strconv.Itoa(n))
	}
	res := generic0(m+n, func(v, w int) bool {
		return v < m && w >= m || v >= m && w < m
	})
	res.degree = func(v int) int {
		if v < m {
			return n
		}
		return m
	}
	res.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		if v < m {
			for w := max(a, m); w < m+n; w++ {
				if do(w, 0) {
					return true
				}
			}
			return
		}
		for w := a; w < m; w++ {
			if do(w, 0) {
				return true
			}
		}
		return
	}
	return res
}
