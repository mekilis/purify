package structures

import (
	"errors"
)

var (
	// ErrDuplicateWord returns a soft error indicating a word's duplicate status
	ErrDuplicateWord = errors.New("word already exist")
)

// Trie implements the data structure of same name
type Trie struct {
	End        bool
	ChildTries Dict
}

// NewTrie returns a new Trie object
func NewTrie() *Trie {
	return &Trie{
		End:        false,
		ChildTries: make(Dict),
	}
}

// Find returns true if a given word exists in the word map or false otherwise
func (t *Trie) Find(word string) bool {
	subTrie := t
	for _, r := range word {
		trie, ok := subTrie.ChildTries[r]
		if !ok {
			return false // not found
		}

		subTrie = trie
	}

	return subTrie.End
}

// AddWord adds a new distinct word to the trie
func (t *Trie) AddWord(word string) error {
	if ok := t.Find(word); ok {
		return ErrDuplicateWord
	}

	subTrie := t
	for _, r := range word {
		trie, ok := subTrie.ChildTries[r]
		if !ok {
			trie = new(Trie)
			trie.ChildTries = make(Dict)
			subTrie.ChildTries[r] = trie
		}
		subTrie = trie
	}

	if !subTrie.End {
		subTrie.End = true
	}

	return nil
}
