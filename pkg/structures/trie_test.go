package structures

import "testing"

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

		got := trie.FindWord(tt.in)
		if got != tt.want {
			t.Errorf("FindWord(...): got %v, wanted %v", got, tt.want)
		}
	}
}
