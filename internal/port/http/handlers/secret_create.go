package handlers

import (
	"encoding/json"
	"github.com/artarts36/potaynik/internal/app/operation/secret/creator"
	"github.com/artarts36/potaynik/internal/app/repository"
	"github.com/artarts36/potaynik/internal/port/http/routing"
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

func (h *SecretCreateHandler) Handle(r routing.Request) routing.Response {
	var params SecretCreateParams

	err := r.DecodeBody(&params)

	if err != nil || params.Value == "" {
		return routing.NewInvalidEntityResponse("Invalid value")
	}

	if params.TTL == 0 {
		return routing.NewInvalidEntityResponse("Invalid TTL")
	}

	sec, err := h.creator.Create(creator.SecretCreateParams{
		Value:       params.Value,
		TTL:         params.TTL,
		AuthFactors: params.AuthFactors,
	})

	if err != nil {
		_, isAlreadyExistsErr := err.(*repository.SecretAlreadyExistsError)

		if isAlreadyExistsErr {
			return routing.NewAlreadyReportedResponse(err.Error())
		}

		return routing.NewInvalidEntityResponse(err.Error())
	}

	resp := SecretCreateResponse{
		Key: sec.Key,
	}

	respJSON, _ := json.Marshal(resp)

	return routing.NewCreatedResponse(respJSON)
}
