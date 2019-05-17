package purify

import (
	"fmt"
	"testing"
)

func TestClean(t *testing.T) {
	chatterbox := NewChatterbox(false)
	chatterbox.NumberOfWords = 2
	chatterbox.VocabularySize = 10
	words := CleanText(chatterbox.BadWords, chatterbox.Rant())
	fmt.Println(words)
}
