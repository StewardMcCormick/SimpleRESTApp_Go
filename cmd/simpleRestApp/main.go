package main

import (
	"fmt"
	"net/http"
)

const (
	host = "localhost"
	port = "8080"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/hello", hello)

	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: mux,
	}

	_ = server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
