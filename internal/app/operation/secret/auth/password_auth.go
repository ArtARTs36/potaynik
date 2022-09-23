package auth

import (
	"fmt"

	"github.com/artarts36/potaynik/internal/app/entity"
)

const PasswordAuthorizerKey = "password"
const ParamsPasswordKey = "password"

type PasswordAuthorizer struct {
}

func (auth *PasswordAuthorizer) CreateAuthFactor(params CreateAuthFactorParams) (entity.AuthFactor, error) {
	pswd, exists := params[ParamsPasswordKey]

	if !exists {
		return entity.AuthFactor{}, fmt.Errorf("password not defined")
	}

	return entity.AuthFactor{
		Key: PasswordAuthorizerKey,
		Params: map[string]interface{}{
			ParamsPasswordKey: pswd,
		},
	}, nil
}

//nolint:unparam
func (auth *PasswordAuthorizer) Authorize(factor entity.AuthFactor, req AuthorizeRequest) (Access, error) {
	pswd, exists := factor.Params[ParamsPasswordKey]

	if !exists {
		return Access{
			Access: false,
			Reason: "password not defined",
		}, nil
	}

	userPswd := req.UserFactorValue

	if !exists {
		return Access{
			Access: false,
			Reason: "password not given",
		}, nil
	}

	if pswd != userPswd {
		return Access{
			Access: false,
			Reason: "passwords not equals",
		}, nil
	}

	return Access{
		Access: true,
		Reason: "passwords equals",
	}, nil
}
