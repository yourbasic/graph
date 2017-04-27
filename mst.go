package graph

import (
	"container/heap"
)

// MST computes a minimum spanning tree for each connected component
// of an undirected weighted graph.
// The forest of spanning trees is returned as a slice of parent pointers:
// parent[v] is either the parent of v in a tree,
// or -1 if v is the root of a tree.
//
// The time complexity is O(|E|⋅log|V|), where |E| is the number of edges
// and |V| the number of vertices in the graph.
func MST(g Iterator) (parent []int) {
	n := g.Order()
	parent = make([]int, n)
	cost := make([]int64, n)
	for i := range parent {
		parent[i] = -1
		cost[i] = Max
	}

	// Prim's algorithm
	Q := newMstQueue(cost)
	for Q.Len() > 0 {
		v := heap.Pop(Q).(int)
		g.Visit(v, func(w int, c int64) (skip bool) {
			if Q.Contains(w) && c < cost[w] {
				cost[w] = c
				Q.Update(w)
				parent[w] = v
			}
			return
		})
	}
	return
}

type mstQueue struct {
	heap  []int // vertices in heap order
	index []int // index of each vertex in the heap
	cost  []int64
}

func newMstQueue(cost []int64) *mstQueue {
	n := len(cost)
	h := &mstQueue{
		heap:  make([]int, n),
		index: make([]int, n),
		cost:  cost,
	}
	for i := range h.heap {
		h.heap[i] = i
		h.index[i] = i
	}
	return h
}

func (m *mstQueue) Len() int { return len(m.heap) }

func (m *mstQueue) Less(i, j int) bool {
	return m.cost[m.heap[i]] < m.cost[m.heap[j]]
}

func (m *mstQueue) Swap(i, j int) {
	m.heap[i], m.heap[j] = m.heap[j], m.heap[i]
	m.index[m.heap[i]] = i
	m.index[m.heap[j]] = j
}

func (m *mstQueue) Push(x interface{}) {} // Not used

func (m *mstQueue) Pop() interface{} {
	n := len(m.heap) - 1
	v := m.heap[n]
	m.index[v] = -1
	m.heap = m.heap[:n]
	return v
}

func (m *mstQueue) Update(v int) {
	heap.Fix(m, m.index[v])
}

func (m *mstQueue) Contains(v int) bool {
	return m.index[v] >= 0
}
