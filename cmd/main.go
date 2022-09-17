package main

import (
	"github.com/artarts36/potaynik/internal/app"
	"net/http"
)

func main() {
	application := app.NewApplication()

	http.HandleFunc("/api/createSecret", application.Services.Http.Handlers.SecretCreateHandler.Handle)
	http.HandleFunc("/api/showSecret", application.Services.Http.Handlers.SecretShowHandler.Handle)
	http.ListenAndServe(":8080", nil)
}
