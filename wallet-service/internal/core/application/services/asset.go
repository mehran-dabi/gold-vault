package services

import (
	"context"
	"database/sql"
	"net/http"

	"goldvault/wallet-service/internal/core/application/ports"
	"goldvault/wallet-service/internal/infrastructure/db"
	"goldvault/wallet-service/pkg/serr"
)

type AssetService struct {
	assetDomainService ports.AssetDomainService
}

func NewAssetService(assetDomainService ports.AssetDomainService) *AssetService {
	return &AssetService{assetDomainService: assetDomainService}
}

func (a *AssetService) AddAsset(ctx context.Context, walletID int64, assetType string, amount float64) error {
	err := a.assetDomainService.AddAsset(ctx, walletID, assetType, amount)
	if err != nil {
		return serr.ServiceErr("AssetApplicationService.AddAsset", err.Error(), err, http.StatusInternalServerError)
	}

	return nil
}

func (a *AssetService) UpdateAssetBalance(ctx context.Context, userID int64, assetType string, amount float64) error {
	// Start a transaction with read commited isolation level to update the asset balance without causing dirty reads.
	err := db.Transaction(ctx, sql.LevelReadCommitted, func(tx *sql.Tx) error {
		err := a.assetDomainService.UpdateAssetBalance(ctx, tx, userID, assetType, amount)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return serr.ServiceErr("AssetApplicationService.UpsertAssetBalance", err.Error(), err, http.StatusInternalServerError)
	}

	return nil
}

func (a *AssetService) GetAssetBalance(ctx context.Context, userID int64, assetType string) (float64, error) {
	balance, err := a.assetDomainService.GetAssetBalance(ctx, userID, assetType)
	if err != nil {
		return 0, serr.ServiceErr("AssetApplicationService.GetAssetBalance", err.Error(), err, http.StatusInternalServerError)
	}

	return balance, nil
}
