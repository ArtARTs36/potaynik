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
	SecretKey   string            `query:"secret_key"`
	AuthFactors map[string]string `query:"auth_factors"`
}

type SecretShowResponse struct {
	Value string `json:"value"`
}

func NewSecretShowHandler(viewer *viewer.Viewer) *SecretShowHandler {
	return &SecretShowHandler{viewer}
}

func (h *SecretShowHandler) Handle(r routing.Request) routing.Response {
	var params SecretShowParams

	err := r.DecodeQuery(&params)

	if err != nil {
		return routing.NewInvalidEntityResponse("invalid data")
	}

	if params.SecretKey == "" {
		return routing.NewInvalidEntityResponse("invalid secret key")
	}

	val, err := h.viewer.View(params.SecretKey, params.AuthFactors)

	notFoundErr, isNotFoundErr := err.(*viewer.SecretNotFoundError)

	if isNotFoundErr {
		return routing.NewNotFoundResponse(notFoundErr.Error())
	}

	forbiddenErr, isForbiddenErr := err.(*viewer.SecretViewForbiddenError)

	if isForbiddenErr {
		return routing.NewForbiddenResponseWithText(forbiddenErr.Reason)
	}

	resp, _ := json.Marshal(&SecretShowResponse{
		Value: val,
	})

	return routing.NewOKResponse(resp)
}
