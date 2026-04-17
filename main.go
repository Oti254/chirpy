package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"
	// Allocates and returns a new HTTP multiplexer
	mux := http.NewServeMux()

	// Adding a handler for the root path
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	// Creates an instance of http.Server
	srv := &http.Server{
		// The NewServeMux becomes the server's handler
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	// Starts the server
	log.Fatal(srv.ListenAndServe())

}
