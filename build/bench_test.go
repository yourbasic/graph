package build

import (
	"testing"
)

func BenchmarkKn(b *testing.B) {
	g := Kn(100)
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkComplement(b *testing.B) {
	g := Empty(100).Complement()
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkSubgraph(b *testing.B) {
	g := Kn(1000).Subgraph(Range(0, 100))
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkConnect(b *testing.B) {
	g := Kn(50).Connect(0, Kn(50))
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkJoin(b *testing.B) {
	g := Kn(1).Join(Kn(100), AllEdges())
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkCost(b *testing.B) {
	g := Kn(100).AddCost(8)
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkCostFunc(b *testing.B) {
	g := Kn(100).AddCostFunc(Cost(8))
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkDeleteNone(b *testing.B) {
	g := Kn(100).Delete(NoEdges())
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkDeleteOne(b *testing.B) {
	g := Kn(100).Delete(Edge(3, 8))
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkDeleteAll(b *testing.B) {
	g := Kn(100).Delete(AllEdges())
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkAddNone(b *testing.B) {
	g := Kn(100).Add(NoEdges())
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkAddOne(b *testing.B) {
	g := Kn(100).Add(Edge(3, 8))
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkAddAll(b *testing.B) {
	g := Kn(100).Add(AllEdges())
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkAddAllCache(b *testing.B) {
	g := Specific(Kn(100).Add(AllEdges()))
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkUnion(b *testing.B) {
	g := Kn(100).Union(Kn(100))
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkIntersect(b *testing.B) {
	g := Kn(100).Intersect(Kn(100))
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkTensor(b *testing.B) {
	g := Kn(10).Tensor(Kn(10))
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}

func BenchmarkCartesian(b *testing.B) {
	g := Kn(50).Cartesian(Kn(50))
	for i := 0; i < b.N; i++ {
		g.Visit(0, func(w int, c int64) (skip bool) {
			return
		})
	}
}
