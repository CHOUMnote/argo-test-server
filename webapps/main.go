package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	greeting := os.Getenv("HELLO_ARGOCD")
	if greeting == "" {
		greeting = "Hello from Go (but GREETING not set)"
	}
	fmt.Fprintf(w, "%s\n", greeting)
}

func main() {
	http.HandleFunc("/", handler)
	port := ":8080"
	log.Printf("Starting server on %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
