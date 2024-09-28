package persistence

import (
	"context"
	"database/sql"

	"goldvault/user-service/internal/core/application/ports"
	"goldvault/user-service/internal/core/domain/entity"
	"goldvault/user-service/internal/infrastructure/persistence/models"
	"goldvault/user-service/internal/infrastructure/persistence/queries"
	"goldvault/user-service/pkg/serr"
)

type BankCardPersistence struct {
	db *sql.DB
}

func NewBankCardPersistence(db *sql.DB) ports.BankCardPersistence {
	return &BankCardPersistence{
		db: db,
	}
}

func (b *BankCardPersistence) CreateBankCard(ctx context.Context, bankCard *entity.BankCard) error {
	dbModel := models.ToBankCardDB(bankCard)
	err := b.db.QueryRowContext(ctx, queries.CreateBankCard, dbModel.UserID, dbModel.CardNumber).Scan(&dbModel.ID)
	if err != nil {
		return serr.DBError("CreateBankCard", "bank_card", err)
	}

	return nil
}

func (b *BankCardPersistence) GetBankCardsByUserID(ctx context.Context, userID int64) ([]*entity.BankCard, error) {
	rows, err := b.db.QueryContext(ctx, queries.GetBankCardsByUserID, userID)
	if err != nil {
		return nil, serr.DBError("GetBankCardsByUserID", "bank_card", err)
	}
	defer rows.Close()

	var bankCards []*entity.BankCard
	for rows.Next() {
		dbModel := new(models.BankCard)
		err := dbModel.Scan(rows)
		if err != nil {
			return nil, serr.DBError("GetBankCardsByUserID", "bank_card", err)
		}

		bankCards = append(bankCards, dbModel.ToEntity())
	}

	return bankCards, nil
}
