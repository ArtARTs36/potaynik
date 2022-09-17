package routing

import "net/http"

type HttpHandler func(http.ResponseWriter, *http.Request)

type Router struct {
	handlers map[string]map[string]HttpHandler
}

func NewRouter() *Router {
	return &Router{handlers: map[string]map[string]HttpHandler{}}
}

func (r *Router) Add(uri string, method string, handler HttpHandler) *Router {
	if _, exists := r.handlers[uri]; !exists {
		r.handlers[uri] = map[string]HttpHandler{}
	}

	r.handlers[uri][method] = handler
	r.addOptionsRoute(uri)

	return r
}

func (r *Router) Find(uri string, method string) HttpHandler {
	uriHandlers, exists := r.handlers[uri]

	if !exists {
		return handleNotFound
	}

	handler, exists := uriHandlers[method]

	if exists {
		return handler
	}

	methodNotAllowedHandler := newMethodNotAllowedHandler(uriHandlers)

	return methodNotAllowedHandler.Handle
}

func (r *Router) addOptionsRoute(uri string) {
	if _, exists := r.handlers[uri]["OPTIONS"]; exists {
		return
	}

	uriRoutes := r.handlers[uri]

	optsHandler := &OptionsHandler{routes: &uriRoutes}

	r.handlers[uri]["OPTIONS"] = optsHandler.Handle
}
