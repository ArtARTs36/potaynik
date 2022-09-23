package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/artarts36/potaynik/internal/app/entity"
)

type RedisSecretRepository struct {
	redis  *redis.Client
	prefix string
}

func NewRedisSecretRepository(redis *redis.Client, prefix string) *RedisSecretRepository {
	return &RedisSecretRepository{redis: redis, prefix: prefix}
}

func (repo *RedisSecretRepository) Add(secret *entity.Secret) error {
	secretJSON, err := json.Marshal(secret)

	if err != nil {
		return err
	}

	return repo.redis.Set(
		context.Background(),
		repo.createKey(secret.Key),
		secretJSON,
		secret.Duration(),
	).Err()
}

func (repo *RedisSecretRepository) Find(secretKey string) (*entity.Secret, error) {
	res, err := repo.redis.Get(context.Background(), repo.createKey(secretKey)).Result()

	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}

		return nil, err
	}

	if res == "" {
		return nil, nil
	}

	var secret entity.Secret

	err = json.Unmarshal([]byte(res), &secret)

	if err != nil {
		return nil, err
	}

	return &secret, nil
}

func (repo *RedisSecretRepository) Delete(secretKey string) {
	repo.redis.Del(context.Background(), repo.createKey(secretKey))
}

func (repo *RedisSecretRepository) createKey(secretKey string) string {
	return fmt.Sprintf("%s_secret_%s", repo.prefix, secretKey)
}
