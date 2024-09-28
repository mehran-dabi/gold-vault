package services

import (
	"context"
	"net/http"

	"goldvault/asset-service/internal/core/application/ports"
	"goldvault/asset-service/internal/core/domain/entity"
	"goldvault/asset-service/pkg/serr"
)

type PriceHistoryService struct {
	priceHistoryDomainService ports.PriceHistoryDomainService
}

func NewPriceHistoryService(priceHistoryDomainService ports.PriceHistoryDomainService) *PriceHistoryService {
	return &PriceHistoryService{priceHistoryDomainService: priceHistoryDomainService}
}

func (p *PriceHistoryService) GetAssetPriceHistory(ctx context.Context, assetType string, limit, offset int64) ([]*entity.PriceHistory, error) {
	priceHistory, err := p.priceHistoryDomainService.GetAssetPriceHistory(ctx, assetType, limit, offset)
	if err != nil {
		return nil, serr.ServiceErr("PriceHistory.GetAssetPriceHistory", err.Error(), err, http.StatusInternalServerError)
	}

	return priceHistory, nil
}
