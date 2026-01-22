package main

import (
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/handler"
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/repository"
	"net/http"
)

const (
	host = "localhost"
	port = "8080"
)

func main() {
	userRepo := repository.NewInMemoryUserRepository()
	h := handler.InitHttpHandler(userRepo)

	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: h,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
