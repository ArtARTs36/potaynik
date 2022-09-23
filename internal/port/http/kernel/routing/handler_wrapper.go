package routing

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

const defaultContentType = "application/json"

type HTTPHandlerWrapper struct {
	handler AppHandler
}

func WrapHandler(handler AppHandler) GoHTTPHandler {
	wrapper := HTTPHandlerWrapper{handler: handler}

	return wrapper.Handle
}

func (r *HTTPHandlerWrapper) Handle(writer http.ResponseWriter, request *http.Request) {
	resp := r.handler(NewRequest(request))

	writer.Header().Set("Content-Type", defaultContentType)
	writer.WriteHeader(resp.Code)
	_, err := writer.Write(resp.Message)

	if err != nil {
		log.Error().Msgf("[HTTP][HandlerWrapper] Error on response writing: %v", err)
	}
}
