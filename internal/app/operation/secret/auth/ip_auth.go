package auth

import (
	"fmt"

	"github.com/artarts36/potaynik/internal/app/entity"
)

const ParamsIPKey = "ip"

type IPAuthorizer struct {
}

func (auth *IPAuthorizer) CreateAuthFactor(params CreateAuthFactorParams) (entity.AuthFactor, error) {
	ip, exists := params[ParamsIPKey]

	if !exists {
		return entity.AuthFactor{}, fmt.Errorf("ip not defined")
	}

	ipString, ipValid := ip.(string)

	if ipValid {
		return entity.AuthFactor{}, fmt.Errorf("ip isn't valid")
	}

	ip, err := newIP(ipString)

	if err != nil {
		return entity.AuthFactor{}, fmt.Errorf("ip isn't valid")
	}

	return entity.AuthFactor{
		Key: ParamsIPKey,
		Params: map[string]interface{}{
			ParamsIPKey: ip,
		},
	}, nil
}

//nolint:unparam
func (auth *IPAuthorizer) Authorize(factor entity.AuthFactor, req AuthorizeRequest) (Access, error) {
	ipVal, exists := factor.Params[ParamsIPKey]

	if !exists {
		return Access{
			Access: false,
			Reason: "ip not defined",
		}, nil
	}

	ip, ipValid := ipVal.(*IP)

	if ipValid {
		return Access{
			Access: false,
			Reason: "ip not defined",
		}, nil
	}

	return Access{
		Access: ip.Contains(req.User.IPAddress),
		Reason: "user ip",
	}, nil
}
