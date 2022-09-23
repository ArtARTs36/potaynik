package viewer

import (
	"github.com/rs/zerolog/log"

	"github.com/artarts36/potaynik/internal/app/entity"
	"github.com/artarts36/potaynik/internal/app/operation/secret/auth"
)

type Viewer struct {
	secrets     Repository
	authorizers map[string]auth.Authorizer
	metrics     *Metrics
}

type Repository interface {
	Find(key string) (*entity.Secret, error)
	Delete(key string)
}

func New(repository Repository, authorizers map[string]auth.Authorizer, metrics *Metrics) *Viewer {
	return &Viewer{secrets: repository, authorizers: authorizers, metrics: metrics}
}

func (v *Viewer) View(params ViewParams) (string, error) {
	v.metrics.IncViewTotalAttempts()

	log.Info().Msgf("[SecretViewer] finding secret with key %s", params.SecretKey)

	secret, err := v.secrets.Find(params.SecretKey)

	if err != nil {
		v.metrics.IncSearchFails()
		log.Error().Msgf(
			"[SecretViewer] fail on finding secret with key %s: %s",
			params.SecretKey,
			err.Error(),
		)

		return "", err
	}

	if secret == nil {
		v.metrics.IncSecretNotFound()
		log.Info().Msgf("[SecretViewer] secret with key %s not found", params.SecretKey)

		return "", newSecretNotFoundError(params.SecretKey)
	}

	v.metrics.IncSecretFound()

	log.Info().Msgf("[SecretViewer] secret with key %s found", params.SecretKey)

	access := v.authorize(secret, params)

	if !access.Access {
		v.metrics.IncAccessFail()

		log.Info().Msgf("[SecretViewer] user cant show secret with key %s", params.SecretKey)

		return "", newSecretViewForbiddenError(params.SecretKey, access.Reason)
	}

	v.metrics.IncAccessFail()

	log.Info().Msgf("[SecretViewer] user can show secret with key %s", params.SecretKey)

	v.secrets.Delete(params.SecretKey)

	log.Info().Msgf("[SecretViewer] secret with key %s deleted", params.SecretKey)

	return secret.Value, nil
}

func (v *Viewer) authorize(secret *entity.Secret, params ViewParams) auth.Access {
	for factorKey, factor := range secret.AuthFactors {
		authorizer := v.authorizers[factorKey]

		access, err := authorizer.Authorize(factor, auth.NewAuthorizeRequest(
			params.AuthFactors[factorKey],
			auth.NewUser(params.UserIP),
		))

		if err != nil {
			log.Error().Msgf("[SecretViewer] Authorizer %s returns error: %s", factorKey, err.Error())
		}

		if !access.Access {
			v.metrics.IncAuthPassFail(factorKey)

			return access
		}

		v.metrics.IncAuthPassOk(factorKey)
	}

	return auth.Access{
		Access: true,
		Reason: "OK",
	}
}
