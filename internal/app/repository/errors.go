package repository

import "fmt"

type SecretAlreadyExistsError struct {
	SecretKey string
}

func newSecretAlreadyExistsError(secretKey string) *SecretAlreadyExistsError {
	return &SecretAlreadyExistsError{SecretKey: secretKey}
}

func (err *SecretAlreadyExistsError) Error() string {
	return fmt.Sprintf(
		"Secret with name %s already exists",
		err.SecretKey,
	)
}
