package main

//Server is a minimal echo and counter server
import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int = 0

var mu sync.Mutex

func main() {

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		// To handle concurrency issues we also use a lock here
		mu.Lock()
		fmt.Fprintf(w, "%d", count)
		mu.Unlock()
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//To handle concurrency issues we could use a mutex
		mu.Lock()
		count++
		mu.Unlock()
		fmt.Fprintf(w, "Hello from %s : Request: %d", r.URL.Path, count)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
