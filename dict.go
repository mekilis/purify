package purify

import (
	"bufio"
	"errors"
	"os"
)

// Dict implements a quasi hash map
type Dict struct {
	Key   rune
	Value *Trie
}

const pathToWords = "data/bad.txt"

var (
	errFileNotFound = errors.New("file could not be opened")
	errScanningFile = errors.New("an error occurred while scanning")
)

// ParseDictionary returns a slice containing either good or bad words. Defaults to bad words
// if no argument is specified.
func ParseDictionary(path ...string) ([]string, error) {
	words := make([]string, 0)
	p := pathToWords
	if len(path) > 0 {
		p = path[0]
	}

	file, err := os.Open(p)
	if err != nil {
		// log.Printf("%v\n", err.Error())
		return words, errFileNotFound
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		// log.Printf("%v", err)
		return words, err
	}

	return words, nil
}
