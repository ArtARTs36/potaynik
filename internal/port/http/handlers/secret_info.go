package handlers

import (
	"encoding/json"

	"github.com/rs/zerolog/log"

	"github.com/artarts36/potaynik/internal/app/operation/secret/informer"
	"github.com/artarts36/potaynik/internal/port/http/kernel/responses"
	"github.com/artarts36/potaynik/internal/port/http/kernel/routing"
)

type SecretInfoHandler struct {
	informer *informer.Informer
}

type SecretInfoParams struct {
	SecretKey string `query:"secret_key"`
}

func NewSecretInfoHandler(informer *informer.Informer) *SecretInfoHandler {
	return &SecretInfoHandler{informer: informer}
}

func (h *SecretInfoHandler) Handle(r routing.Request) responses.Response {
	var params SecretInfoParams

	err := r.DecodeQuery(&params)

	if err != nil || params.SecretKey == "" {
		return responses.UnprocessableEntity("invalid secret key")
	}

	info, err := h.informer.Info(params.SecretKey)

	if err == nil {
		infoJSON, _ := json.Marshal(info)

		return responses.OK(infoJSON)
	}

	notFoundErr, isNotFoundErr := err.(*informer.SecretNotFoundError)

	if isNotFoundErr {
		return responses.NotFound(notFoundErr.Error())
	}

	log.Error().Msgf("[SecretInfoHandler] informer returns error: %s", err)

	return responses.ServerError("")
}
