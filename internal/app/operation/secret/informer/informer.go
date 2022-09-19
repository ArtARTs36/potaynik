package informer

import "github.com/artarts36/potaynik/internal/app/entity"

type Informer struct {
	secrets Repository
}

type Repository interface {
	Find(key string) (*entity.Secret, error)
}

type SecretInfo struct {
	SecretKey   string   `json:"secret_key"`
	AuthFactors []string `json:"auth_factors"`
}

func New(secrets Repository) *Informer {
	return &Informer{secrets: secrets}
}

func (i *Informer) Info(secretKey string) (SecretInfo, error) {
	secret, err := i.secrets.Find(secretKey)

	if err != nil {
		return SecretInfo{}, err
	}

	if secret == nil {
		return SecretInfo{}, newSecretNotFoundError(secretKey)
	}

	return SecretInfo{
		SecretKey:   secret.Key,
		AuthFactors: secret.AuthFactorNames(),
	}, nil
}
