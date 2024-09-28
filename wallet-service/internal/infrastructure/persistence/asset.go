package persistence

import (
	"context"
	"database/sql"
	"errors"

	"goldvault/wallet-service/internal/core/application/ports"
	"goldvault/wallet-service/internal/core/domain/entity"
	"goldvault/wallet-service/internal/infrastructure/persistence/models"
	"goldvault/wallet-service/internal/infrastructure/persistence/queries"
	"goldvault/wallet-service/pkg/serr"
)

type AssetPersistence struct {
	db *sql.DB
}

func NewAssetPersistence(db *sql.DB) ports.AssetPersistence {
	return &AssetPersistence{db: db}
}

func (a *AssetPersistence) AddAsset(ctx context.Context, asset *entity.Asset) error {
	dbModel, err := models.ToAssetDB(asset)
	if err != nil {
		return serr.DBError("AddAsset", "asset", err)
	}

	err = a.db.QueryRowContext(ctx, queries.AddAsset, dbModel.WalletID, dbModel.Type, dbModel.Balance).
		Scan(&dbModel.ID)
	if err != nil {
		return serr.DBError("AddAsset", "asset", err)
	}

	return nil
}

func (a *AssetPersistence) UpsertAssetBalance(ctx context.Context, tx *sql.Tx, walletID int64, assetType string, amount float64) error {
	_, err := tx.ExecContext(ctx, queries.UpsertAssetBalance, walletID, assetType, amount)
	if err != nil {
		return serr.DBError("UpsertAssetBalance", "asset", err)
	}

	return nil
}

func (a *AssetPersistence) GetAssetForUpdate(ctx context.Context, tx *sql.Tx, walletID int64, assetType string) (*entity.Asset, error) {
	row := tx.QueryRowContext(ctx, queries.GetAssetForUpdate, walletID, assetType)

	var dbModel models.Asset
	err := dbModel.Scan(row)
	if err != nil {
		return nil, serr.DBError("GetAssetForUpdate", "asset", err)
	}

	return dbModel.ToEntity(), nil
}

func (a *AssetPersistence) GetAssetsByWalletID(ctx context.Context, walletID int64) ([]entity.Asset, error) {
	rows, err := a.db.QueryContext(ctx, queries.GetAssetsByWalletID, walletID)
	if err != nil {
		return nil, serr.DBError("GetAssetsByWalletID", "asset", err)
	}
	defer rows.Close()

	var assets []entity.Asset
	for rows.Next() {
		var dbModel models.Asset
		err = dbModel.Scan(rows)
		if err != nil {
			return nil, serr.DBError("GetAssetsByWalletID", "asset", err)
		}

		assets = append(assets, *dbModel.ToEntity())
	}

	return assets, nil
}

func (a *AssetPersistence) GetAssetBalance(ctx context.Context, walletID int64, assetType string) (float64, error) {
	row := a.db.QueryRowContext(ctx, queries.GetAssetBalance, walletID, assetType)

	var balance float64
	err := row.Scan(&balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil // No asset found
		}
		return 0, serr.DBError("GetAssetBalance", "asset", err)
	}

	return balance, nil
}
