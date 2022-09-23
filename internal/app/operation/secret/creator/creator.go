package creator

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/artarts36/potaynik/internal/app/entity"
	"github.com/artarts36/potaynik/internal/app/operation/secret/auth"
)

type Creator struct {
	keyGenerator *KeyGenerator
	secrets      SecretRepository
	authorizers  map[string]auth.Authorizer
	metrics      *Metrics
}

type SecretCreateParams struct {
	Value       string
	TTL         int
	AuthFactors map[string]map[string]interface{}
}

func NewCreator(secrets SecretRepository, authorizers map[string]auth.Authorizer, metrics *Metrics) *Creator {
	return &Creator{secrets: secrets, keyGenerator: &KeyGenerator{}, authorizers: authorizers, metrics: metrics}
}

func (c *Creator) Create(params SecretCreateParams) (*entity.Secret, error) {
	c.metrics.IncSecretCreateAttempts()

	secretKey := c.keyGenerator.Generate()

	log.Info().Msgf("[SecretCreator] try to creating new secret with key %s", secretKey)

	authFactors, err := c.createAuthFactors(params)

	if err != nil {
		log.Info().Msg(err.Error())

		return nil, err
	}

	log.Info().Msgf(
		"[SecretCreator] attached %d auth factors to secret with key %s",
		len(authFactors),
		secretKey,
	)

	secret := &entity.Secret{
		Key:         c.keyGenerator.Generate(),
		Value:       params.Value,
		TTL:         params.TTL,
		AuthFactors: authFactors,
	}

	err = c.secrets.Add(secret)

	if err != nil {
		fmt.Printf(
			"[SecretCreator] cannot create secret: %v",
			err,
		)

		log.Error().Msgf("[SecretCreator] try to creating new secret with key %s", secretKey)

		return &entity.Secret{}, err
	}

	log.Info().Msgf("[SecretCreator] secret with key %s was created", secretKey)

	c.metrics.IncSecretCreateSuccessAttempts()

	return secret, nil
}

func (c *Creator) createAuthFactors(params SecretCreateParams) (map[string]entity.AuthFactor, error) {
	factors := map[string]entity.AuthFactor{}

	for authorizerKey, factorData := range params.AuthFactors {
		authorizer, authorizerExists := c.authorizers[authorizerKey]

		if !authorizerExists {
			return factors, fmt.Errorf("[SecretCreator] authorizer with key %s not found", authorizer)
		}

		c.metrics.IncUseAuthFactor(authorizerKey)

		factor, err := authorizer.CreateAuthFactor(factorData)

		if err != nil {
			return factors, fmt.Errorf(
				"[SecretCreator] fail on creating auth factor with authorizer key %s: %s",
				authorizerKey,
				err.Error(),
			)
		}

		factors[authorizerKey] = factor
	}

	return factors, nil
}
