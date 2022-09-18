package repository

import "github.com/artarts36/potaynik/internal/app/entity"

type MemorySecretRepository struct {
	secrets map[string]*entity.Secret
}

func NewMemorySecretRepository() *MemorySecretRepository {
	return &MemorySecretRepository{
		secrets: map[string]*entity.Secret{},
	}
}

func (repo *MemorySecretRepository) Add(secret *entity.Secret) error {
	_, exists := repo.secrets[secret.Key]

	if exists {
		return newSecretAlreadyExistsError(secret.Key)
	}

	repo.secrets[secret.Key] = secret

	return nil
}

func (repo *MemorySecretRepository) Find(secretKey string) (*entity.Secret, error) {
	secret, exists := repo.secrets[secretKey]

	if exists {
		return secret, nil
	}

	return nil, nil
}

func (repo *MemorySecretRepository) Delete(secretKey string) {
	delete(repo.secrets, secretKey)
}
