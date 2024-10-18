package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"goldvault/asset-service/internal/config"
	"goldvault/asset-service/internal/core/application/ports"
	"goldvault/asset-service/internal/core/domain/entity"
)

// PriceCache provides persistence operations for OTPs
type PriceCache struct {
	cacheRedisClient *config.CacheRedisConfig
}

// NewPriceCache creates a new instance of PriceCache
func NewPriceCache(cacheRedisClient *config.CacheRedisConfig) ports.PriceCachePorts {
	return &PriceCache{cacheRedisClient: cacheRedisClient}
}

func (p *PriceCache) GetPrice(ctx context.Context, assetType string) (*entity.PriceDetails, error) {
	key := fmt.Sprintf("%s:%s", assetPriceKey, assetType)
	pricesString, err := p.cacheRedisClient.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	// unmarshal the string
	var prices entity.PriceDetails
	err = json.Unmarshal([]byte(pricesString), &prices)
	if err != nil {
		return nil, err
	}
	return &prices, nil
}

func (p *PriceCache) SavePrice(ctx context.Context, assetType string, prices *entity.PriceDetails) error {
	// marshal the prices entity
	pricesBytes, err := json.Marshal(prices)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s:%s", assetPriceKey, assetType)
	return p.cacheRedisClient.RedisClient.Set(ctx, key, string(pricesBytes), 0).Err()
}

func (p *PriceCache) RemovePrice(ctx context.Context, assetType string) error {
	key := fmt.Sprintf("%s:%s", assetPriceKey, assetType)
	return p.cacheRedisClient.RedisClient.Del(ctx, key).Err()
}

func (p *PriceCache) GetAllAssetPrices(ctx context.Context) (map[string]*entity.PriceDetails, error) {
	// Build keys for MGET based on asset types
	keys := make([]string, len(entity.AssetTypes))
	for i, asset := range entity.AssetTypes {
		keys[i] = fmt.Sprintf("%s:%s", assetPriceKey, asset)
	}

	// Fetch all prices using MGET
	prices, err := p.cacheRedisClient.RedisClient.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	// Build the map of asset types and prices
	assetPrices := make(map[string]*entity.PriceDetails)
	for i, price := range prices {
		if price == nil {
			continue
		}

		// Unmarshal the price
		var priceDetails entity.PriceDetails
		err = json.Unmarshal([]byte(price.(string)), &priceDetails)
		if err != nil {
			return nil, err
		}

		assetPrices[entity.AssetTypes[i].String()] = &priceDetails
	}

	return assetPrices, nil
}

func (p *PriceCache) SetPriceChangeStep(ctx context.Context, step float64) error {
	key := assetPriceChangeStep
	return p.cacheRedisClient.RedisClient.Set(ctx, key, step, 0).Err()
}

func (p *PriceCache) GetPriceChangeStep(ctx context.Context) (float64, error) {
	key := assetPriceChangeStep
	return p.cacheRedisClient.RedisClient.Get(ctx, key).Float64()
}
