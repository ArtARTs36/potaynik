package routing

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Controller struct {
	router *Router
}

func NewController(callback func(router *Router)) *Controller {
	router := NewRouter()

	callback(router)

	return NewControllerWithRouter(router)
}

func NewControllerWithRouter(router *Router) *Controller {
	return &Controller{router: router}
}

func (c *Controller) HandleRequest(writer http.ResponseWriter, request *http.Request) {
	var rootLogCtx zerolog.Context

	log.Logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		rootLogCtx = c

		return c.Str("user_request_id", uuid.New().String())
	})

	c.router.Find(request.URL.Path, request.Method)(writer, request)

	log.Logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return rootLogCtx
	})
}

func (c *Controller) Serve(port int) error {
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.HandleFunc("/", c.HandleRequest)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	chErr := make(chan error)

	go func() {
		log.Info().Msgf("Start listening on '%s'", server.Addr)

		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Info().Msg("http server closed")

			return
		}

		chErr <- err
	}()

	select {
	case err := <-chErr:
		return err
	case <-ctx.Done():
		return server.Shutdown(ctx)
	}
}
