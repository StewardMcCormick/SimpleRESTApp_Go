package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	host = "localhost"
	port = "8080"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/hello", hello)

	handler := loggingMiddleware(mux)

	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[NEW REQUEST]: Addr: %s, Method: %s", r.URL, r.Method)

		start := time.Now()
		next.ServeHTTP(w, r)

		log.Printf("[RESPONSE]: Addr: %s, Method: %s, Total millis: %d", r.URL, r.Method, time.Since(start))
	})
}
