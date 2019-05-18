package purify

import (
	"testing"
)

func BenchmarkFilter(b *testing.B) {
	cb := NewChatterbox(false)

	tests := []struct {
		Name      string
		VocabSize int
	}{
		{"Normal", -1},
		{"Vulgar", 10},
	}

	for _, tt := range tests {
		if tt.VocabSize != -1 {
			cb.VocabularySize = tt.VocabSize
		}
		temp := cb.NumberOfWords
		cb.NumberOfWords += b.N
		Filter(cb.BadWords, cb.Rant())
		cb.NumberOfWords = temp
	}
}
