package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Tag     string `json:"tag"`
	Value   any    `json:"value"`
}

func sendError(w http.ResponseWriter, err error, code int) {
	var response ErrorResponse

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errorSlice := make([]ValidationError, len(ve))
		for i, fe := range ve {
			errorSlice[i] = ValidationError{
				Field:   fe.Field(),
				Message: getValidationErrorMessage(fe),
				Tag:     fe.Tag(),
				Value:   fe.Value(),
			}
		}

		response = ErrorResponse{
			Status:  code,
			Message: "Validation error",
			Errors:  errorSlice,
		}
	} else {
		response = ErrorResponse{
			Status:  code,
			Message: err.Error(),
		}
	}

	log.Printf("%s", err)
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		http.Error(w, "JSON Marshalling error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(jsonResponse)
}

func getValidationErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Incorrect email"
	case "min":
		return fmt.Sprintf("Minimum lenght - %s", fe.Param())
	case "max":
		return fmt.Sprintf("Maximum lenght - %s", fe.Param())
	default:
		return fe.Error()
	}
}
