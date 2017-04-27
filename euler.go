package graph

// EulerDirected returns an Euler walk in a directed graph.
// If no such walk exists, it returns an empty walk and sets ok to false.
func EulerDirected(g Iterator) (walk []int, ok bool) {
	n := g.Order()
	// Compute outdegree - indegree for each vertex.
	degree := make([]int, n)
	for v := range degree {
		g.Visit(v, func(w int, _ int64) (skip bool) {
			degree[v]++
			degree[w]--
			return
		})
	}

	start, end := -1, -1
	for v := range degree {
		switch {
		case degree[v] == 0:
		case degree[v] == 1 && start == -1:
			start = v
		case degree[v] == -1 && end == -1:
			end = v
		default:
			return []int{}, false
		}
	}

	// Make a copy of g
	edgeCount := 0
	h := make([][]int, n)
	for v := range h {
		g.Visit(v, func(w int, _ int64) (skip bool) {
			h[v] = append(h[v], w)
			edgeCount++
			return
		})
	}
	if edgeCount == 0 {
		return []int{}, true
	}

	// Find a starting point with neighbors.
	for v := 0; v < n && start == -1; v++ {
		if len(h[v]) > 0 {
			start = v
		}
	}

	stack := []int{start}
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for len(h[v]) > 0 {
			stack = append(stack, v)
			v, h[v] = h[v][0], h[v][1:]
			edgeCount--
		}
		walk = append(walk, v)
	}

	if edgeCount != 0 {
		return []int{}, false
	}

	for i, j := 0, len(walk)-1; i < j; i, j = i+1, j-1 {
		walk[i], walk[j] = walk[j], walk[i]
	}
	return walk, true
}

// EulerUndirected returns an Euler walk following undirected edges
// in only one direction. If no such walk exists, it returns an empty walk
// and sets ok to false.
func EulerUndirected(g Iterator) (walk []int, ok bool) {
	n := g.Order()
	// Compute outdegree for each vertex.
	out := make([]int, n)
	for v := range out {
		g.Visit(v, func(w int, _ int64) (skip bool) {
			if v != w {
				out[v]++
			}
			return
		})
	}

	start, oddDeg := -1, 0
	for v := range out {
		if out[v]&1 == 1 {
			start = v
			oddDeg++
		}
	}
	if !(oddDeg == 0 || oddDeg == 2) {
		return []int{}, false
	}

	// Make a copy of g.
	edgeCount := 0
	h := New(n)
	for v := 0; v < n; v++ {
		g.Visit(v, func(w int, _ int64) (skip bool) {
			h.Add(v, w)
			edgeCount++
			return
		})
	}
	if edgeCount == 0 {
		return []int{}, true
	}

	// Find a starting point with neighbors.
	for v := 0; v < n && start == -1; v++ {
		h.Visit(v, func(w int, _ int64) (skip bool) {
			start = w
			return true
		})
	}

	stack := []int{start}
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for h.Degree(v) > 0 {
			stack = append(stack, v)
			var w int
			h.Visit(v, func(u int, _ int64) (skip bool) {
				w = u
				return true
			})
			if v != w {
				h.DeleteBoth(v, w)
				edgeCount -= 2
			} else {
				h.Delete(v, v)
				edgeCount--
			}
			v = w
		}
		walk = append(walk, v)
	}

	if edgeCount != 0 {
		return []int{}, false
	}
	return walk, true
}
