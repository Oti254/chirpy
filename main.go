package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}
	// Allocates and returns a new HTTP multiplexer
	// Responsible for routing
	// Matches requests with specified patterns
	mux := http.NewServeMux()

	// Adding a handler for the root path
	// Adding the fileserver location for root
	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.Handle("/app/", fsHandler)
	mux.HandleFunc("GET /api/healthz", handleReadiness)
	mux.HandleFunc("GET /api/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /api/reset", apiCfg.resetHandler)

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
