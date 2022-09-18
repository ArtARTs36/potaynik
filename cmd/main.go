package main

import (
	"fmt"
	"github.com/artarts36/potaynik/internal/app"
	"github.com/artarts36/potaynik/internal/port/http/routing"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"sync"
)

func main() {
	application, err := app.NewApplication("potaynik")

	if err != nil {
		panic(fmt.Sprintf("cant build application: %s", err))
	}

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		err = routing.
			NewController(func(router *routing.Router) {
				router.
					Add("/api/secrets", "POST", application.Services.Http.Handlers.SecretCreateHandler.Handle).
					Add("/api/secrets", "GET", application.Services.Http.Handlers.SecretShowHandler.Handle)
			}).
			Serve(application.Environment.Http.Public.Port)

		if err != nil {
			log.Error().Msg(err.Error())
		}

		wg.Done()
	}()

	go func() {
		err = routing.
			NewController(func(router *routing.Router) {
				router.
					AddGoHandler("/metrics", "GET", promhttp.HandlerFor(
						application.Metrics.Registry,
						promhttp.HandlerOpts{
							EnableOpenMetrics: true,
						}).ServeHTTP)
			}).
			Serve(application.Environment.Http.Health.Port)

		if err != nil {
			log.Error().Msg(err.Error())
		}

		wg.Done()
	}()

	wg.Wait()
}
