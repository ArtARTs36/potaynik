package handlers

import (
	"encoding/json"

	"github.com/artarts36/potaynik/internal/app/operation/secret/creator"
	"github.com/artarts36/potaynik/internal/app/repository"
	"github.com/artarts36/potaynik/internal/port/http/kernel/responses"
	"github.com/artarts36/potaynik/internal/port/http/kernel/routing"
)

type SecretCreateHandler struct {
	creator *creator.Creator
}

type SecretCreateParams struct {
	Value       string                            `json:"value"`
	TTL         int                               `json:"ttl"`
	AuthFactors map[string]map[string]interface{} `json:"auth_factors"`
}

type SecretCreateResponse struct {
	Key string `json:"key"`
}

func NewSecretCreateHandler(creator *creator.Creator) *SecretCreateHandler {
	return &SecretCreateHandler{creator: creator}
}

func (h *SecretCreateHandler) Handle(r routing.Request) responses.Response {
	var params SecretCreateParams

	err := r.DecodeBody(&params)

	if err != nil || params.Value == "" {
		return responses.UnprocessableEntity("Invalid value")
	}

	if params.TTL == 0 {
		return responses.UnprocessableEntity("Invalid TTL")
	}

	sec, err := h.creator.Create(creator.SecretCreateParams{
		Value:       params.Value,
		TTL:         params.TTL,
		AuthFactors: params.AuthFactors,
	})

	if err != nil {
		_, isAlreadyExistsErr := err.(*repository.SecretAlreadyExistsError)

		if isAlreadyExistsErr {
			return responses.AlreadyReported(err.Error())
		}

		return responses.UnprocessableEntity(err.Error())
	}

	resp := SecretCreateResponse{
		Key: sec.Key,
	}

	respJSON, _ := json.Marshal(resp)

	return responses.Created(respJSON)
}
