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

func (v *Viewer) View(secretKey string, authFactors map[string]string) (string, error) {
	v.metrics.IncViewTotalAttempts()

	log.Info().Msgf("[SecretViewer] finding secret with key %s", secretKey)

	secret, err := v.secrets.Find(secretKey)

	if err != nil {
		v.metrics.IncSearchFails()
		log.Error().Msgf("[SecretViewer] fail on finding secret with key %s: %s", secretKey, err.Error())

		return "", err
	}

	if secret == nil {
		v.metrics.IncSecretNotFound()
		log.Info().Msgf("[SecretViewer] secret with key %s not found", secretKey)

		return "", newSecretNotFoundError(secretKey)
	}

	v.metrics.IncSecretFound()

	log.Info().Msgf("[SecretViewer] secret with key %s found", secretKey)

	access := v.authorize(secret, authFactors)

	if !access.Access {
		log.Info().Msgf("[SecretViewer] user cant show secret with key %s", secretKey)

		return "", newSecretViewForbiddenError(secretKey, access.Reason)
	}

	log.Info().Msgf("[SecretViewer] user can show secret with key %s", secretKey)

	v.secrets.Delete(secretKey)

	log.Info().Msgf("[SecretViewer] secret with key %s deleted", secretKey)

	return secret.Value, nil
}

func (v *Viewer) authorize(secret *entity.Secret, authFactors map[string]string) auth.Access {
	for factorKey, factor := range secret.AuthFactors {
		authorizer := v.authorizers[factorKey]

		access, err := authorizer.Authorize(factor, auth.AuthorizeRequest{
			UserFactorValue: authFactors[factorKey],
		})

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
