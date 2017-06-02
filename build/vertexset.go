package build

import "sort"

// VertexSet represents a set of vertices in a graph.
// The zero value of a VertexSet is the universe;
// the set containing all vertices.
type VertexSet struct {
	// A set is an immutable sorted list of non-empty disjoint intervals.
	// The zero value VertexSet{nil} represents the universe.
	set []interval
}

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

// An interval represents the numbers [a, b).
type interval struct {
	a, b  int
	index int // index of a in the whole set
}

// update updates the index values.
func (s VertexSet) update() {
	prev := 0
	for i, in := range s.set {
		s.set[i].index = prev
		prev += in.b - in.a
	}
}

// empty returns the empty set.
func empty() VertexSet {
	return VertexSet{[]interval{}}
}

// Range returns a set containing all vertices v, a ≤ v < b.
func Range(a, b int) VertexSet {
	if a >= b {
		return empty()
	}
	return VertexSet{[]interval{{a, b, 0}}}
}

// Vertex returns a set containing the single vertex v.
func Vertex(v int) VertexSet {
	return VertexSet{[]interval{{v, v + 1, 0}}}
}

// size returns the number of elements in this set, or -1 for the universe.
func (s VertexSet) size() (size int) {
	switch {
	case s.set == nil:
		return -1
	case len(s.set) == 0:
		return 0
	}
	in := s.set[len(s.set)-1]
	return in.index + in.b - in.a
}

// get returns the i:th element in the set, or -1 if not available.
func (s VertexSet) get(i int) int {
	if i < 0 || i >= s.size() {
		return -1
	}
	// The smallest index j such that i < set[j].index.
	j := sort.Search(len(s.set), func(j int) bool {
		return i < s.set[j].index
	})
	in := s.set[j-1]
	return in.a + i - in.index
}

// rank returns the position of n in the set, or -1 if not available.
func (s VertexSet) rank(n int) int {
	if len(s.set) == 0 || n < s.set[0].a {
		return -1
	}
	// The smallest index i such that n < set[i].a
	i := sort.Search(len(s.set), func(i int) bool {
		return n < s.set[i].a
	})
	in := s.set[i-1]
	if n >= in.b {
		return -1
	}
	return in.index + n - in.a
}

// Contains tells if v is a member of the set.
func (s VertexSet) Contains(v int) bool {
	switch {
	case s.set == nil:
		return true
	case len(s.set) == 0 || v < s.set[0].a:
		return false
	}
	// The smallest index i such that v < set[i].a
	i := sort.Search(len(s.set), func(i int) bool {
		return v < s.set[i].a
	})
	if v >= s.set[i-1].b {
		return false
	}
	return true
}

// AndNot returns the set of all vertices belonging to s1 but not to s2.
func (s1 VertexSet) AndNot(s2 VertexSet) VertexSet {
	return s1.And(s2.complement())
}

// complement returns the set of all vertices not belonging to s.
func (s VertexSet) complement() VertexSet {
	switch {
	case s.set == nil:
		return empty()
	case len(s.set) == 0:
		return VertexSet{}
	}
	t := empty()
	prev := minInt
	for _, in := range s.set {
		if prev != in.a {
			t.set = append(t.set, interval{prev, in.a, 0})
		}
		prev = in.b
	}
	if prev < maxInt {
		t.set = append(t.set, interval{prev, maxInt, 0})
	}
	t.update()
	return t
}

// And returns the set of all vertices belonging to both s1 and s2.
func (s1 VertexSet) And(s2 VertexSet) VertexSet {
	switch {
	case s1.set == nil:
		return s2
	case s2.set == nil:
		return s1
	}

	type point struct {
		x int
		a bool // Tells if x is a of [a, b).
	}
	points := make([]point, 0, 2*(len(s1.set)+len(s2.set)))
	for _, in := range s1.set {
		points = append(points, point{in.a, true})
		points = append(points, point{in.b, false})
	}
	for _, in := range s2.set {
		points = append(points, point{in.a, true})
		points = append(points, point{in.b, false})
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].x == points[j].x {
			return !points[i].a
		}
		return points[i].x < points[j].x
	})

	s := empty()
	start, count := 0, 0
	for _, p := range points {
		switch count {
		case 0:
			count++
		case 1:
			if p.a {
				start = p.x
				count++
			} else {
				count--
			}
		case 2:
			s.set = append(s.set, interval{start, p.x, 0})
			count--
		}
	}
	s.update()
	return s
}

// Or returns the set of all vertices belonging to either s1 or s2.
func (s1 VertexSet) Or(s2 VertexSet) VertexSet {
	if s1.set == nil || s2.set == nil {
		return VertexSet{nil}
	}

	type point struct {
		x int
		a bool // Tells if x is a of [a, b).
	}
	points := make([]point, 0, 2*(len(s1.set)+len(s2.set)))
	for _, in := range s1.set {
		points = append(points, point{in.a, true})
		points = append(points, point{in.b, false})
	}
	for _, in := range s2.set {
		points = append(points, point{in.a, true})
		points = append(points, point{in.b, false})
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].x == points[j].x {
			return points[i].a
		}
		return points[i].x < points[j].x
	})

	s := empty()
	start, count := 0, 0
	for _, p := range points {
		switch count {
		case 0:
			start = p.x
			count++
		case 1:
			if p.a {
				count++
			} else {
				s.set = append(s.set, interval{start, p.x, 0})
				count--
			}
		case 2:
			count--
		}
	}
	s.update()
	return s
}
