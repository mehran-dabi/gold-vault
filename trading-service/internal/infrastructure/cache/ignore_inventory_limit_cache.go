package cache

import (
	"context"

	"goldvault/trading-service/internal/config"
	"goldvault/trading-service/internal/core/application/ports"
	"goldvault/trading-service/pkg/serr"
)

type IgnoreInventoryLimitCache struct {
	cacheRedisClient *config.CacheRedisConfig
}

func NewIgnoreInventoryLimitCache(cacheRedisClient *config.CacheRedisConfig) ports.IgnoreInventoryLimitCache {
	return &IgnoreInventoryLimitCache{cacheRedisClient: cacheRedisClient}
}

func (i *IgnoreInventoryLimitCache) Set(ctx context.Context) error {
	err := i.cacheRedisClient.RedisClient.Set(ctx, ignoreInventoryLimitKey, true, 0).Err()
	if err != nil {
		return serr.DBError("Set", "ignore_inventory_limit", err)
	}

	return nil
}

func (i *IgnoreInventoryLimitCache) Unset(ctx context.Context) error {
	err := i.cacheRedisClient.RedisClient.Set(ctx, ignoreInventoryLimitKey, false, 0).Err()
	if err != nil {
		return serr.DBError("Unset", "ignore_inventory_limit", err)
	}

	return nil
}

func (i *IgnoreInventoryLimitCache) Get(ctx context.Context) (bool, error) {
	ignoreInventoryLimit, err := i.cacheRedisClient.RedisClient.Get(ctx, ignoreInventoryLimitKey).Bool()
	if err != nil {
		return false, serr.DBError("Get", "ignore_inventory_limit", err)
	}

	return ignoreInventoryLimit, nil
}
