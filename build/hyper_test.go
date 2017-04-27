package build

import "testing"

func TestHyper(t *testing.T) {
	if mess, diff := diff(Hyper(0).String(), "1 []"); diff {
		t.Errorf("Hyper %s", mess)
	}

	if mess, diff := diff(Hyper(1).String(), "2 [{0 1}]"); diff {
		t.Errorf("Hyper %s", mess)
	}

	exp := "4 [{0 1} {0 2} {1 3} {2 3}]"
	if mess, diff := diff(Hyper(2).String(), exp); diff {
		t.Errorf("Hyper %s", mess)
	}

	exp = "8 [{0 1} {0 2} {0 4} {1 3} {1 5} {2 3} {2 6} {3 7} {4 5} {4 6} {5 7} {6 7}]"
	if mess, diff := diff(Hyper(3).String(), exp); diff {
		t.Errorf("Hyper %s", mess)
	}

	for n := 0; n < 4; n++ {
		Consistent("Hyper", t, Hyper(n))
	}

	g := Hyper(bitsPerWord - 2) // maximum possible size
	if mess, diff := diff(g.Edge(0, 1), true); diff {
		t.Errorf("Hyper %s", mess)
	}
}
