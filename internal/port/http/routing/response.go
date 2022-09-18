package routing

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int
	Message []byte
	Headers map[string]string
}

type errorMessage struct {
	Error string `json:"error"`
}

func NewForbiddenResponseWithText(message string) Response {
	return Response{Code: http.StatusForbidden, Message: newErrorResponseMessage(message)}
}

func NewNotFoundResponse(message string) Response {
	return Response{Code: http.StatusNotFound, Message: newErrorResponseMessage(message), Headers: map[string]string{}}
}

func NewAlreadyReportedResponse(message string) Response {
	return Response{Code: http.StatusAlreadyReported, Message: newErrorResponseMessage(message), Headers: map[string]string{}}
}

func NewInvalidEntityResponse(message string) Response {
	return Response{Code: http.StatusUnprocessableEntity, Message: newErrorResponseMessage(message), Headers: map[string]string{}}
}

func NewCreatedResponse(message []byte) Response {
	return Response{Code: http.StatusCreated, Message: message, Headers: map[string]string{}}
}

func NewOKResponse(message []byte) Response {
	return Response{Code: http.StatusOK, Message: message, Headers: map[string]string{}}
}

func newErrorResponseMessage(message string) []byte {
	jsonMsg, _ := json.Marshal(errorMessage{Error: message})

	return jsonMsg
}
