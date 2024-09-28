package services

import (
	"context"
	"fmt"

	"goldvault/asset-service/internal/core/application/ports"
	"goldvault/asset-service/internal/core/domain/entity"
)

type PriceHistoryDomainService struct {
	priceHistoryPersistence ports.PriceHistoryPersistence
}

func NewPriceHistoryDomainService(priceHistoryPersistence ports.PriceHistoryPersistence) ports.PriceHistoryDomainService {
	return &PriceHistoryDomainService{priceHistoryPersistence: priceHistoryPersistence}
}

func (p *PriceHistoryDomainService) GetAssetPriceHistory(ctx context.Context, assetType string, limit, offset int64) ([]*entity.PriceHistory, error) {
	// validate asset type
	if !entity.IsValidAssetType(assetType) {
		return nil, fmt.Errorf("invalid asset type")
	}

	priceHistory, err := p.priceHistoryPersistence.GetHistoryByAssetType(ctx, assetType, limit, offset)
	if err != nil {
		return nil, err
	}

	return priceHistory, nil
}

func (p *PriceHistoryDomainService) AddPriceHistory(ctx context.Context, assetType string, prices *entity.PriceDetails) error {
	// validate asset type
	if !entity.IsValidAssetType(assetType) {
		return fmt.Errorf("invalid asset type")
	}

	priceHistory := &entity.PriceHistory{
		AssetType: entity.AssetType(assetType),
		Prices: entity.PriceDetails{
			BuyPrice:  prices.BuyPrice,
			SellPrice: prices.SellPrice,
		},
	}

	// validate entity
	if err := priceHistory.Validate(); err != nil {
		return err
	}

	err := p.priceHistoryPersistence.InsertHistory(ctx, priceHistory)
	if err != nil {
		return err
	}

	return nil
}
