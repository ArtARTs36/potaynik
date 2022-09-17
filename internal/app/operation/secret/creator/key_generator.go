package creator

import "github.com/google/uuid"

type KeyGenerator struct {
}

func (g *KeyGenerator) Generate() string {
	return uuid.New().String()
}
