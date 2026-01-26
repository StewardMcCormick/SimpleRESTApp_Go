package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/model"
	"github.com/StewardMcCormick/SimpleRESTApp_Go/internal/usecase"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	UserUseCase usecase.UserUseCase
	Validator   *validator.Validate
}

func InitHttpHandler(userUseCase usecase.UserUseCase) http.Handler {
	mux := http.NewServeMux()
	handler := &Handler{
		UserUseCase: userUseCase,
		Validator:   validator.New(),
	}

	mux.HandleFunc("GET /", handler.getHello)
	mux.HandleFunc("GET /users/{id}", handler.getById)
	mux.HandleFunc("GET /users", handler.getAll)
	mux.Handle("POST /users", isJSONValidMiddleware(http.HandlerFunc(handler.postSave)))
	mux.Handle("PUT /users/{id}", isJSONValidMiddleware(http.HandlerFunc(handler.putUser)))
	mux.Handle("PATCH /users/{id}", isJSONValidMiddleware(http.HandlerFunc(handler.patchUser)))
	mux.HandleFunc("DELETE /users/{id}", handler.delete)

	h := Chain(
		loggingMiddleware,
		jsonContentTypeMiddleware,
	)(mux)

	return h
}

func isJSONValidMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			sendError(w, errors.New("cannot read request body"), http.StatusBadRequest)
			return
		}

		if len(body) == 0 {
			sendError(w, errors.New("request body is required"), http.StatusBadRequest)
			return
		}

		if !json.Valid(body) {
			sendError(w, errors.New("invalid JSON body"), http.StatusBadRequest)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(w, r)
	})
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
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
	fmt.Fprint(w, `{"message": "Hello from User REST-service!"}`)
}

func (h *Handler) getById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		sendError(w, errors.New(fmt.Sprintf("incorrect value for id: %s", r.PathValue("id"))),
			http.StatusBadRequest)
		return
	}

	user, err := h.UserUseCase.GetById(id)
	if err != nil {
		sendError(w, err, http.StatusNotFound)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	users := h.UserUseCase.GetAll()

	response, err := json.Marshal(users)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func (h *Handler) postSave(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var user model.PostUserRequest
	json.Unmarshal(body, &user)

	if err := h.Validator.Struct(user); err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	savedUser, err := h.UserUseCase.Create(user)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(savedUser)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func (h *Handler) putUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	var request model.PutUserRequest
	json.Unmarshal(body, &request)

	if err = h.Validator.Struct(request); err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.UserUseCase.Put(id, request)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	w.Write(response)
}

func (h *Handler) patchUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	err = h.UserUseCase.Delete(id)
	if err != nil {
		sendError(w, err, http.StatusNotFound)
		return
	}
}
