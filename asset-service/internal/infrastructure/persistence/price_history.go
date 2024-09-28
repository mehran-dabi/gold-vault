package persistence

import (
	"context"
	"database/sql"

	"goldvault/asset-service/internal/core/application/ports"
	"goldvault/asset-service/internal/core/domain/entity"
	"goldvault/asset-service/internal/infrastructure/persistence/models"
	"goldvault/asset-service/internal/infrastructure/persistence/queries"
)

type PriceHistoryPersistence struct {
	db *sql.DB
}

func NewPriceHistoryPersistence(db *sql.DB) ports.PriceHistoryPersistence {
	return &PriceHistoryPersistence{db: db}
}

func (p *PriceHistoryPersistence) GetHistoryByAssetType(ctx context.Context, assetType string, limit, offset int64) ([]*entity.PriceHistory, error) {
	rows, err := p.db.QueryContext(ctx, queries.GetPriceHistory, assetType, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var priceHistories []*entity.PriceHistory
	for rows.Next() {
		var ph models.PriceHistory
		err := ph.Scan(rows)
		if err != nil {
			return nil, err
		}
		priceHistories = append(priceHistories, ph.ToEntity())
	}

	return priceHistories, nil
}

func (p *PriceHistoryPersistence) InsertHistory(ctx context.Context, priceHistory *entity.PriceHistory) error {
	dbModel := models.ToPriceHistoryDB(priceHistory)

	_, err := p.db.ExecContext(ctx, queries.InsertPriceHistory, dbModel.AssetType, dbModel.BuyPrice, dbModel.SellPrice)
	if err != nil {
		return err
	}

	return nil
}
