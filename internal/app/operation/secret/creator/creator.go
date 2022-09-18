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

type SecretCreateParams struct {
	Value       string
	TTL         int
	AuthFactors map[string]interface{}
}

func NewCreator(secrets SecretRepository) *Creator {
	return &Creator{secrets: secrets, keyGenerator: &KeyGenerator{}}
}

func (c *Creator) Create(params SecretCreateParams) (*entity.Secret, error) {
	secretKey := c.keyGenerator.Generate()

	log.Info().Msgf("[SecretCreator] try to creating new secret with key %s", secretKey)

	secret := &entity.Secret{
		Key:   c.keyGenerator.Generate(),
		Value: params.Value,
		TTL:   params.TTL,
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

	log.Info().Msgf("[SecretCreator] secret with key %s was created", secretKey)

	return secret, nil
}
