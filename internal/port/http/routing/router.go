package routing

type Router struct {
	handlers map[string]map[string]GoHttpHandler
}

func NewRouter() *Router {
	return &Router{handlers: map[string]map[string]GoHttpHandler{}}
}

func (r *Router) Add(uri string, method string, handler Handler) *Router {
	if _, exists := r.handlers[uri]; !exists {
		r.handlers[uri] = map[string]GoHttpHandler{}
	}

	r.handlers[uri][method] = WrapHandler(handler)
	r.addOptionsRoute(uri)

	return r
}

func (r *Router) Find(uri string, method string) GoHttpHandler {
	uriHandlers, exists := r.handlers[uri]

	if !exists {
		return WrapHandler(handleNotFound)
	}

	handler, exists := uriHandlers[method]

	if exists {
		return handler
	}

	methodNotAllowedHandler := newMethodNotAllowedHandler(uriHandlers)

	return WrapHandler(methodNotAllowedHandler.Handle)
}

func (r *Router) addOptionsRoute(uri string) {
	if _, exists := r.handlers[uri]["OPTIONS"]; exists {
		return
	}

	uriRoutes := r.handlers[uri]

	optsHandler := &OptionsHandler{routes: &uriRoutes}

	r.handlers[uri]["OPTIONS"] = WrapHandler(optsHandler.Handle)
}
