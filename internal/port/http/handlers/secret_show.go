package handlers

import (
	"encoding/json"
	"github.com/artarts36/potaynik/internal/app/operation/secret/viewer"
	"github.com/artarts36/potaynik/internal/port/http/routing"
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

func (h *SecretShowHandler) Handle(r routing.Request) routing.Response {
	var params SecretShowParams

	json.NewDecoder(r.Request.Body).Decode(&params)

	val, err := h.viewer.View(params.SecretKey)

	notFoundErr, isNotFoundErr := err.(*viewer.SecretNotFoundError)

	if isNotFoundErr {
		return routing.NewNotFoundResponse(notFoundErr.Error())
	}

	resp, _ := json.Marshal(&SecretShowResponse{
		Value: val,
	})

	return routing.NewOKResponse(resp)
}
