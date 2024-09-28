package queries

import "goldvault/trading-service/internal/infrastructure/persistence/models"

const (
	CreateOrder = `
		INSERT INTO ` + models.OrderTableName + ` (` + models.OrderColumnsNoID + `)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id;
	`

	GetOrderByUserIDAndStatus = `
		SELECT ` + models.OrderColumns + `
		FROM ` + models.OrderTableName + `
		WHERE user_id = $1 AND status = $2;
	`

	GetOrderByStatus = `
		SELECT ` + models.OrderColumns + `
		FROM ` + models.OrderTableName + `
		WHERE status = $1;
	`
)
