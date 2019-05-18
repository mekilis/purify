package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mekilis/purify"
	"github.com/pborman/getopt"
)

var (
	optPortNumber     *int
	optHelp           *bool
	optCurrentVersion *bool
)

const (
	// InvalidRequest puts 400 status code in English
	InvalidRequest = "invalid request body"
	// CurrentVersion at the moment
	CurrentVersion = "0.1"
)

func init() {
	const (
		// Port specifies the port to run on
		Port = 9002
	)

	optPortNumber = getopt.IntLong("port", 'p', Port, "port number to run on")
	optCurrentVersion = getopt.BoolLong("version", 'v', "prints current version of Purify")
	optHelp = getopt.BoolLong("help", 'h', "show help")
	getopt.SetUsage(func() {
		fmt.Println("Purify server\n\nA simple word filter API written in Go." +
			"\n\nUsage:\n\tpurify [options]" +
			"\n\nOptions:" +
			"\n\t-p <number>,\t--port <number>" +
			"\n\t-v,\t\t--version" +
			"\n\t-h,\t\t--help")
	})
}

// Request implements a basic request structure comprising
// only the message to clean
type Request struct {
	Message string `json:"message"`
}

// Response implements a JSON structure to parse to the client
type Response struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func main() {
	getopt.Parse()
	if *optCurrentVersion {
		fmt.Println("purify version", CurrentVersion)
		os.Exit(0)
	}

	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}

	color.Set(color.FgYellow)
	log.Print("Starting purify...\t")

	badWords, err := purify.ParseDictionary()
	if err != nil {
		log.Fatal(err)
	}
	color.Unset()

	color.Set(color.FgRed)
	log.Print("Setting up trie...\t")
	color.Unset()

	trie := purify.NewTrie()

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

func rootHandler(t *purify.Trie) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request Request
		var response = Response{
			StatusCode: -1,
			Status:     InvalidRequest,
			Message:    "",
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Connection", "open")

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil || request.Message == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&response)
			return
		}

		response.Message = purify.Filter(t, request.Message)
		response.StatusCode = 1
		response.Status = "successfully filtered"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response)
	}
}

func start(t *purify.Trie) {
	r := chi.NewRouter()
	// r.Use(middleware.Logger)
	r.Use(middleware.Timeout(20 * time.Second))

	r.Post("/", rootHandler(t))

	color.Set(color.FgGreen)
	log.Printf("Listening on port: %d\n", *optPortNumber)
	color.Unset()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *optPortNumber), r))
}
