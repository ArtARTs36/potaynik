package routing

import (
	"net/http"
	"strings"
)

type OptionsHandler struct {
	routes      *map[string]GoHttpHandler
	headerValue *string
}

func (h *OptionsHandler) Handle(_ Request) Response {
	h.retrieveHeaderValue()

	return Response{
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
