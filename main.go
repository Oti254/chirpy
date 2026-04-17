package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	// Allocates and returns a new HTTP multiplexer
	mux := http.NewServeMux()

	// Creates an instance of http.Server
	srv := &http.Server{
		// The NewServeMux becomes the server's handler
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Printf("You are listening on port: %v", port)
	// Starts the server
	log.Fatal(srv.ListenAndServe())

}
