package purify

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const pathToSystemDictionary = "/usr/share/dict/words"

// ChatterBox implements a struct to generate (sic)intelligible random words
type ChatterBox struct {
	NumberOfWords  int
	Words          []string
	BadWords       *Trie
	VocabularySize int
}

// NewChatterbox returns a new ChatterBox object
func NewChatterbox(clean bool) *ChatterBox {
	badWordsTrie := NewTrie()
	words, err := ParseDictionary(pathToSystemDictionary)
	if err != nil {
		log.Fatal("an error occurred while parsing dictionary:", err)
	}

	if !clean {
		badWords, err := ParseDictionary()
		if err != nil {
			log.Fatal(err)
		}
		for _, word := range badWords {
			badWordsTrie.AddWord(word)
		}

		words = append(badWords, words...)
	}

	return &ChatterBox{
		NumberOfWords:  10,
		Words:          words,
		BadWords:       badWordsTrie,
		VocabularySize: len(words),
	}
}

// Rant parrots whatever randomly comes to mind
func (c *ChatterBox) Rant() string {
	rant := ""
	for i := 0; i < c.NumberOfWords; i++ {
		rant += c.Words[rand.Int()%c.VocabularySize] + " "
	}
	return rant
}
