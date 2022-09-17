package handlers

import (
	"encoding/json"
	"github.com/artarts36/potaynik/internal/app/operation/secret/viewer"
	"net/http"
)

type SecretShowHandler struct {
	viewer *viewer.Viewer
}

type SecretShowParams struct {
	SecretKey string `json:"secret_key"`
}

type SecretShowResponse struct {
	Value string `json:"value"`
}

func NewSecretShowHandler(viewer *viewer.Viewer) *SecretShowHandler {
	return &SecretShowHandler{viewer}
}

func (h *SecretShowHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var params SecretShowParams

	json.NewDecoder(r.Body).Decode(&params)

	val, err := h.viewer.View(params.SecretKey)

	notFoundErr, isNotFoundErr := err.(*viewer.SecretNotFoundError)

	if isNotFoundErr {
		resp, _ := json.Marshal(&ErrorResponse{
			Error: notFoundErr.Error(),
		})

		w.WriteHeader(404)
		w.Write(resp)

		return
	}

	if val != nil {
		resp, _ := json.Marshal(&SecretShowResponse{
			Value: *val,
		})

		w.Write(resp)
	}
}
