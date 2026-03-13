package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// A minimal echo and counter server
// that returns header and form data

var count int = 0
var mu sync.Mutex

func main() {

	http.HandleFunc("/count", counterHandler)
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s:%s", k, v)
	}

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	fmt.Fprintf(w, "\nHello from %q: (%d)", r.URL.Path, count)
}

func counterHandler(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	fmt.Fprintf(w, "\n Count : %d \n", count)
	mu.Unlock()
}
