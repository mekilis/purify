package purify

import (
	"log"
	"regexp"
	"strings"
)

// This variable matches basic patterns to the possibly intended runes
var replacer *strings.Replacer

func init() {
	replacer = strings.NewReplacer(
		"@", "a", "(", "c", "3", "e",
		"1", "i", "!", "i", "0", "o",
		"9", "q", "5", "s", "v", "u",
		"2", "z", "+", "t",
	)
}

// Filter sanitizes a given text based on a passed trie of bad words
func Filter(trie *Trie, words string) string {
	r, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}

	wordsSlice := strings.Split(words, " ")
	wordsDone := make([]string, 0)
	for _, word := range wordsSlice {
		filtered := word
		word = strings.ToLower(word)
		replaced := replacer.Replace(word)
		rWord := r.ReplaceAllString(word, "")
		rReplaced := r.ReplaceAllString(replaced, "")

		if trie.Find(word) || trie.Find(replaced) || trie.Find(rWord) || trie.Find(rReplaced) {
			filtered = clean(word)
		}

		wordsDone = append(wordsDone, filtered)
	}

	return strings.Join(wordsDone, " ")
}

func clean(s string) string {
	var sb strings.Builder
	sb.WriteByte(s[0])

	size := len(s)
	for i := 2; i < size; i++ {
		sb.WriteRune('*')
	}

	sb.WriteByte(s[size-1])
	return sb.String()
}
