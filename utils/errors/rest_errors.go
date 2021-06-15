package errors

import "net/http"

type RestError struct {
	Status int `json:"status"`
	Error string `json:"error"`
	Message string `json:"message"`
}

func NewBadReqError(message string) *RestError{
	return &RestError{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}

}
func NewNotFoundError(message string) *RestError{
	return &RestError{
		Message: message,
		Status: http.StatusNotFound,
		Error: "not_found",
	}

}