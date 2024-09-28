package queries

import "goldvault/asset-service/internal/infrastructure/persistence/models"

const (
	InsertPriceHistory = `
		INSERT INTO ` + models.PricesHistoryTableName + ` (` + models.PriceHistoryColumnsNoID + `)
		VALUES ($1, $2, $3, NOW())
		RETURNING id;
	`

	GetPriceHistory = `
		SELECT ` + models.PriceHistoryColumns + `
		FROM ` + models.PricesHistoryTableName + `
		WHERE asset_type = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3;
	`
)
