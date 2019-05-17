package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/mekilis/purify"
)

func TestMain(m *testing.M) {
	flag.Parse()
	if testing.Short() {
		fmt.Println("short test")
		os.Exit(m.Run())
	}

	fmt.Println("long test")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("working directory could not be accessed", err)
	}
	dir = filepath.Join(dir, "purify")

	build := exec.Command("go", "build", "-o", dir)
	o, err := build.CombinedOutput()
	if err != nil {
		log.Fatalln(err, string(o))
	}

	time.Sleep(2 * time.Second)
	os.Remove(dir)
	os.Exit(0)
}

func TestRootHandler(t *testing.T) {
	trie := purify.NewTrie()
	got := rootHandler(trie)
	if got == nil {
		t.Errorf("handler cannot be nil")
	}
}

func BenchmarkMain(b *testing.B) {
	url := fmt.Sprintf("http://localhost:%d", *optPortNumber)
	chatterBox := purify.NewChatterbox(false)
	var json []byte
	var request = new(http.Request)
	var response = new(http.Response)
	var err error
	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	for i := 0; i < b.N; i++ {
		json = []byte(fmt.Sprintf("{\"message\": \"%s\"}", chatterBox.Rant()))
		request, err = http.NewRequest("POST", url, bytes.NewBuffer(json))
		if err != nil {
			// log.Println("failed to generate request", err)
			continue
		}
		request.Header.Set("Content-Type", "application/json")

		response, err = client.Do(request)
		if err != nil {
			// log.Println("failed to do request", err)
			continue
		}
		response.Body.Close()
	}
}

func BenchmarkClean(b *testing.B) {
	cases := []struct {
		Name      string
		VocabSize int
	}{
		{"Normal", -1},
		{"ExcessivelyVulgar", 10},
	}

	for _, c := range cases {
		b.Run(c.Name, func(b *testing.B) {
			randomUser := purify.NewChatterbox(false)
			if c.VocabSize != -1 {
				randomUser.VocabularySize = c.VocabSize
			}
			randomUser.NumberOfWords += b.N
			words := strings.Split(randomUser.Rant(), " ")
			Clean(randomUser, words)
		})
	}
}

// Clean does the actual job of filtering a given string of possible profane
// words (testing only)
func Clean(chatterBox *purify.ChatterBox, wordsSlice []string) (int, int) {
	var goodWords, badWords int
	for i, word := range wordsSlice {
		if chatterBox.BadWords.Find(word) {
			w := string(word[0])
			k := len(word)
			for j := 2; j < k; j++ {
				w += "*"
			}
			w += string(word[k-1])
			wordsSlice[i] = w
			badWords++
		} else {
			goodWords++
		}
	}
	return goodWords, badWords
}
