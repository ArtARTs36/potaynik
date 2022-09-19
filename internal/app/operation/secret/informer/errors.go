package informer

import "fmt"

type SecretNotFoundError struct {
	secretKey string
}

func newSecretNotFoundError(secretKey string) *SecretNotFoundError {
	return &SecretNotFoundError{secretKey: secretKey}
}

func (e *SecretNotFoundError) Error() string {
	return fmt.Sprintf("Secret with key %s not found", e.secretKey)
}
