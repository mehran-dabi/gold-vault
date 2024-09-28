package queries

import "goldvault/trading-service/internal/infrastructure/persistence/models"

const (
	CreateInventory = `
		INSERT INTO ` + models.InventoryTableName + ` (` + models.InventoryColumnsNoID + `)
		VALUES ($1, $2, NOW(), NOW())
		RETURNING id;
	`

	GetInventoryForUpdate = `
		SELECT ` + models.InventoryColumns + `
		FROM ` + models.InventoryTableName + `
		WHERE asset_type = $1
		FOR UPDATE;
	`

	UpdateInventoryQuantity = `
		UPDATE ` + models.InventoryTableName + `
		SET total_quantity = $2, updated_at = NOW()
		WHERE asset_type = $1;
	`

	DeleteInventory = `
		DELETE FROM ` + models.InventoryTableName + `
		WHERE asset_type = $1;
	`

	GetInventory = `
		SELECT ` + models.InventoryColumns + `
		FROM ` + models.InventoryTableName + `;
	`
)
