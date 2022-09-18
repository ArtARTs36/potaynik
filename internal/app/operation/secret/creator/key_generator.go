package creator

import "github.com/google/uuid"

type KeyGenerator struct {
}

func (g *KeyGenerator) Generate() string {
	return "1"
	return uuid.New().String()
}
