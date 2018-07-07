package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mekilis/purify/pkg/structures"
)

const (
	Port           = "9002"
	InvalidRequest = "invalid request body"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func main() {
	color.Set(color.FgYellow)
	log.Print("Starting purify...\t")
	badWords, err := structures.ParseDictionary()
	if err != nil {
		log.Fatal(err)
	}
	color.Unset()

	color.Set(color.FgRed)
	log.Print("Setting up trie...\t")
	color.Unset()
	trie := structures.NewTrie()
	for _, word := range badWords {
		err = trie.AddWord(word)
		if err != nil {
			log.Fatal(err)
		}
	}
	color.Unset()

	var stopSignal = make(chan os.Signal)
	signal.Notify(stopSignal, syscall.SIGTERM)
	signal.Notify(stopSignal, syscall.SIGINT)

	go func() {
		<-stopSignal
		// finish pending processes if any
		time.Sleep(time.Second)
		fmt.Println("\nBye!")
		os.Exit(0)
	}()

	start(trie)
}

func root(t *structures.Trie) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request Request
		var response = Response{
			StatusCode: -1,
			Status:     InvalidRequest,
			Message:    "",
		}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil || request.Message == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&response)
			return
		}

		wordsSlice := strings.Split(request.Message, " ")
		for i, word := range wordsSlice {
			if t.Find(word) {
				w, k := string(word[0]), len(word)
				for j := 2; j < k; j++ {
					w += "*"
				}
				w += string(word[k-1])
				wordsSlice[i] = w
			}
		}

		response.Message = strings.Join(wordsSlice, " ")
		response.StatusCode = 1
		response.Status = "successfully filtered"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response)
	}
}

func start(t *structures.Trie) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(20 * time.Second))

	r.Post("/", root(t))

	color.Set(color.FgGreen)
	log.Printf("Listening on port: %s\n", Port)
	color.Unset()

	log.Fatal(http.ListenAndServe(":"+Port, r))
}
