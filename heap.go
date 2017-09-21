package graph

import (
	"container/heap"
)

type prioQueue struct {
	heap  []int // vertices in heap order
	index []int // index of each vertex in the heap
	cost  []int64
}

func emptyQueue(cost []int64) *prioQueue {
	return &prioQueue{
		index: make([]int, len(cost)),
		cost:  cost,
	}
}

func newQueue(cost []int64) *prioQueue {
	n := len(cost)
	h := &prioQueue{
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

func (m *prioQueue) Len() int { return len(m.heap) }

func (m *prioQueue) Less(i, j int) bool {
	return m.cost[m.heap[i]] < m.cost[m.heap[j]]
}

func (m *prioQueue) Swap(i, j int) {
	m.heap[i], m.heap[j] = m.heap[j], m.heap[i]
	m.index[m.heap[i]] = i
	m.index[m.heap[j]] = j
}

func (pq *prioQueue) Push(x interface{}) {
	n := len(pq.heap)
	v := x.(int)
	pq.heap = append(pq.heap, v)
	pq.index[v] = n
}

func (m *prioQueue) Pop() interface{} {
	n := len(m.heap) - 1
	v := m.heap[n]
	m.index[v] = -1
	m.heap = m.heap[:n]
	return v
}

func (m *prioQueue) Update(v int) {
	heap.Fix(m, m.index[v])
}

func (m *prioQueue) Contains(v int) bool {
	return m.index[v] >= 0
}
