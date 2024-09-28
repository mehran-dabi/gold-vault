package queries

import "goldvault/wallet-service/internal/infrastructure/persistence/models"

const (
	CreateWallet = `
		INSERT INTO wallets (` + models.WalletsColumnsNoID + `)
		VALUES ($1, $2, $3) RETURNING id;
	`

	GetWalletByUserID = `
		SELECT ` + models.WalletsColumns + `
		FROM wallets
		WHERE user_id = $1;
	`

	GetWalletByUserIDWithLock = `
		SELECT ` + models.WalletsColumns + `
		FROM wallets
		WHERE user_id = $1
		FOR UPDATE;
	`

	GetWalletsWithPagination = `
		SELECT ` + models.WalletsColumns + `
		FROM wallets
		LIMIT $1 OFFSET $2;
	`
)
