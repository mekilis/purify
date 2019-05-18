package purify

import (
	"testing"
)

func TestParseDictionary(t *testing.T) {
	// valid
	_, err := ParseDictionary()
	if err != nil {
		t.Errorf("parseDictionary(...): got %s, wanted %v", err.Error(), nil)
	}

	// file 404
	_, err = ParseDictionary("data/fakelist.txt")
	if err != errFileNotFound {
		t.Errorf("parseDictionary(...): got %s, wanted %s", err.Error(), errFileNotFound)
	}
}
