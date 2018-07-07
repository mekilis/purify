package structures

import (
	"bufio"
	"errors"
	"os"
)

// Dict implements a quasi hash map
type Dict map[string]*Trie

const pathToBadWords = "/mnt/data/purify/bad.txt"

var (
	errFileNotFound = errors.New("file could not be opened")
	errScanningFile = errors.New("an error occured while scanning")
)

// ParseDictionary returns a slice containing either good or bad words
func ParseDictionary(path ...string) ([]string, error) {
	words := make([]string, 0)
	p := pathToBadWords
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

	// if err := scanner.Err(); err != nil {
	// 	// log.Printf("%v", err)
	// 	return words, err
	// }

	return words, nil
}
