package build

import "testing"

func TestKmn(t *testing.T) {
	if mess, diff := diff(Kmn(0, 0).String(), "0 []"); diff {
		t.Errorf("Kmn %s", mess)
	}

	if mess, diff := diff(Kmn(1, 1).String(), "2 [{0 1}]"); diff {
		t.Errorf("Kmn %s", mess)
	}

	if mess, diff := diff(Kmn(1, 2).String(), "3 [{0 1} {0 2}]"); diff {
		t.Errorf("Kmn %s", mess)
	}

	for m := 0; m < 5; m++ {
		for n := 0; n < 10; n++ {
			Consistent("Kmn", t, Kmn(m, n))
		}
	}
}
