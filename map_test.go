package iterloopbenchmarks

import (
	"testing"

	"github.com/barweiss/go-tuple"
	"mtoohey.com/iter"
)

func BenchmarkMapFiltrationIter(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
	iter.KVZip(m).Filter(func(t tuple.T2[int, int]) bool {
		return t.V1%2 == 0
	}).Consume()
}

func BenchmarkMapFiltrationGoIshLoop(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
	for k := range m {
		if k%2 != 0 {
			delete(m, k)
		}
	}
}

func BenchmarkMapFiltrationUnrealisticLoop(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
	t := make([]tuple.T2[int, int], b.N/2+1)
	for k, v := range m {
		if k%2 == 0 {
			t[k/2] = tuple.New2(k, v)
		}
	}
}

func BenchmarkMapFiltrationEquivalentLoop(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
	var t []tuple.T2[int, int]
	for k, v := range m {
		if k%2 == 0 {
			t = append(t, tuple.New2(k, v))
		}
	}
}
