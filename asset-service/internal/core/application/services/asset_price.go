package services

import (
	"context"
	"net/http"

	"goldvault/asset-service/internal/core/application/ports"
	"goldvault/asset-service/internal/core/domain/entity"
	"goldvault/asset-service/pkg/serr"
)

type AssetPriceService struct {
	assetPriceDomainService   ports.AssetPriceDomainService
	priceCache                ports.PriceCachePorts
	priceHistoryDomainService ports.PriceHistoryDomainService
}

func NewAssetPriceService(
	assetPriceDomainService ports.AssetPriceDomainService,
	priceCache ports.PriceCachePorts,
	priceHistoryDomainService ports.PriceHistoryDomainService,
) *AssetPriceService {
	return &AssetPriceService{
		assetPriceDomainService:   assetPriceDomainService,
		priceCache:                priceCache,
		priceHistoryDomainService: priceHistoryDomainService,
	}
}

func (a *AssetPriceService) GetLatestPrice(ctx context.Context, assetType string) (*entity.PriceDetails, error) {
	// check cache
	price, err := a.priceCache.GetPrice(ctx, assetType)
	if err == nil {
		return price, nil
	}

	// get from domain service
	price, err = a.assetPriceDomainService.GetLatestPrice(ctx, assetType)
	if err != nil {
		return nil, serr.ServiceErr("AssetPriceService.GetLatestPrice", err.Error(), err, http.StatusInternalServerError)
	}

	// save to cache
	err = a.priceCache.SavePrice(ctx, assetType, price)
	if err != nil {
		return nil, serr.ServiceErr("AssetPriceService.GetLatestPrice", err.Error(), err, http.StatusInternalServerError)
	}

	return price, nil
}

func (a *AssetPriceService) GetLatestPrices(ctx context.Context, assetTypes []string) (map[string]*entity.PriceDetails, error) {
	prices := make(map[string]*entity.PriceDetails)

	for _, assetType := range assetTypes {
		price, err := a.GetLatestPrice(ctx, assetType)
		if err != nil {
			return nil, serr.ServiceErr("AssetPriceService.GetLatestPrices", err.Error(), err, http.StatusInternalServerError)
		}

		prices[assetType] = price
	}

	return prices, nil
}

func (a *AssetPriceService) UpsertPrice(ctx context.Context, assetType string, prices *entity.PriceDetails) error {
	// upsert price in database
	err := a.assetPriceDomainService.UpsertPrice(ctx, assetType, prices)
	if err != nil {
		return serr.ServiceErr("AssetPriceService.UpsertPrice", err.Error(), err, http.StatusInternalServerError)
	}

	// update price in cache
	err = a.priceCache.SavePrice(ctx, assetType, prices)
	if err != nil {
		return serr.ServiceErr("AssetPriceService.UpsertPrice", err.Error(), err, http.StatusInternalServerError)
	}

	// add to price history
	err = a.priceHistoryDomainService.AddPriceHistory(ctx, assetType, prices)
	if err != nil {
		return serr.ServiceErr("AssetPriceService.UpsertPrice", err.Error(), err, http.StatusInternalServerError)
	}

	return nil
}

func (a *AssetPriceService) DeleteAssetPrice(ctx context.Context, assetType string) error {
	// remove from database
	err := a.assetPriceDomainService.DeleteAssetPrice(ctx, assetType)
	if err != nil {
		return serr.ServiceErr("AssetPriceService.DeleteAssetPrice", err.Error(), err, http.StatusInternalServerError)
	}

	// remove price from cache
	err = a.priceCache.RemovePrice(ctx, assetType)
	if err != nil {
		return serr.ServiceErr("AssetPriceService.DeleteAssetPrice", err.Error(), err, http.StatusInternalServerError)
	}

	return nil
}

func (a *AssetPriceService) GetAllAssetPrices(ctx context.Context) (map[string]*entity.PriceDetails, error) {
	// get from cache
	prices, err := a.priceCache.GetAllAssetPrices(ctx)
	if err == nil {
		return prices, nil
	}

	// get from domain service
	prices, err = a.assetPriceDomainService.GetAllAssetPrices(ctx)
	if err != nil {
		return nil, serr.ServiceErr("AssetPriceService.GetAllAssetPrices", err.Error(), err, http.StatusInternalServerError)
	}

	return prices, nil
}

func (a *AssetPriceService) GetPriceChangeStep(ctx context.Context) (float64, error) {
	return a.priceCache.GetPriceChangeStep(ctx)
}

func (a *AssetPriceService) SetPriceChangeStep(ctx context.Context, step float64) error {
	return a.priceCache.SetPriceChangeStep(ctx, step)
}

func (a *AssetPriceService) UpdateAssetPriceByStep(ctx context.Context, assetType string) error {
	step, err := a.GetPriceChangeStep(ctx)
	if err != nil {
		return serr.ServiceErr("AssetPriceService.UpdateAssetPriceByStep", err.Error(), err, http.StatusInternalServerError)
	}

	err = a.assetPriceDomainService.UpdateAssetPriceByStep(ctx, step, assetType)
	if err != nil {
		return serr.ServiceErr("AssetPriceService.UpdateAssetPriceByStep", err.Error(), err, http.StatusInternalServerError)
	}

	return nil
}
