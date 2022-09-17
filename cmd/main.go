package main

import (
	"github.com/artarts36/potaynik/internal/app"
	"github.com/artarts36/potaynik/internal/port/http/routing"
)

func main() {
	application := app.NewApplication()

	routing.
		NewController(func(router *routing.Router) {
			router.
				Add("/api/secrets", "POST", application.Services.Http.Handlers.SecretCreateHandler.Handle).
				Add("/api/secrets", "GET", application.Services.Http.Handlers.SecretShowHandler.Handle)
		}).
		Serve()
}
