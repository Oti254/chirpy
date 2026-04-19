package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"
	// Allocates and returns a new HTTP multiplexer
	// Responsible for routing
	// Matches requests with specified patterns
	mux := http.NewServeMux()

	// Adding a handler for the root path
	// Adding the fileserver location for root
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.HandleFunc("/healthz", handlerReadiness)

	// Creates an instance of http.Server
	srv := &http.Server{
		// The NewServeMux becomes the server's handler config
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	// Starts the server
	log.Fatal(srv.ListenAndServe())

}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text.html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
