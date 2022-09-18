package routing

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

const defaultContentType = "application/json"

type HttpHandlerWrapper struct {
	handler AppHandler
}

func WrapHandler(handler AppHandler) GoHttpHandler {
	wrapper := HttpHandlerWrapper{handler: handler}

	return wrapper.Handle
}

func (r *HttpHandlerWrapper) Handle(writer http.ResponseWriter, request *http.Request) {
	resp := r.handler(NewRequest(request))

	writer.Header().Set("Content-Type", defaultContentType)
	writer.WriteHeader(resp.Code)
	_, err := writer.Write(resp.Message)

	if err != nil {
		log.Error().Msgf("[Http][HandlerWrapper] Error on response writing: %v", err)
	}
}
