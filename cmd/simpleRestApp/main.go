package main

import (
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/handler"
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/repository"
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/usecase"
	"log"
	"net/http"
)

const (
	host = "localhost"
	port = "8080"
)

func main() {
	userRepo := repository.NewInMemoryUserRepository()
	userUseCase := usecase.NewUserUseCase(userRepo)
	h := handler.InitHttpHandler(userUseCase)

	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: h,
	}

	log.Printf("Start server on %s:%s", host, port)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
