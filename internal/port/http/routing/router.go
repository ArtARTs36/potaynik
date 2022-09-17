package routing

import "net/http"

type httpHandler func(http.ResponseWriter, *http.Request)

type Router struct {
	routes map[string]map[string]httpHandler
}

func NewRouter() *Router {
	return &Router{routes: map[string]map[string]httpHandler{}}
}

func (r *Router) Add(uri string, method string, handler httpHandler) *Router {
	if _, exists := r.routes[uri]; !exists {
		r.routes[uri] = map[string]httpHandler{}
	}

	r.routes[uri][method] = handler

	return r
}

func (r *Router) Find(uri string, method string) *httpHandler {
	route, exists := r.routes[uri][method]

	if exists {
		return &route
	}

	return nil
}
