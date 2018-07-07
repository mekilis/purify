package chatterbox

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRant(t *testing.T) {
	cleanSpeech := false
	chatterBox := New(cleanSpeech)
	chatterBox.NumberOfWords = 1 + rand.Int()%5
	rant := chatterBox.Rant()
	if rant == "" {
		t.Error("chatterBox is suddenly mute")
	}

	fmt.Println(rant)
}