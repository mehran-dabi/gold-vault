package queries

import "goldvault/asset-service/internal/infrastructure/persistence/models"

const (
	GetPriceByAssetType = `
		SELECT ` + models.AssetPriceColumns + `
		FROM ` + models.AssetPriceTableName + `
		WHERE asset_type = $1;
	`

	GetAllAssetPrices = `
		SELECT ` + models.AssetPriceColumns + `
		FROM ` + models.AssetPriceTableName + `;
	`

	UpsertAssetPrice = `
		INSERT INTO ` + models.AssetPriceTableName + ` (` + models.AssetPriceColumnsNoID + `)
		VALUES ($1, $2, $3, NOW(), NOW())
		ON CONFLICT (asset_type) DO UPDATE
		SET buy_price = $2, sell_price = $3, updated_at = NOW()
		RETURNING id;
	`

	DeleteAssetPrice = `
		DELETE FROM ` + models.AssetPriceTableName + `
		WHERE asset_type = $1;
	`
)
