package iterloopbenchmarks

import (
	"testing"

	"mtoohey.com/iter"
)

func BenchmarkIntCollectionIter(b *testing.B) {
	iter.Ints[int]().Take(b.N).Collect()
}

func BenchmarkIntCollectionPracticalLoop(b *testing.B) {
	ints := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		ints[i] = i
	}
}

func BenchmarkIntCollectionEquivalentLoop(b *testing.B) {
	var ints []int
	for i := 0; i < b.N; i++ {
		ints = append(ints, i)
	}
}
