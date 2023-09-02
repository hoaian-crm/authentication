package repositories

import (
	"context"
	"encoding/json"
	"main/config"
)

type CacheRepository[Model any] struct {
}

var ctx = context.Background()

func (cache CacheRepository[Model]) Read(key string, result *Model) (error) {

	p, err := config.RedisClient.Get(ctx, key).Result()
  if err := json.Unmarshal([]byte(p), result); err != nil {
    return err;
  }

  return err;
}

func (cache *CacheRepository[Model]) Set(key string, value *Model) error {

	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return config.RedisClient.Set(ctx, key, string(b), 0).Err()
}
