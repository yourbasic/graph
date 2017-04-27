package build

import "strconv"

// Hyper returns a virtual hypercube graph with 2ⁿ vertices.
// Two vertices of a hypercube are adjacent when their binary
// representations differ in a single digit.
func Hyper(n int) *Virtual {
	switch {
	case n < 0:
		return nil
	case n == 0:
		return singleton()
	case n == 1:
		return line(2)
	case n > bitsPerWord-2:
		panic("n=" + strconv.Itoa(n) + " too big; max=" + strconv.Itoa(bitsPerWord-2))
	}

	res := generic0(1<<uint(n), func(v, w int) (edge bool) {
		d := v ^ w
		return d&-d == d
	})

	res.degree = func(v int) (deg int) {
		return n
	}

	res.visit = func(v int, a int, do func(w int, c int64) bool) (aborted bool) {
		if v >= a {
			for b := 1 << uint(n); b > 0; b >>= uint(1) {
				if v&b == 0 {
					continue
				}
				if w := v ^ b; w >= a && do(w, 0) {
					return true
				}
			}
		}
		for b := 1; b < 1<<uint(n); b <<= uint(1) {
			if v&b != 0 {
				continue
			}
			if w := v ^ b; w >= a && do(w, 0) {
				return true
			}
		}
		return
	}
	return res
}
