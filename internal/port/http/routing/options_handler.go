package routing

import (
	"net/http"
	"strings"
)

type OptionsHandler struct {
	routes      *map[string]HttpHandler
	headerValue *string
}

func (h *OptionsHandler) Handle(w http.ResponseWriter, _ *http.Request) {
	h.retrieveHeaderValue()

	w.Header().Set("Allow", *h.headerValue)
	w.WriteHeader(http.StatusNoContent)
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
