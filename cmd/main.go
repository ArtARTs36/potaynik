package main

import (
	"github.com/artarts36/potaynik/internal/app"
	"github.com/artarts36/potaynik/internal/port/http/routing"
	"github.com/rs/zerolog/log"
)

func main() {
	application := app.NewApplication()

	err := routing.
		NewController(func(router *routing.Router) {
			router.
				Add("/api/secrets", "POST", application.Services.Http.Handlers.SecretCreateHandler.Handle).
				Add("/api/secrets", "GET", application.Services.Http.Handlers.SecretShowHandler.Handle)
		}).
		Serve()

	if err != nil {
		log.Error().Msg(err.Error())
	}
}
