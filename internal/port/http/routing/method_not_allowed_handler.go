package routing

import (
	"net/http"
	"strings"
)

type MethodNotAllowedHandler struct {
	routes map[string]HttpHandler
}

func newMethodNotAllowedHandler(routes map[string]HttpHandler) MethodNotAllowedHandler {
	return MethodNotAllowedHandler{routes: routes}
}

func (h *MethodNotAllowedHandler) Handle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Allow", h.buildHeaderValue())
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *MethodNotAllowedHandler) buildHeaderValue() string {
	methods := make([]string, 0, len(h.routes))

	for method, _ := range h.routes {
		methods = append(methods, method)
	}

	return strings.Join(methods, ", ")
}
