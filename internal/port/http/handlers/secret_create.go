package handlers

import (
	"encoding/json"
	"github.com/artarts36/potaynik/internal/app/operation/secret/creator"
	"github.com/rs/zerolog/log"
	"net/http"
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

func (h *SecretCreateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var params SecretCreateParams

	log.Ctx(r.Context())

	json.NewDecoder(r.Body).Decode(&params)

	if params.Value == "" {
		w.WriteHeader(422)

		return
	}

	sec, _ := h.creator.Create(params.Value)

	resp := SecretCreateResponse{
		Key: sec.Key,
	}

	respJSON, _ := json.Marshal(resp)

	w.Write(respJSON)
	w.WriteHeader(200)
}
