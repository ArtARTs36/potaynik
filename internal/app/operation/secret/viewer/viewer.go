package viewer

import (
	"github.com/artarts36/potaynik/internal/app/entity"
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

func (v *Viewer) View(secretKey string) (*string, error) {
	secret, _ := v.secrets.Find(secretKey)

	if secret == nil {
		return nil, newSecretNotFoundError(secretKey)
	}

	v.secrets.Delete(secretKey)

	return &secret.Value, nil
}
