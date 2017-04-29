package build

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSize(t *testing.T) {
	if mess, diff := diff(VertexSet{}.size(), -1); diff {
		t.Errorf("size %s", mess)
	}
	if mess, diff := diff(Vertex(0).size(), 1); diff {
		t.Errorf("size %s", mess)
	}
	if mess, diff := diff(Range(0, 0).size(), 0); diff {
		t.Errorf("size %s", mess)
	}
	if mess, diff := diff(Range(1, 1).size(), 0); diff {
		t.Errorf("size %s", mess)
	}
	if mess, diff := diff(Range(0, -1).size(), 0); diff {
		t.Errorf("size %s", mess)
	}
	if mess, diff := diff(Range(0, 1).size(), 1); diff {
		t.Errorf("size %s", mess)
	}
	if mess, diff := diff(Range(2, 4).size(), 2); diff {
		t.Errorf("size %s", mess)
	}
}

func TestContains(t *testing.T) {
	if mess, diff := diff(VertexSet{}.Contains(0), true); diff {
		t.Errorf("Contains %s", mess)
	}
	if mess, diff := diff(VertexSet{}.Contains(-1), true); diff {
		t.Errorf("Contains %s", mess)
	}
	if mess, diff := diff(VertexSet{}.Contains(1), true); diff {
		t.Errorf("Contains %s", mess)
	}
	if mess, diff := diff(Vertex(0).Contains(0), true); diff {
		t.Errorf("Contains %s", mess)
	}
	if mess, diff := diff(Vertex(0).Contains(1), false); diff {
		t.Errorf("Contains %s", mess)
	}
	if mess, diff := diff(Range(0, 0).Contains(0), false); diff {
		t.Errorf("Contains %s", mess)
	}
	if mess, diff := diff(Range(1, 1).Contains(1), false); diff {
		t.Errorf("Contains %s", mess)
	}
	if mess, diff := diff(Range(0, -1).Contains(0), false); diff {
		t.Errorf("Contains %s", mess)
	}
	if mess, diff := diff(Range(0, 1).Contains(0), true); diff {
		t.Errorf("Contains %s", mess)
	}
	if mess, diff := diff(Range(0, 1).Contains(1), false); diff {
		t.Errorf("Contains %s", mess)
	}
	if mess, diff := diff(Range(2, 4).Contains(1), false); diff {
		t.Errorf("Contains %s", mess)
	}
	if mess, diff := diff(Range(2, 4).Contains(2), true); diff {
		t.Errorf("Contains %s", mess)
	}
	if mess, diff := diff(Range(2, 4).Contains(3), true); diff {
		t.Errorf("Contains %s", mess)
	}
	if mess, diff := diff(Range(2, 4).Contains(4), false); diff {
		t.Errorf("Contains %s", mess)
	}
}

func TestGetRank(t *testing.T) {
	v1 := VertexSet{set: []interval{{0, 2, 0}, {4, 6, 0}, {8, 10, 0}, {12, 14, 0}, {16, 18, 0}}}
	v1.update()
	v2 := VertexSet{set: []interval{{2, 3, 0}, {5, 8, 0}, {9, 10, 0}, {11, 13, 0}, {15, 19, 0}}}
	v2.update()

	var res bytes.Buffer
	for i := 0; i < 12; i++ {
		fmt.Fprint(&res, v1.get(i), " ")
	}
	exp := "0 1 4 5 8 9 12 13 16 17 -1 -1 "
	if !(res.String() == exp) {
		t.Errorf("VertexSet.get: %v; want %v", res, exp)
	}

	res.Reset()
	for i := 0; i < 12; i++ {
		fmt.Fprint(&res, v2.get(i), " ")
	}
	exp = "2 5 6 7 9 11 12 15 16 17 18 -1 "
	if !(res.String() == exp) {
		t.Errorf("VertexSet.get: %v; want %v", res.String(), exp)
	}

	res.Reset()
	for i := 0; i < 20; i++ {
		fmt.Fprint(&res, v1.rank(i), " ")
	}
	exp = "0 1 -1 -1 2 3 -1 -1 4 5 -1 -1 6 7 -1 -1 8 9 -1 -1 "
	if !(res.String() == exp) {
		t.Errorf("VertexSet.rank: %v; want %v", res.String(), exp)
	}

	res.Reset()
	for i := 0; i < 20; i++ {
		fmt.Fprint(&res, v2.rank(i), " ")
	}
	exp = "-1 -1 0 -1 -1 1 2 3 -1 4 -1 5 6 -1 -1 7 8 9 10 -1 "
	if !(res.String() == exp) {
		t.Errorf("VertexSet.rank: %v; want %v", res.String(), exp)
	}
}

