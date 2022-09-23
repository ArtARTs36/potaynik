package handlers

import (
	"encoding/json"

	"github.com/artarts36/potaynik/internal/app/operation/health"
	"github.com/artarts36/potaynik/internal/port/http/kernel/responses"
	"github.com/artarts36/potaynik/internal/port/http/kernel/routing"
)

type HealthCheckHandler struct {
	checkers []health.Checker
}

type HealthCheckResponse struct {
	Status   bool                `json:"status"`
	Services []HealthCheckResult `json:"services"`
}

type HealthCheckResult struct {
	Status  bool   `json:"status"`
	Service string `json:"service"`
}

func NewHealthCheckHandler(checkers []health.Checker) *HealthCheckHandler {
	return &HealthCheckHandler{checkers: checkers}
}

func (h *HealthCheckHandler) Handle(r routing.Request) responses.Response {
	response := HealthCheckResponse{
		Status: true,
	}

	for _, check := range health.RunHealthChecks(r.Context(), h.checkers) {
		response.Status = response.Status && check.Status

		response.Services = append(response.Services, HealthCheckResult{
			Service: check.Service,
			Status:  check.Status,
		})
	}

	respJSON, err := json.Marshal(response)

	if err != nil {
		return responses.ServerError(err.Error())
	}

	if response.Status {
		return responses.OK(respJSON)
	}

	return responses.ServerErrorFromBytes(respJSON)
}
