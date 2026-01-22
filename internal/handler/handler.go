package handler

import (
	"encoding/json"
	"fmt"
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/model"
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/repository"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	UserRepo repository.UserRepository
}

func InitHttpHandler(userRepo repository.UserRepository) http.Handler {
	mux := http.NewServeMux()
	handler := &Handler{UserRepo: userRepo}

	mux.HandleFunc("GET /", handler.getHello)
	mux.HandleFunc("GET /users/{id}", handler.getById)
	mux.HandleFunc("GET /users/", handler.getAll)
	mux.HandleFunc("POST /users/", handler.postSave)

	h := loggingMiddleware(mux)
	h = JSONContentTypeMiddleware(h)

	return h
}

func JSONContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[NEW REQUEST]: Addr: %s, Method: %s", r.URL, r.Method)

		start := time.Now()
		next.ServeHTTP(w, r)

		log.Printf("[RESPONSE]: Addr: %s, Method: %s, Total millis: %d", r.URL, r.Method, time.Since(start).Milliseconds())
	})
}

func (h *Handler) getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from User REST-service!")
}

func (h *Handler) getById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Incorrect value in request: %v", err)
		http.Error(w, fmt.Sprintf("incorrect value for id: %s", r.PathValue("id")), http.StatusBadRequest)
		return
	}

	user, err := h.UserRepo.GetById(id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusNoContent)
	}

	response, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error JSON Marshaling: %v", err)
		http.Error(w, "Error JSON Marshaling", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	users := h.UserRepo.GetAll()

	response, err := json.Marshal(users)
	if err != nil {
		log.Printf("Error JSON Marshaling: %v", err)
		http.Error(w, "Error JSON Marshaling", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func (h *Handler) postSave(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Reading request body error: %v", err)
		http.Error(w, "Reading request body error", http.StatusInternalServerError)
		return
	}

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("JSON unmarshaling error: %v", err)
		http.Error(w, "Incorrect request body", http.StatusBadRequest)
		return
	}

	savedUser, err := h.UserRepo.Save(user)
	if err != nil {
		log.Printf("Save User error: %v", err)
		http.Error(w, "Saving error", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(savedUser)
	w.Write(response)
}
