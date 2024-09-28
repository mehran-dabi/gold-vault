package queries

import "goldvault/user-service/internal/infrastructure/persistence/models"

const (
	CreateBankCard = `
		INSERT INTO ` + models.BankCardsTableName + ` (` + models.BankCardsColumnsNoID + `)
		VALUES ($1, $2, NOW(), NOW())
		RETURNING id;
	`

	GetBankCardsByUserID = `
		SELECT ` + models.BankCardsColumns + `
		FROM ` + models.BankCardsTableName + `
		WHERE user_id = $1;
	`
)
