package graph

// ShortestPath computes a shortest path from v to w.
// Only edges with non-negative costs are included.
// The number dist is the length of the path, or -1 if w cannot be reached.
//
// The time complexity is O((|E| + |V|)⋅log|V|), where |E| is the number of edges
// and |V| the number of vertices in the graph.
func ShortestPath(g Iterator, v, w int) (path []int, dist int64) {
	n := g.Order()
	dists := make([]int64, n)
	parents := make([]int, n)
	for i := range dists {
		dists[i], parents[i] = -1, -1
	}
	dists[v] = 0

	// Dijkstra's algorithm
	q := emptyPrioQueue(dists)
	q.Push(v)
	for q.Len() > 0 {
		u := q.Pop()
		if u == w {
			for x := w; x != -1; x = parents[x] {
				path = append(path, x)
			}
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			return path, dists[w]
		}
		g.Visit(u, func(x int, d int64) (skip bool) {
			if d < 0 {
				return
			}
			alt := dists[u] + d
			switch {
			case dists[x] == -1:
				dists[x], parents[x] = alt, u
				q.Push(x)
			case alt < dists[x]:
				dists[x], parents[x] = alt, u
				q.Fix(x)
			}
			return
		})
	}
	return []int{}, -1
}

// ShortestPaths computes the shortest paths from v to all other vertices.
// Only edges with non-negative costs are included.
// The number parent[w] is the predecessor of w on a shortest path from v to w,
// or -1 if none exists.
// The number dist[w] equals the length of a shortest path from v to w,
// or is -1 if w cannot be reached.
//
// The time complexity is O((|E| + |V|)⋅log|V|), where |E| is the number of edges
// and |V| the number of vertices in the graph.
func ShortestPaths(g Iterator, v int) (parent []int, dist []int64) {
	n := g.Order()
	dist = make([]int64, n)
	parent = make([]int, n)
	for i := range dist {
		dist[i], parent[i] = -1, -1
	}
	dist[v] = 0

	// Dijkstra's algorithm
	Q := emptyPrioQueue(dist)
	Q.Push(v)
	for Q.Len() > 0 {
		v := Q.Pop()
		g.Visit(v, func(w int, d int64) (skip bool) {
			if d < 0 {
				return
			}
			alt := dist[v] + d
			switch {
			case dist[w] == -1:
				dist[w], parent[w] = alt, v
				Q.Push(w)
			case alt < dist[w]:
				dist[w], parent[w] = alt, v
				Q.Fix(w)
			}
			return
		})
	}
	return
}
