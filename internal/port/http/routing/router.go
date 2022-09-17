package routing

import "net/http"

type HttpHandler func(http.ResponseWriter, *http.Request)

type Router struct {
	routes map[string]map[string]HttpHandler
}

func NewRouter() *Router {
	return &Router{routes: map[string]map[string]HttpHandler{}}
}

func (r *Router) Add(uri string, method string, handler HttpHandler) *Router {
	if _, exists := r.routes[uri]; !exists {
		r.routes[uri] = map[string]HttpHandler{}
	}

	r.routes[uri][method] = handler
	r.addOptionsRoute(uri)

	return r
}

func (r *Router) Find(uri string, method string) HttpHandler {
	uriRoutes, exists := r.routes[uri]

	if !exists {
		return handleNotFound
	}

	route, exists := uriRoutes[method]

	if exists {
		return route
	}

	handler := MethodNotAllowedHandler{
		routes: &uriRoutes,
	}

	return handler.Handle
}

func (r *Router) addOptionsRoute(uri string) {
	if _, exists := r.routes[uri]["OPTIONS"]; exists {
		return
	}

	uriRoutes := r.routes[uri]

	optsHandler := &OptionsHandler{routes: &uriRoutes}

	r.routes[uri]["OPTIONS"] = optsHandler.Handle
}
