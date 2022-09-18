package auth

import "github.com/artarts36/potaynik/internal/app/entity"

type Authorizer interface {
	CreateAuthFactor(params CreateAuthFactorParams) (entity.AuthFactor, error)
	Authorize(factor entity.AuthFactor, request AuthorizeRequest) (Access, error)
}

type Access struct {
	Access bool
	Reason string
}

type CreateAuthFactorParams map[string]interface{}

type AuthorizeRequest struct {
	UserFactorValue string
}
