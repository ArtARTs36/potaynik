package repository

import (
	"time"

	"github.com/artarts36/potaynik/internal/app/entity"
)

type MemorySecretRepository struct {
	secrets      map[string]*entity.Secret
	secretsDates map[string]time.Time
}

func NewMemorySecretRepository() *MemorySecretRepository {
	return &MemorySecretRepository{
		secrets:      map[string]*entity.Secret{},
		secretsDates: map[string]time.Time{},
	}
}

func (repo *MemorySecretRepository) Add(secret *entity.Secret) error {
	_, exists := repo.secrets[secret.Key]

	if exists {
		return newSecretAlreadyExistsError(secret.Key)
	}

	repo.secrets[secret.Key] = secret
	repo.secretsDates[secret.Key] = time.Now()

	return nil
}

func (repo *MemorySecretRepository) Find(secretKey string) (*entity.Secret, error) {
	secret, exists := repo.secrets[secretKey]

	if !exists {
		return nil, nil
	}

	currTimeDur := time.Now().Sub(repo.secretsDates[secretKey])

	if currTimeDur.Seconds() > float64(secret.TTL) {
		return nil, nil
	}

	return secret, nil
}

func (repo *MemorySecretRepository) Delete(secretKey string) {
	delete(repo.secrets, secretKey)
	delete(repo.secretsDates, secretKey)
}
