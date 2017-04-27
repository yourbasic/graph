package build

import "testing"

func TestTree(t *testing.T) {
	exp := "0 []"
	if mess, diff := diff(Tree(1, 0).String(), exp); diff {
		t.Errorf("Tree %s", mess)
	}

	exp = "1 []"
	if mess, diff := diff(Tree(1, 1).String(), exp); diff {
		t.Errorf("Tree %s", mess)
	}

	exp = "2 [{0 1}]"
	if mess, diff := diff(Tree(1, 2).String(), exp); diff {
		t.Errorf("Tree %s", mess)
	}

	exp = "0 []"
	if mess, diff := diff(Tree(2, 0).String(), exp); diff {
		t.Errorf("Tree %s", mess)
	}

	exp = "1 []"
	if mess, diff := diff(Tree(2, 1).String(), exp); diff {
		t.Errorf("Tree %s", mess)
	}

	exp = "3 [{0 1} {0 2}]"
	if mess, diff := diff(Tree(2, 2).String(), exp); diff {
		t.Errorf("Tree %s", mess)
	}

	exp = "7 [{0 1} {0 2} {1 3} {1 4} {2 5} {2 6}]"
	if mess, diff := diff(Tree(2, 3).String(), exp); diff {
		t.Errorf("Tree %s", mess)
	}

	exp = "4 [{0 1} {0 2} {0 3}]"
	if mess, diff := diff(Tree(3, 2).String(), exp); diff {
		t.Errorf("Tree %s", mess)
	}

	exp = "13 [{0 1} {0 2} {0 3} {1 4} {1 5} {1 6} {2 7} {2 8} {2 9} {3 10} {3 11} {3 12}]"
	if mess, diff := diff(Tree(3, 3).String(), exp); diff {
		t.Errorf("Tree %s", mess)
	}

	for k := 1; k < 4; k++ {
		for n := 0; n < 4; n++ {
			Consistent("Tree", t, Tree(k, n))
		}
	}
}
