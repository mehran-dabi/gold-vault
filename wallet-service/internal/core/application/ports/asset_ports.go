package ports

import (
	"context"
	"database/sql"

	"goldvault/wallet-service/internal/core/domain/entity"
)

type (
	AssetPersistence interface {
		AddAsset(ctx context.Context, asset *entity.Asset) error
		UpsertAssetBalance(ctx context.Context, tx *sql.Tx, walletID int64, assetType string, amount float64) error
		GetAssetsByWalletID(ctx context.Context, walletID int64) ([]entity.Asset, error)
		GetAssetForUpdate(ctx context.Context, tx *sql.Tx, walletID int64, assetType string) (*entity.Asset, error)
		GetAssetBalance(ctx context.Context, walletID int64, assetType string) (float64, error)
	}

	AssetDomainService interface {
		AddAsset(ctx context.Context, walletID int64, assetType string, amount float64) error
		GetAssetsByWalletID(ctx context.Context, walletID int64) ([]entity.Asset, error)
		UpdateAssetBalance(ctx context.Context, tx *sql.Tx, userID int64, assetType string, amount float64) error
		GetAssetBalance(ctx context.Context, userID int64, assetType string) (float64, error)
	}
)
