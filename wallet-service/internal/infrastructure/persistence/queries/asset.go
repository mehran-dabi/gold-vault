package queries

import "goldvault/wallet-service/internal/infrastructure/persistence/models"

const (
	AddAsset = `
		INSERT INTO ` + models.AssetsTableName + ` (` + models.AssetsColumnsNoID + `)
		VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id;
	`

	GetAssetsByWalletID = `
		SELECT ` + models.AssetsColumns + `
		FROM ` + models.AssetsTableName + `
		WHERE wallet_id = $1;
	`

	GetAssetForUpdate = `
		SELECT ` + models.AssetsColumns + `
		FROM ` + models.AssetsTableName + `
		WHERE wallet_id = $1 AND type = $2
		FOR UPDATE;
	`

	UpsertAssetBalance = `
		INSERT INTO ` + models.AssetsTableName + ` (wallet_id, type, balance, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		ON CONFLICT (wallet_id, type) 
		DO UPDATE 
		SET balance = ` + models.AssetsTableName + `.balance + $3, updated_at = NOW();
	`

	GetAssetBalance = `
		SELECT balance
		FROM ` + models.AssetsTableName + `
		WHERE wallet_id = $1 AND type = $2;
	`
)