// equals returns true if the sets are identical.
func (s1 VertexSet) equals(s2 VertexSet) bool {
	if len(s1.set) != len(s2.set) {
		return false
	}
	for i, in := range s1.set {
		if in.a != s2.set[i].a || in.b != s2.set[i].b {
			return false
		}
	}
	return true
}

func TestAndNot(t *testing.T) {
	res := Range(1, 5).AndNot(Vertex(3))
	exp := Range(1, 3).Or(Range(4, 5))
	if !res.equals(exp) {
		t.Errorf("AndNot: %v; want %v", res, exp)
	}

	res = VertexSet{}.AndNot(VertexSet{})
	exp = empty()
	if !res.equals(exp) {
		t.Errorf("AndNot: %v; want %v", res, exp)
	}

	res = empty().AndNot(Vertex(3))
	exp = empty()
	if !res.equals(exp) {
		t.Errorf("AndNot: %v; want %v", res, exp)
	}

	res = Range(1, 5).AndNot(Range(1, 3)).AndNot(Range(3, 5))
	exp = empty()
	if !res.equals(exp) {
		t.Errorf("AndNot: %v; want %v", res, exp)
	}
}

func TestAndOr(t *testing.T) {
	v1 := VertexSet{set: []interval{{0, 2, 0}, {4, 6, 0}, {8, 10, 0}, {12, 14, 0}, {16, 18, 0}}}
	v1.update()
	v2 := VertexSet{set: []interval{{2, 3, 0}, {5, 8, 0}, {9, 10, 0}, {11, 13, 0}, {15, 19, 0}}}
	v2.update()

	expInter := VertexSet{set: []interval{{5, 6, 0}, {9, 10, 0}, {12, 13, 0}, {16, 18, 0}}}
	if !v1.And(v2).equals(expInter) {
		t.Errorf("And: %v; want %v", v1.And(v2), expInter)
	}

	expUnion := VertexSet{set: []interval{{0, 3, 0}, {4, 10, 0}, {11, 14, 0}, {15, 19, 0}}}
	if !v1.Or(v2).equals(expUnion) {
		t.Errorf("Or: %v; want %v", v1.Or(v2), expUnion)
	}

	v := Range(1, 3)
	if mess, diff := diff(VertexSet{}.And(v).Contains(0), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(VertexSet{}.And(v).Contains(-1), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(VertexSet{}.And(v).Contains(1), true); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Vertex(0).And(v).Contains(0), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Vertex(0).And(v).Contains(1), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(0, 0).And(v).Contains(0), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(1, 1).And(v).Contains(1), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(0, -1).And(v).Contains(0), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(0, 1).And(v).Contains(0), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(0, 1).And(v).Contains(1), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(2, 4).And(v).Contains(1), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(2, 4).And(v).Contains(2), true); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(2, 4).And(v).Contains(3), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(2, 4).And(v).Contains(4), false); diff {
		t.Errorf("And %s", mess)
	}

	if mess, diff := diff(Range(0, 4).And(v).Contains(0), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(0, 4).And(v).Contains(1), true); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(0, 4).And(v).Contains(2), true); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(0, 4).And(v).Contains(3), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(v.And(Range(0, 4)).Contains(0), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(v.And(Range(0, 4)).Contains(1), true); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(v.And(Range(0, 4)).Contains(2), true); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(v.And(Range(0, 4)).Contains(3), false); diff {
		t.Errorf("And %s", mess)
	}

	if mess, diff := diff(Range(0, 2).And(v).Contains(0), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(0, 2).And(v).Contains(1), true); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(Range(0, 2).And(v).Contains(2), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(v.And(Range(0, 2)).Contains(0), false); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(v.And(Range(0, 2)).Contains(1), true); diff {
		t.Errorf("And %s", mess)
	}
	if mess, diff := diff(v.And(Range(0, 2)).Contains(2), false); diff {
		t.Errorf("And %s", mess)
	}

	if mess, diff := diff(VertexSet{nil}.Or(VertexSet{nil}), VertexSet{nil}); diff {
		t.Errorf("Or %s", mess)
	}
}
