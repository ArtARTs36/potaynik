package routing

import (
	"net/http"
	"strings"
)

type MethodNotAllowedHandler struct {
	routes map[string]Handler
}

func newMethodNotAllowedHandler(routes map[string]Handler) MethodNotAllowedHandler {
	return MethodNotAllowedHandler{routes: routes}
}

func (h *MethodNotAllowedHandler) Handle(_ Request) Response {
	return Response{
		Code: http.StatusMethodNotAllowed,
		Headers: map[string]string{
			"Allow": h.buildHeaderValue(),
		},
	}
}

func (h *MethodNotAllowedHandler) buildHeaderValue() string {
	methods := make([]string, 0, len(h.routes))

	for method, _ := range h.routes {
		methods = append(methods, method)
	}

	return strings.Join(methods, ", ")
}
