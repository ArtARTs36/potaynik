package routing

import (
	"net/http"
	"strings"

	"github.com/artarts36/potaynik/internal/port/http/kernel/responses"
)

type MethodNotAllowedHandler struct {
	routes map[string]GoHTTPHandler
}

func newMethodNotAllowedHandler(routes map[string]GoHTTPHandler) MethodNotAllowedHandler {
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

	for method := range h.routes {
		methods = append(methods, method)
	}

	return strings.Join(methods, ", ")
}
