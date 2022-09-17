package creator

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/artarts36/potaynik/internal/app/entity"
)

type Creator struct {
	keyGenerator *KeyGenerator
	secrets      SecretRepository
}

func NewCreator(secrets SecretRepository) *Creator {
	return &Creator{secrets: secrets, keyGenerator: &KeyGenerator{}}
}

type SecretRepository interface {
	Add(secret *entity.Secret) error
}

func (c *Creator) Create(secretVal string) (*entity.Secret, error) {
	secretKey := c.keyGenerator.Generate()

	log.Info().Msgf("[SecretCreator] try to creating new secret with key %s", secretKey)

	secret := &entity.Secret{
		Key:   c.keyGenerator.Generate(),
		Value: secretVal,
	}

	err := c.secrets.Add(secret)

	if err != nil {
		fmt.Printf(
			"[SecretCreator] cannot create secret: %v",
			err,
		)

		log.Error().Msgf("[SecretCreator] try to creating new secret with key %s", secretKey)

		return &entity.Secret{}, err
	}

	return secret, nil
}
