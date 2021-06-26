package errors

import (
	"errors"
	"net/http"
)

type RestError struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadReqError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}

}
func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}

}
func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}

}
