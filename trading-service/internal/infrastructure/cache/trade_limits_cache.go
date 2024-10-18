package cache

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"goldvault/trading-service/internal/config"
	"goldvault/trading-service/internal/core/application/ports"
	"goldvault/trading-service/pkg/serr"
)

type TradeLimitsCache struct {
	cacheRedisClient *config.CacheRedisConfig
}

func NewTradeLimitsCache(cacheRedisClient *config.CacheRedisConfig) ports.TradeLimitsCache {
	return &TradeLimitsCache{cacheRedisClient: cacheRedisClient}
}

func (t *TradeLimitsCache) SetGlobalLimits(ctx context.Context, assetType string, minBuy, minSell, maxBuy, maxSell, dailyBuyLimit, dailySellLimit float64) error {
	key := fmt.Sprintf("%s:%s", globalLimits, assetType)
	err := t.cacheRedisClient.RedisClient.HSet(ctx, key,
		"minBuy", minBuy,
		"maxBuy", maxBuy,
		"minSell", minSell,
		"maxSell", maxSell,
		"dailyBuyLimit", dailyBuyLimit,
		"dailySellLimit", dailySellLimit,
	).Err()
	if err != nil {
		return serr.DBError("HSet", key, err)
	}

	return nil
}

func (t *TradeLimitsCache) GetGlobalLimits(ctx context.Context, assetType string) (map[string]float64, error) {
	key := fmt.Sprintf("%s:%s", globalLimits, assetType)
	limits, err := t.cacheRedisClient.RedisClient.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, serr.DBError("HGetAll", key, err)
	}

	result := make(map[string]float64)
	for key, value := range limits {
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, serr.ServiceErr("GetGlobalLimits", "failed to convert limit value to float64", err, http.StatusInternalServerError)
		}
		result[key] = floatValue
	}

	return result, nil
}

func (t *TradeLimitsCache) GetUserDailyLimits(ctx context.Context, userID int64, assetType string) (map[string]float64, error) {
	key := fmt.Sprintf("%s:%d:%s", userDailyLimits, userID, assetType)
	dailyLimit, err := t.cacheRedisClient.RedisClient.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, serr.DBError("Get", key, err)
	}

	result := make(map[string]float64)
	for key, value := range dailyLimit {
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, serr.ServiceErr("GetUserDailyLimits", "failed to convert limit value to float64", err, http.StatusInternalServerError)
		}
		result[key] = floatValue
	}
	return result, nil
}

func (t *TradeLimitsCache) SetUserDailyLimits(ctx context.Context, userID int64, assetType string, dailySellLimit, dailyBuyLimit float64) error {
	key := fmt.Sprintf("%s:%d:%s", userDailyLimits, userID, assetType)
	err := t.cacheRedisClient.RedisClient.HSet(ctx, key,
		"dailyBuyLimit", dailyBuyLimit,
		"dailySellLimit", dailySellLimit,
	).Err()
	if err != nil {
		return serr.DBError("HSet", userDailyLimits, err)
	}

	// set the expiration time for the user daily limit
	err = t.cacheRedisClient.RedisClient.Expire(ctx, key, 24*time.Hour).Err()
	if err != nil {
		return serr.DBError("Expire", userDailyLimits, err)
	}

	return nil
}
