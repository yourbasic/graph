package build

// Tree returns a full k-ary tree with n levels and (kⁿ-1)/(k-1) vertices.
// The parent of vertex v is (v-1)/k,
// and its children are kv + 1, kv + 2,… , kv + k.
func Tree(k, n int) *Virtual {
	switch {
	case n < 0 || k < 1:
		return nil
	case n == 0:
		return null
	case n == 1:
		return singleton()
	case k == 1:
		return line(n)
	}

	size := 1
	for i := n; i > 0; i-- {
		if k*size/k != size {
			panic("tree too big")
		}
		size *= k
	}
	size = (size - 1) / (k - 1)
	lastRow := 1 + (size-2)/k

	res := generic0(size, func(v, w int) (edge bool) {
		switch {
		case v == 0:
			return w >= 1 && w <= k
		case v < lastRow:
			return w == (v-1)/k || w >= k*v+1 && w <= k*v+k
		default:
			return w == (v-1)/k
		}
	})

	res.degree = func(v int) (deg int) {
		if v != 0 {
			deg++
		}
		if v < lastRow {
			deg += k
		}
		return
	}

	res.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		if v != 0 {
			if w := (v - 1) / k; w >= a && do(w, 0) {
				return true
			}
		}
		if v < lastRow {
			for w := max(a, k*v+1); w <= k*(v+1); w++ {
				if do(w, 0) {
					return true
				}
			}
		}
		return
	}
	return res
}
