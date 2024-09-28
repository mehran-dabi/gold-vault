package queries

import "goldvault/trading-service/internal/infrastructure/persistence/models"

const (
	InsertTransaction = `
		INSERT INTO ` + models.TransactionTableName + ` (` + models.TransactionColumnsNoID + `)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING id;
	`

	UpdateTransactionStatus = `
		UPDATE ` + models.TransactionTableName + `
		SET status = $1
		WHERE id = $2;
	`

	GetUserTransactions = `
		SELECT ` + models.TransactionColumns + `
		FROM ` + models.TransactionTableName + `
		WHERE user_id = $1;
	`

	GetAllTransactions = `
		SELECT ` + models.TransactionColumns + `
		FROM ` + models.TransactionTableName + `
		LIMIT $1 OFFSET $2;
	`
)
