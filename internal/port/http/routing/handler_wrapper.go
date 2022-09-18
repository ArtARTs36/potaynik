package routing

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

type HttpHandlerWrapper struct {
	handler Handler
}

func WrapHandler(handler Handler) GoHttpHandler {
	wrapper := HttpHandlerWrapper{handler: handler}

	return wrapper.Handle
}

func (r *HttpHandlerWrapper) Handle(writer http.ResponseWriter, request *http.Request) {
	resp := r.handler(NewRequest(request))

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(resp.Code)
	_, err := writer.Write(resp.Message)

	if err != nil {
		log.Error().Msgf("[Http][HandlerWrapper] Error on response writing: %v", err)
	}
}
