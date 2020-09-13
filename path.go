package graph

// ShortestPath computes a shortest path from v to w.
// Only edges with non-negative costs are included.
// The number dist is the length of the path, or -1 if w cannot be reached.
//
// The time complexity is O((|E| + |V|)⋅log|V|), where |E| is the number of edges
// and |V| the number of vertices in the graph.
func ShortestPath(g Iterator, v, w int) (path []int, dist int64) {
	parent, distances := ShortestPaths(g, v)
	path, dist = []int{}, distances[w]
	if dist == -1 {
		return
	}
	for v := w; v != -1; v = parent[v] {
		path = append(path, v)
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return
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
	p := &pathFinder{dist: dist, parent: parent, Q: Q}
	do := p.Do
	for Q.Len() > 0 {
		v := Q.Pop()
		p.v = v
		g.Visit(v, do)
	}
	return
}

type pathFinder struct {
	dist   []int64
	parent []int
	Q      *prioQueue
	v      int
}

// Do performs a part of the shortest path algorithm.
func (p *pathFinder) Do(w int, d int64) (skip bool) {
	if d < 0 {
		return
	}
	alt := p.dist[p.v] + d
	switch {
	case p.dist[w] == -1:
		p.dist[w], p.parent[w] = alt, p.v
		p.Q.Push(w)
	case alt < p.dist[w]:
		p.dist[w], p.parent[w] = alt, p.v
		p.Q.Fix(w)
	}
	return
}
