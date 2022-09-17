package routing

type Router struct {
	handlers map[string]map[string]Handler
}

func NewRouter() *Router {
	return &Router{handlers: map[string]map[string]Handler{}}
}

func (r *Router) Add(uri string, method string, handler Handler) *Router {
	if _, exists := r.handlers[uri]; !exists {
		r.handlers[uri] = map[string]Handler{}
	}

	r.handlers[uri][method] = handler
	r.addOptionsRoute(uri)

	return r
}

func (r *Router) Find(uri string, method string) Handler {
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
