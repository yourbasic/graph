package build

import (
	"github.com/yourbasic/graph"
	"testing"
)

func TestCirculant(t *testing.T) {
	res := Circulant(1, 0)
	if mess, diff := diff(res.String(), "1 []"); diff {
		t.Errorf("Circulant %s", mess)
	}
	Consistent("Circulant", t, res)

	res = Circulant(-1)
	if mess, diff := diff(res, (*Virtual)(nil)); diff {
		t.Errorf("Circulant %s", mess)
	}
	Consistent("Circulant", t, res)

	res = Circulant(0)
	if mess, diff := diff(res.String(), "0 []"); diff {
		t.Errorf("Circulant %s", mess)
	}
	Consistent("Circulant", t, res)

	res = Circulant(1)
	if mess, diff := diff(res.String(), "1 []"); diff {
		t.Errorf("Circulant %s", mess)
	}
	Consistent("Circulant", t, res)

	res = Circulant(2, 0)
	if mess, diff := diff(res.String(), "2 []"); diff {
		t.Errorf("Circulant %s", mess)
	}
	Consistent("Circulant", t, res)

	res = Circulant(2, 1)
	if mess, diff := diff(res.String(), "2 [{0 1}]"); diff {
		t.Errorf("Circulant %s", mess)
	}
	Consistent("Circulant", t, res)

	res = Circulant(2, 0, 1, 1, 2)
	if mess, diff := diff(res.String(), "2 [{0 1}]"); diff {
		t.Errorf("Circulant %s", mess)
	}
	Consistent("Circulant", t, res)

	res = Circulant(2, 0, -1, -3, -2)
	if mess, diff := diff(res.String(), "2 [{0 1}]"); diff {
		t.Errorf("Circulant %s", mess)
	}
	Consistent("Circulant", t, res)

	res = Circulant(2, 0, 1, 2)
	if mess, diff := diff(res.String(), "2 [{0 1}]"); diff {
		t.Errorf("Circulant %s", mess)
	}
	Consistent("Circulant", t, res)

	res = Circulant(4, 1)
	if mess, diff := diff(res.String(), "4 [{0 1} {0 3} {1 2} {2 3}]"); diff {
		t.Errorf("Circulant %s", mess)
	}
	Consistent("Circulant", t, res)

	res = Circulant(4, -5)
	if mess, diff := diff(res.String(), "4 [{0 1} {0 3} {1 2} {2 3}]"); diff {
		t.Errorf("Circulant %s", mess)
	}
	Consistent("Circulant", t, res)

	res = Circulant(6, 1, 2, 3)
	exp := Kn(6)
	if mess, diff := diff(graph.Equal(res, exp), true); diff {
		t.Errorf("Circulant %s", mess)
	}
	Consistent("Circulant", t, res)

	for n := 1; n < 10; n++ {
		Consistent("Circulant", t, Circulant(n, n/2, n/3-n, n/4))
	}
}
