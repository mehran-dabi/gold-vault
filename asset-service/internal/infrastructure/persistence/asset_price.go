package persistence

import (
	"context"
	"database/sql"

	"goldvault/asset-service/internal/core/application/ports"
	"goldvault/asset-service/internal/core/domain/entity"
	"goldvault/asset-service/internal/infrastructure/persistence/models"
	"goldvault/asset-service/internal/infrastructure/persistence/queries"
	"goldvault/asset-service/pkg/serr"
)

type AssetPricePersistence struct {
	db *sql.DB
}

func NewAssetPricePersistence(db *sql.DB) ports.AssetPricePersistence {
	return &AssetPricePersistence{
		db: db,
	}
}

func (a *AssetPricePersistence) GetPrice(ctx context.Context, assetType string) (*entity.PriceDetails, error) {
	row := a.db.QueryRowContext(ctx, queries.GetPriceByAssetType, assetType)

	var dbModel models.AssetPrice
	err := dbModel.Scan(row)
	if err != nil {
		return nil, serr.DBError("GetPrice", "asset_price", err)
	}

	return &entity.PriceDetails{BuyPrice: dbModel.BuyPrice, SellPrice: dbModel.SellPrice}, nil
}

func (a *AssetPricePersistence) UpsertPrice(ctx context.Context, assetType string, prices *entity.PriceDetails) error {
	_, err := a.db.ExecContext(ctx, queries.UpsertAssetPrice, assetType, prices.BuyPrice, prices.SellPrice)
	if err != nil {
		return serr.DBError("UpsertPrice", "asset_price", err)
	}

	return nil
}

func (a *AssetPricePersistence) DeleteAssetPrice(ctx context.Context, assetType string) error {
	_, err := a.db.ExecContext(ctx, queries.DeleteAssetPrice, assetType)
	if err != nil {
		return serr.DBError("DeleteAssetPrice", "asset_price", err)
	}

	return nil
}

func (a *AssetPricePersistence) GetAllAssetPrices(ctx context.Context) (map[string]*entity.PriceDetails, error) {
	rows, err := a.db.QueryContext(ctx, queries.GetAllAssetPrices)
	if err != nil {
		return nil, serr.DBError("GetAllAssetPrices", "asset_price", err)
	}
	defer rows.Close()

	assetPrices := make(map[string]*entity.PriceDetails)
	for rows.Next() {
		var dbModel models.AssetPrice
		err := dbModel.Scan(rows)
		if err != nil {
			return nil, serr.DBError("GetAllAssetPrices", "asset_price", err)
		}

		assetPrices[dbModel.AssetType] = &entity.PriceDetails{BuyPrice: dbModel.BuyPrice, SellPrice: dbModel.SellPrice}
	}

	return assetPrices, nil
}

func (a *AssetPricePersistence) UpdateAssetPriceByStep(ctx context.Context, step float64, assetType string) error {
	_, err := a.db.ExecContext(ctx, queries.UpdateAssetPriceByStep, step, assetType)
	if err != nil {
		return serr.DBError("UpdateAssetPriceByStep", "asset_price", err)
	}

	return nil
}
