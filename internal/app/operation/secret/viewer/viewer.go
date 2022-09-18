package viewer

import (
	"github.com/artarts36/potaynik/internal/app/entity"
	"github.com/rs/zerolog/log"
)

type Viewer struct {
	secrets Repository
}

type Repository interface {
	Find(key string) (*entity.Secret, error)
	Delete(key string)
}

func New(repository Repository) *Viewer {
	return &Viewer{secrets: repository}
}

func (v *Viewer) View(secretKey string) (string, error) {
	log.Info().Msgf("[SecretViewer] finding secret with key %s", secretKey)

	secret, err := v.secrets.Find(secretKey)

	if err != nil {
		log.Error().Msgf("[SecretViewer] fail on finding secret with key %s: %s", secretKey, err.Error())

		return "", err
	}

	if secret == nil {
		log.Info().Msgf("[SecretViewer] secret with key %s not found", secretKey)

		return "", newSecretNotFoundError(secretKey)
	}

	log.Info().Msgf("[SecretViewer] finding secret with key %s found", secretKey)

	v.secrets.Delete(secretKey)

	log.Info().Msgf("[SecretViewer] finding secret with key %s deleted", secretKey)

	return secret.Value, nil
}
