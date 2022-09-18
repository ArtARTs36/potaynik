package creator

import "github.com/artarts36/potaynik/internal/app/entity"

type SecretRepository interface {
	Add(secret *entity.Secret) error
}
