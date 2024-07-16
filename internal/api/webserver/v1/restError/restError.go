package resterror

import "net/http"

type RestError struct {
	Message string   `json:"message"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	File    string `json:"file"`
	Message string `json:"message"`
}

func (r *RestError) Error() string {
	return r.Message
}

func NewBadRequestError(message string, causes ...Causes) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusNotFound,
		Causes:  nil,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusInternalServerError,
		Causes:  nil,
	}
}
