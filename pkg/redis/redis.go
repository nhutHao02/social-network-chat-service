package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/nhutHao02/social-network-chat-service/config"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisClient struct {
	Rdb *redis.Client
}

func InitRedis(cfg *config.RedisConfig) *RedisClient {
	return &RedisClient{
		Rdb: redis.NewClient(&redis.Options{
			Addr:     cfg.Address,
			Password: cfg.Password,
			DB:       cfg.Db,
			PoolSize: cfg.PoolSize,
		}),
	}
}

func (r *RedisClient) SetCacheStringData(ctx context.Context, key string, data string, ttl time.Duration) (err error) {
	err = r.Rdb.Set(ctx, key, data, ttl).Err()
	if err != nil {
		logger.Error("SetCacheStringData: set Cache fail = ", zap.Error(err))
		return
	}
	return
}

func (r *RedisClient) SetCacheStructData(
	ctx context.Context,
	key string,
	data interface{},
	ttl time.Duration,
) (err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Error("SetCacheStructData: Convert data to json fail = ", zap.Error(err))
		return
	}

	err = r.Rdb.Set(ctx, key, jsonData, ttl).Err()
	if err != nil {
		logger.Error("SetCacheStructData: Save data to cache fail = ", zap.Error(err))
		return
	}
	return nil
}

func (r *RedisClient) GetCache(ctx context.Context, key string) (val string, err error) {
	val, err = r.Rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return val, nil
	} else if err != nil {
		return
	}
	return
}

func (r *RedisClient) DeleteCache(ctx context.Context, key string) (err error) {
	err = r.Rdb.Del(ctx, key).Err()
	if err != nil {
		logger.Error("DeleteCache: delete cache fail = ", zap.Error(err))
		return
	}
	return
}

func (r *RedisClient) ConvertDataToStruct(dest interface{}, data string) (err error) {
	err = json.Unmarshal([]byte(data), &dest)
	if err != nil {
		logger.Error("ConvertDataToStruct: convert data from cache to struct fail = ", zap.Error(err))
		return
	}
	return
}
