package routing

import (
	"net/http"

	"github.com/artarts36/potaynik/internal/port/http/kernel/responses"
)

type AppHandler func(Request) responses.Response

type GoHTTPHandler func(writer http.ResponseWriter, request *http.Request)
