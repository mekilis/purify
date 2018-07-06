package main

import (
	"fmt"

	"github.com/mekilis/purify/pkg/structures"
)

func main() {
	words := make([]string, 0)
	words = append(words, "boy")
	trie, err := structures.NewTrie(words)
	fmt.Println(trie, err)
	// fmt.Println(trie.FindWord("boy"))
	// fmt.Println(trie.FindWord("by"))
}
