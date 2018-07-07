package structures

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
	_, err = ParseDictionary("/mnt/data/purify/fakelist.txt")
	if err != errFileNotFound {
		t.Errorf("parseDictionary(...): got %s, wanted %s", err.Error(), errFileNotFound)
	}

	// invalid format
	_, err = ParseDictionary("/mnt/data/purify/bad.txt")
	// TODO: Find an actual error (!=)
	if err == errScanningFile {
		t.Errorf("parseDictionary(...): got %v, wanted %s", err, errScanningFile)
	}

}
