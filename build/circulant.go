package build

import "sort"

// Circulant returns a virtual circulant graph with n vertices
// in which vertex i is adjacent to vertex (i + j) mod n and vertex (i - j) mod n
// for each j in the list s.
func Circulant(n int, s ...int) *Virtual {
	switch {
	case n < 0:
		return nil
	case n == 0:
		return null
	case n == 1:
		return singleton()
	}

	s1 := make([]int, 0, 2*len(s))
	for _, j := range s {
		j %= n
		if j < 0 {
			j += n
		}
		s1 = append(s1, j)
		s1 = append(s1, n-j)
	}
	sort.Ints(s1)
	var t []int
	prev := -1
	for _, j := range s1 {
		// Remove zeroes and duplicates.
		if j == prev || j == 0 || j == n {
			continue
		}
		prev = j
		t = append(t, j)
	}
	deg := len(t)
	switch {
	case deg == 0:
		return Empty(n)
	case deg == 1 && t[0] == 1:
		return Cycle(n)
	case deg == n-1:
		return Kn(n)
	}

	g := generic0(n, func(v, w int) (edge bool) {
		d := v - w
		if d < 0 {
			d = -d
		}
		i := sort.SearchInts(t, d)
		return i < deg && t[i] == d
	})

	g.degree = func(v int) int { return deg }

	g.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		start := sort.SearchInts(t, n-v+a)
		for i := start; i < deg; i++ {
			if do(v+t[i]-n, 0) {
				return true
			}
		}
		start = sort.SearchInts(t, a-v)
		stop := sort.SearchInts(t, n-v)
		for i := start; i < stop; i++ {
			if do(v+t[i], 0) {
				return true
			}
		}
		return
	}
	return g
}
