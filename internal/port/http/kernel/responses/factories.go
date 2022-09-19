package responses

import "net/http"

func ServerError(message string) Response {
	return ServerErrorFromBytes(newErrorResponseMessage(message))
}

func ServerErrorFromBytes(message []byte) Response {
	return Response{Code: http.StatusInternalServerError, Message: message}
}

func Forbidden(message string) Response {
	return Response{Code: http.StatusForbidden, Message: newErrorResponseMessage(message)}
}

func NotFound(message string) Response {
	return Response{Code: http.StatusNotFound, Message: newErrorResponseMessage(message), Headers: map[string]string{}}
}

func AlreadyReported(message string) Response {
	return Response{Code: http.StatusAlreadyReported, Message: newErrorResponseMessage(message), Headers: map[string]string{}}
}

func UnprocessableEntity(message string) Response {
	return Response{Code: http.StatusUnprocessableEntity, Message: newErrorResponseMessage(message), Headers: map[string]string{}}
}

func Created(message []byte) Response {
	return Response{Code: http.StatusCreated, Message: message, Headers: map[string]string{}}
}

func OK(message []byte) Response {
	return Response{Code: http.StatusOK, Message: message, Headers: map[string]string{}}
}
