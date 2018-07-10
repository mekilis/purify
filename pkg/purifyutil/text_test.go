package purifyutil

import (
	"fmt"
	"testing"

	"github.com/mekilis/purify/pkg/chatterbox"
)

func TestClean(t *testing.T) {
	chatterbox := chatterbox.New(false)
	chatterbox.NumberOfWords = 2
	chatterbox.VocabularySize = 10
	words := Clean(chatterbox.BadWords, chatterbox.Rant())
	fmt.Println(words)
}
