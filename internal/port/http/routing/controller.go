package routing

import (
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
	handler := c.router.Find(request.RequestURI, request.Method)

	if handler != nil {
		handleFunc := *handler

		handleFunc(writer, request)

		return
	}

	log.Debug().Msgf("Route with uri %s not found", request.RequestURI)

	writer.WriteHeader(404)
}

func (c *Controller) Serve() {
	http.HandleFunc("/", c.HandleRequest)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Error().Msg(err.Error())
	}
}
