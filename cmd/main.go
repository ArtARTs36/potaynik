package main

import (
	"fmt"
	"sync"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"

	"github.com/artarts36/potaynik/internal/app"
	"github.com/artarts36/potaynik/internal/port/http/kernel/routing"
)

type serverRunner func(application *app.Application) error

func main() {
	application, err := app.NewApplication("potaynik")

	if err != nil {
		panic(fmt.Sprintf("cant build application: %s", err))
	}

	servers := []serverRunner{
		runApplicationServer,
		runHealthServer,
	}

	wg := new(sync.WaitGroup)

	for _, server := range servers {
		wg.Add(1)

		server := server

		go func() {
			err := server(application)

			if err != nil {
				log.Error().Msg(err.Error())
			}
		}()
	}

	wg.Wait()
}

func runApplicationServer(application *app.Application) error {
	return routing.NewController(func(router *routing.Router) {
		router.
			AddAppHandler("/api/secrets/add", "POST", application.Services.HTTP.Handlers.SecretCreateHandler.Handle).
			AddAppHandler("/api/secrets/show", "GET", application.Services.HTTP.Handlers.SecretShowHandler.Handle).
			AddAppHandler("/api/secrets/info", "GET", application.Services.HTTP.Handlers.SecretInfoHandler.Handle)
	}).
		Serve(application.Environment.HTTP.Public.Port)
}

func runHealthServer(application *app.Application) error {
	return routing.NewController(func(router *routing.Router) {
		router.
			AddGoHandler("/metrics", "GET", promhttp.HandlerFor(
				application.Metrics.Registry,
				promhttp.HandlerOpts{
					EnableOpenMetrics: true,
				}).ServeHTTP).
			AddAppHandler("/health/check", "GET", application.Services.HTTP.Handlers.HealthCheckHandler.Handle)
	}).
		Serve(application.Environment.HTTP.Health.Port)
}
