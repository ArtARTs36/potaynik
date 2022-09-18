package routing

import (
	"github.com/artarts36/potaynik/internal/port/http/kernel/responses"
	"net/http"
	"strings"
)

type OptionsHandler struct {
	routes      *map[string]GoHttpHandler
	headerValue *string
}

func (h *OptionsHandler) Handle(_ Request) responses.Response {
	h.retrieveHeaderValue()

	return responses.Response{
		Code: http.StatusNoContent,
		Headers: map[string]string{
			"Allow": *h.headerValue,
		},
	}
}

func (h *OptionsHandler) retrieveHeaderValue() {
	if h.headerValue != nil {
		return
	}

	methods := make([]string, 0, len(*h.routes))

	for method, _ := range *h.routes {
		methods = append(methods, method)
	}

	val := strings.Join(methods, ", ")

	h.headerValue = &val
}
