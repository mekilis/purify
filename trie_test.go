package purify

import (
	"testing"
)

func TestAddWord(t *testing.T) {
	tests := []struct {
		in   string
		want error
	}{
		{"smart", nil},
		{"mandarin", nil},
		{"mand", nil},
		{"artquake", nil},
		{"smart", ErrDuplicateWord},
	}

	trie := NewTrie()
	for _, tt := range tests {
		got := trie.AddWord(tt.in)
		if got != tt.want {
			t.Errorf("AddWord(...): got %s, wanted %s", got, tt.want)
		}
	}
}

func TestFindWord(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"smart", false},
		{"mandarin", true},
		{"mand", false},
		{"artquake", false},
		{"smart", true},
	}

	trie := NewTrie()
	trie.AddWord("mandarin")
	for i, tt := range tests {
		if i == 4 {
			trie.AddWord("smart")
		}

		got := trie.Find(tt.in)
		if got != tt.want {
			t.Errorf("FindWord(...): got %v, wanted %v", got, tt.want)
		}
	}
}

func BenchmarkAddWord(b *testing.B) {
	trie := NewTrie()
	for i := 0; i < b.N; i++ {
		trie.AddWord("word")
	}
}

func BenchmarkFind_NotExist(b *testing.B) {
	trie := NewTrie()
	for i := 0; i < b.N; i++ {
		trie.Find("word")
	}
}

func BenchmarkFind_ItExists(b *testing.B) {
	trie := NewTrie()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		trie.AddWord("word")
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		trie.Find("word")
	}
}
