package main

import "net/http"

const (
	host = "localhost"
	port = "8080"
)

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: mux,
	}

	_ = server.ListenAndServe()
}
