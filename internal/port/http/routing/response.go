package routing

import "net/http"

type Response struct {
	Code    int
	Message []byte
	Headers map[string]string
}

func NewNotFoundResponse(message string) Response {
	return Response{Code: http.StatusNotFound, Message: []byte(message), Headers: map[string]string{}}
}

func NewInvalidEntityResponse(message string) Response {
	return Response{Code: http.StatusUnprocessableEntity, Message: []byte(message), Headers: map[string]string{}}
}

func NewCreatedResponse(message []byte) Response {
	return Response{Code: http.StatusCreated, Message: message, Headers: map[string]string{}}
}

func NewOKResponse(message []byte) Response {
	return Response{Code: http.StatusOK, Message: message, Headers: map[string]string{}}
}
