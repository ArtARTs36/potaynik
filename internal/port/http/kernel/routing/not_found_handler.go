package routing

import (
	"github.com/rs/zerolog/log"
)

func handleNotFound(request Request) Response {
	log.Debug().Msgf("Route with uri %s not found", request.URI())

	return NewNotFoundResponse("Route not found")
}
