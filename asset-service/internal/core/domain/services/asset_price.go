package services

import (
	"context"
	"fmt"

	"goldvault/asset-service/internal/core/application/ports"
	"goldvault/asset-service/internal/core/domain/entity"
)

type AssetPriceDomainService struct {
	assetPricePersistence ports.AssetPricePersistence
}

func NewAssetPriceDomainService(assetPricePersistence ports.AssetPricePersistence) ports.AssetPriceDomainService {
	return &AssetPriceDomainService{
		assetPricePersistence: assetPricePersistence,
	}
}

func (a *AssetPriceDomainService) GetAllAssetPrices(ctx context.Context) (map[string]*entity.PriceDetails, error) {
	prices, err := a.assetPricePersistence.GetAllAssetPrices(ctx)
	if err != nil {
		return nil, err
	}

	return prices, nil
}

func (a *AssetPriceDomainService) GetLatestPrice(ctx context.Context, assetType string) (*entity.PriceDetails, error) {
	// validate asset type
	if !entity.IsValidAssetType(assetType) {
		return nil, fmt.Errorf("invalid asset type")
	}

	price, err := a.assetPricePersistence.GetPrice(ctx, assetType)
	if err != nil {
		return nil, err
	}

	return price, nil
}

func (a *AssetPriceDomainService) UpsertPrice(ctx context.Context, assetType string, prices *entity.PriceDetails) error {
	// validate asset type
	if !entity.IsValidAssetType(assetType) {
		return fmt.Errorf("invalid asset type")
	}

	err := a.assetPricePersistence.UpsertPrice(ctx, assetType, prices)
	if err != nil {
		return err
	}

	return nil
}

func (a *AssetPriceDomainService) DeleteAssetPrice(ctx context.Context, assetType string) error {
	// validate asset type
	if !entity.IsValidAssetType(assetType) {
		return fmt.Errorf("invalid asset type")
	}

	err := a.assetPricePersistence.DeleteAssetPrice(ctx, assetType)
	if err != nil {
		return err
	}

	return nil
}
