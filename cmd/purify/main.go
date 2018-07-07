package main

import (
	"fmt"

	"github.com/mekilis/purify/pkg/structures"
)

func main() {
	trie := structures.NewTrie()
	fmt.Println(trie)
	fmt.Println(trie.FindWord("boy"))
	fmt.Println(trie.AddWord("boy"))
	fmt.Println(trie.FindWord("boy"))
	fmt.Println(trie.AddWord("boy"))
}
