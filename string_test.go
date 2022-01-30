package iterloopbenchmarks

import (
	"testing"

	"mtoohey.com/iter"
)

func F(r rune) {}

func BenchmarkRunesMapIter(b *testing.B) {
	iter.Runes(string(make([]rune, b.N))).ForEach(F)
}

func BenchmarkRunesConsumeLoop(b *testing.B) {
	runes := make([]rune, b.N)
	for i := 0; i < b.N; i++ {
		F(runes[i])
	}
}
