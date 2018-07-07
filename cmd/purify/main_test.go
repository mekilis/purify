package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/mekilis/purify/pkg/chatterbox"
)

func TestMain(m *testing.M) {
	flag.Parse()
	i := m.Run()
	fmt.Print(i)
	os.Exit(i)
}

func TestStart(t *testing.T) {
}

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

// Clean does the actual job of filtering a given string of possible profane
// words (testing only)
func Clean(chatterBox *chatterbox.ChatterBox, wordsSlice []string) (int, int) {
	var goodWords, badWords int
	for i, word := range wordsSlice {
		if chatterBox.BadWords.Find(word) {
			w := string(word[0])
			k := len(word)
			for j := 2; j < k; j++ {
				w += "*"
			}
			w += string(word[k-1])
			wordsSlice[i] = w
			badWords++
		} else {
			goodWords++
		}
	}
	return goodWords, badWords
}
