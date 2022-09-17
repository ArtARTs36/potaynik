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
	log.Logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("user_request_id", uuid.New().String())
	})

	c.router.Find(request.RequestURI, request.Method)(writer, request)
}

func (c *Controller) Serve() {
	http.HandleFunc("/", c.HandleRequest)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Error().Msg(err.Error())
	}
}
