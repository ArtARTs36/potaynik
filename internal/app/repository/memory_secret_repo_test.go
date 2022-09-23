package repository

import (
	"testing"

	"github.com/artarts36/potaynik/internal/app/entity"
	"github.com/stretchr/testify/assert"
)

func TestMemorySecretRepository_AddOK(t *testing.T) {
	repo := NewMemorySecretRepository()

	err := repo.Add(&entity.Secret{
		Key: "12",
		TTL: 5,
	})

	assert.Nil(t, err)
}

func TestMemorySecretRepository_AddAlreadyExists(t *testing.T) {
	repo := NewMemorySecretRepository()

	err := repo.Add(&entity.Secret{
		Key: "12",
		TTL: 5,
	})

	assert.Nil(t, err)

	err = repo.Add(&entity.Secret{
		Key: "12",
		TTL: 5,
	})

	assert.NotNil(t, err)
	assert.EqualError(t, err, "Secret with key 12 already exists")
}
