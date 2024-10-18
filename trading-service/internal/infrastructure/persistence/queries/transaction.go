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

	GetTransactionsSummarySingleDay = `
	SELECT 
		SUM(CASE WHEN transaction_type = 'buy' THEN quantity ELSE 0 END) AS total_buys,
		SUM(CASE WHEN transaction_type = 'sell' THEN quantity ELSE 0 END) AS total_sells
	FROM 
		` + models.TransactionTableName + `
	WHERE 
		DATE(created_at) = $1
	    AND status = 'completed'
		AND asset_type = $2;
	`

	GetTransactionsSummary = `
	SELECT 
		SUM(CASE WHEN transaction_type = 'buy' THEN quantity ELSE 0 END) AS total_buys,
		SUM(CASE WHEN transaction_type = 'sell' THEN quantity ELSE 0 END) AS total_sells
	FROM 
		` + models.TransactionTableName + `
	WHERE 
		status = 'completed'
		AND asset_type = $1;
	`
)
