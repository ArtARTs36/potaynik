package routing

import (
	"github.com/artarts36/potaynik/internal/port/http/kernel/responses"
	"net/http"
	"strings"
)

type MethodNotAllowedHandler struct {
	routes map[string]GoHttpHandler
}

func newMethodNotAllowedHandler(routes map[string]GoHttpHandler) MethodNotAllowedHandler {
	return MethodNotAllowedHandler{routes: routes}
}

func (h *MethodNotAllowedHandler) Handle(_ Request) responses.Response {
	return responses.Response{
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
