package routing

import (
	"net/http"
)

type Handler func(Request) Response

type GoHttpHandler func(writer http.ResponseWriter, request *http.Request)
