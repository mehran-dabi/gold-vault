package queries

import "goldvault/wallet-service/internal/infrastructure/persistence/models"

const (
	AddTransaction = `
		INSERT INTO transactions ` + models.TransactionColumnsNoID + `
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	GetTransactionsByAssetID = `
		SELECT ` + models.TransactionColumns + `
		FROM transactions
		WHERE asset_id = $1;
	`

	GetTransactionsByWalletID = `
		SELECT ` + models.TransactionColumns + `
		FROM transactions
		WHERE wallet_id = $1;
	`

	GetTransactionsWithPagination = `
		SELECT ` + models.TransactionColumns + `
		FROM transactions
		LIMIT $1 OFFSET $2;
	`
)
