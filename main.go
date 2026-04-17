package main

import (
	"log"
	"net/http"
)

func main() {
	const port = 8080
	mux := http.NewServeMux()

	srv := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	log.Printf("You are listening on port: %v", port)
	log.Fatal(http.ListenAndServe(srv.Addr, srv.Handler))

}
