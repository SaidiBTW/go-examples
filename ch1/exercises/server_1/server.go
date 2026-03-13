package main

import (
	"fmt"
	"log"
	"net/http"
)

//This is a minimal echo server

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello from %q\n", request.URL.Path)
}
