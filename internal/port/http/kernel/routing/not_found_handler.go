package routing

import (
	"github.com/artarts36/potaynik/internal/port/http/kernel/responses"
	"github.com/rs/zerolog/log"
)

func handleNotFound(request Request) responses.Response {
	log.Debug().Msgf("Route with uri %s not found", request.URI())

	return responses.NotFound("Route not found")
}
