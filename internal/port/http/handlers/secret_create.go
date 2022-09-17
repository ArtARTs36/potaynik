package handlers

import (
	"encoding/json"
	"github.com/artarts36/potaynik/internal/app/operation/secret/creator"
	"github.com/artarts36/potaynik/internal/port/http/routing"
)

type SecretCreateHandler struct {
	creator *creator.Creator
}

type SecretCreateParams struct {
	Value string `json:"value"`
}

type SecretCreateResponse struct {
	Key string `json:"key"`
}

func NewSecretCreateHandler(creator *creator.Creator) *SecretCreateHandler {
	return &SecretCreateHandler{creator: creator}
}

func (h *SecretCreateHandler) Handle(r routing.Request) routing.Response {
	var params SecretCreateParams

	json.NewDecoder(r.Request.Body).Decode(&params)

	if params.Value == "" {
		return routing.NewInvalidEntityResponse("Invalid value")
	}

	sec, _ := h.creator.Create(params.Value)

	resp := SecretCreateResponse{
		Key: sec.Key,
	}

	respJSON, _ := json.Marshal(resp)

	return routing.NewCreatedResponse(respJSON)
}
