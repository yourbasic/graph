package build

import "testing"

func TestCycle(t *testing.T) {
	if mess, diff := diff(Cycle(-1), (*Virtual)(nil)); diff {
		t.Errorf("Cycle %s", mess)
	}

	if mess, diff := diff(Cycle(0).String(), "0 []"); diff {
		t.Errorf("Cycle %s", mess)
	}

	if mess, diff := diff(Cycle(1).String(), "1 []"); diff {
		t.Errorf("Cycle %s", mess)
	}

	if mess, diff := diff(Cycle(2).String(), "2 [{0 1}]"); diff {
		t.Errorf("Cycle %s", mess)
	}

	if mess, diff := diff(Cycle(3).String(), "3 [{0 1} {0 2} {1 2}]"); diff {
		t.Errorf("Cycle %s", mess)
	}

	for n := 0; n < 5; n++ {
		Consistent("Cycle", t, Cycle(n))
	}
}
