package routing

type Router struct {
	handlers map[string]map[string]GoHTTPHandler
}

func NewRouter() *Router {
	return &Router{handlers: map[string]map[string]GoHTTPHandler{}}
}

func (r *Router) AddGoHandler(uri string, method string, handler GoHTTPHandler) *Router {
	if _, exists := r.handlers[uri]; !exists {
		r.handlers[uri] = map[string]GoHTTPHandler{}
	}

	r.handlers[uri][method] = handler
	r.addOptionsRoute(uri)

	return r
}

func (r *Router) AddAppHandler(uri string, method string, handler AppHandler) *Router {
	return r.AddGoHandler(uri, method, WrapHandler(handler))
}

func (r *Router) Find(uri string, method string) GoHTTPHandler {
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
