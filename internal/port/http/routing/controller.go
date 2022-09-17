package routing

import (
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

	c.router.Find(request.RequestURI, request.Method)(writer, request)

	log.Logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return rootLogCtx
	})
}

func (c *Controller) Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", c.HandleRequest)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Error().Msg(err.Error())
	}
}
