package main

import "github.com/mekilis/purify/pkg/chatterbox"

func main() {
}

// Clean does the actual job of filtering a given string of possible profane
// words
func Clean(chatterBox *chatterbox.ChatterBox, wordsSlice []string) (int, int) {
	var goodWords, badWords int
	for i, word := range wordsSlice {
		if chatterBox.BadWords.FindWord(word) {
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
