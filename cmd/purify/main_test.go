package main

import (
	"strings"
	"testing"

	"github.com/mekilis/purify/pkg/chatterbox"
)

func BenchmarkClean(b *testing.B) {
	cases := []struct {
		Name      string
		VocabSize int
	}{
		{"Normal", -1},
		{"ExcessivelyVulgar", 10},
	}

	for _, c := range cases {
		b.Run(c.Name, func(b *testing.B) {
			randomUser := chatterbox.New(false)
			if c.VocabSize != -1 {
				randomUser.VocabularySize = c.VocabSize
			}
			randomUser.NumberOfWords += b.N
			words := strings.Split(randomUser.Rant(), " ")
			Clean(randomUser, words)
		})
	}
}
