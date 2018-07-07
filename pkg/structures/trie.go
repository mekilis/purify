package structures

import (
	"errors"
)

var (
	//ErrTypeAssert returns type assert errors
	ErrTypeAssert = errors.New("casting error")

	// ErrDuplicateWord returns a soft error indicating a word's duplicate status
	ErrDuplicateWord = errors.New("word already exist")
)

// Trie implements the data structure of same name
type Trie struct {
	End      bool
	Children Dict
}

// TODO: Use int instead of string as keys

// NewTrie returns a new Trie object
func NewTrie() *Trie {
	return &Trie{
		End:      false,
		Children: make(Dict),
	}
}

// FindWord returns true if a given word exists in the word map or false otherwise
func (t *Trie) FindWord(word string) bool {
	subTrie := t
	for _, c := range word {
		ch := string(c)

		child, ok := subTrie.Children[ch]
		if !ok {
			return false // not found
		}

		subTrie = child
	}

	return subTrie.End
}

// AddWord adds a new distinct word to the trie
func (t *Trie) AddWord(word string) error {
	if ok := t.FindWord(word); ok {
		return ErrDuplicateWord
	}

	subTrie := t
	for _, r := range word {
		ch := string(r)
		child, ok := subTrie.Children[ch]
		if ok {
			subTrie = child
		} else {
			child = new(Trie)
			child.Children = make(Dict)
			subTrie.Children[ch] = child
			subTrie = child
		}
	}

	// subTrie.Children[t.End] = t.End // use a boolean instead of an interface
	if !subTrie.End {
		subTrie.End = true
	}

	return nil
}
