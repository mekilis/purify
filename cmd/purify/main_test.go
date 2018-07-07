package main

import (
	"strings"
	"testing"

	"github.com/mekilis/purify/pkg/chatterbox"
)

func BenchmarkClean_Normal(b *testing.B) {
	// fmt.Printf("sample size: %d ", b.N)

	randomUser := chatterbox.New(false)
	randomUser.NumberOfWords += b.N

	words := randomUser.Rant()
	wordsSlice := strings.Split(words, " ")

	_, _ = Clean(randomUser, wordsSlice)
	// fmt.Printf("\ngood words: %d\t\t\tbad words: %d", goodWords, badWords)
}

func BenchmarkClean_Excessive(b *testing.B) {
	// fmt.Printf("sample size: %d ", b.N)

	randomUser := chatterbox.New(false)
	randomUser.NumberOfWords = b.N
	randomUser.VocabularySize = 10 //bad words are in the range 0-100

	words := randomUser.Rant()
	wordsSlice := strings.Split(words, " ")

	_, _ = Clean(randomUser, wordsSlice)
	// fmt.Printf("good words: %d\t\t\tbad words: %d", goodWords, badWords)
}
