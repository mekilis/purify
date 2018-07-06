package structures

import (
	"errors"
	"fmt"
)

var (
	//ErrTypeAssert returns type assert errors
	ErrTypeAssert = errors.New("casting error")
)

// Trie implements the data structure of same name
type Trie struct {
	End  string
	Dict map[string]interface{}
}

// TODO: Use int instead of string as keys

// NewTrie returns a new Trie object
func NewTrie(words []string) (*Trie, error) {
	trie := &Trie{
		End:  "*",
		Dict: make(map[string]interface{}),
	}

	for _, word := range words {
		for _, c := range word {
			ch := string(c)
			//TODO: Stuck
			fmt.Print(ch)
		}
		trie.Dict[trie.End] = trie.End
	}
	return trie, nil
}

// FindWord returns the true if word exists in the word map or false otherwise
func (t *Trie) FindWord(word string) (bool, error) {
	trie := t.Dict
	fmt.Println("trie 1", trie)
	for _, c := range word {
		ch := string(c)

		dict, ok := t.Dict[ch]
		if !ok {
			return false, nil // not found
		}

		// fetch inner map
		innerDict, ok := dict.(map[string]interface{})
		if ok {
			trie = innerDict
		} else {
			return false, ErrTypeAssert
		}
	}

	fmt.Println("trie 2", trie)
	_, ok := trie[t.End]
	return ok, nil
}
