package routing

import (
	"github.com/artarts36/potaynik/internal/port/http/kernel/responses"
	"net/http"
)

type AppHandler func(Request) responses.Response

type GoHttpHandler func(writer http.ResponseWriter, request *http.Request)
