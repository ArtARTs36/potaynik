package viewer

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

type SecretViewForbiddenError struct {
	secretKey string
	Reason    string
}

func newSecretViewForbiddenError(secretKey string, reason string) *SecretViewForbiddenError {
	return &SecretViewForbiddenError{secretKey: secretKey, Reason: reason}
}

func (e *SecretViewForbiddenError) Error() string {
	return fmt.Sprintf("Secret with key %s forbidden for view", e.secretKey)
}
