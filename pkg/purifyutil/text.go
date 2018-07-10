package purifyutil

import (
	"strings"
	"sync"

	"github.com/mekilis/purify/pkg/structures"
)

// Regex matches basic patterns to the possibly intended runes
var Regex map[rune]rune

func init() {
	Regex = map[rune]rune{
		'@': 'a',
		'(': 'c',
		'3': 'e',
		'1': 'i',
		'!': 'i',
		'0': 'o',
		'9': 'q',
		'5': 's',
		'v': 'u',
		'2': 'z',
	}
}

// Clean sanitizes a given text based on a passed trie of bad words
func Clean(trie *structures.Trie, words string) string {
	wordsSlice := strings.Split(words, " ")
	wordsDone := make([]bool, len(wordsSlice))
	var mutex = &sync.Mutex{}

	for i, word := range wordsSlice {
		tempWord := word
		word = strings.ToLower(word)

		ch := make(chan int)
		possibleMatch := false

		go func(wordArg string) (wordReturned string) {
			wordRegex := ""

			for _, r := range wordArg {
				if match, ok := Regex[r]; ok {
					possibleMatch = true
					wordRegex += string(match)
					continue
				}
				wordRegex += string(r)
			}

			if possibleMatch && trie.Find(wordRegex) {
				// repeat again only if there's a possible match
				wRegex, k := string(tempWord[0]), len(word)
				for j := 2; j < k; j++ {
					wRegex += "*"
				}
				wRegex += string(tempWord[k-1])

				mutex.Lock()
				wordsDone[i] = true
				wordsSlice[i] = wRegex
				mutex.Unlock()
			}

			ch <- 1
			return
		}(word)

		if trie.Find(word) {
			w, k := string(tempWord[0]), len(word)
			for j := 2; j < k; j++ {
				w += "*"
			}
			w += string(tempWord[k-1])

			mutex.Lock()
			if wordsDone[i] {
				mutex.Unlock()
				continue
			}
			wordsSlice[i] = w
			mutex.Unlock()
		}

		<-ch
		close(ch)
	}

	return strings.Join(wordsSlice, " ")
}
