package handlers

import (
	"encoding/json"
	"github.com/artarts36/potaynik/internal/app/operation/secret/viewer"
	"github.com/artarts36/potaynik/internal/port/http/kernel/responses"
	"github.com/artarts36/potaynik/internal/port/http/kernel/routing"
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

func (h *SecretShowHandler) Handle(r routing.Request) responses.Response {
	var params SecretShowParams

	err := r.DecodeQuery(&params)

	if err != nil {
		return responses.UnprocessableEntity("invalid data")
	}

	if params.SecretKey == "" {
		return responses.UnprocessableEntity("invalid secret key")
	}

	val, err := h.viewer.View(params.SecretKey, params.AuthFactors)

	notFoundErr, isNotFoundErr := err.(*viewer.SecretNotFoundError)

	if isNotFoundErr {
		return responses.NotFound(notFoundErr.Error())
	}

	forbiddenErr, isForbiddenErr := err.(*viewer.SecretViewForbiddenError)

	if isForbiddenErr {
		return responses.Forbidden(forbiddenErr.Reason)
	}

	resp, _ := json.Marshal(&SecretShowResponse{
		Value: val,
	})

	return responses.OK(resp)
}
