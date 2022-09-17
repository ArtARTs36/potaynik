package routing

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func handleNotFound(writer http.ResponseWriter, request *http.Request) {
	log.Debug().Msgf("Route with uri %s not found", request.RequestURI)

	writer.WriteHeader(404)
}
