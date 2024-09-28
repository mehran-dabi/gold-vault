package models

import (
	"time"

	"goldvault/user-service/internal/core/domain/entity"
)

type BankCard struct {
	ID         int64
	UserID     int64
	CardNumber string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

const (
	BankCardsTableName   = "user_bank_cards"
	BankCardsColumns     = "id, user_id, card_number, created_at, updated_at"
	BankCardsColumnsNoID = "user_id, card_number, created_at, updated_at"
)

func (b *BankCard) ToEntity() *entity.BankCard {
	return &entity.BankCard{
		ID:         b.ID,
		UserID:     b.UserID,
		CardNumber: b.CardNumber,
		CreatedAt:  b.CreatedAt,
		UpdatedAt:  b.UpdatedAt,
	}
}

func ToBankCardDB(e *entity.BankCard) *BankCard {
	return &BankCard{
		ID:         e.ID,
		UserID:     e.UserID,
		CardNumber: e.CardNumber,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
	}
}

func (b *BankCard) Scan(scanner Scanner) error {
	return scanner.Scan(&b.ID, &b.UserID, &b.CardNumber, &b.CreatedAt, &b.UpdatedAt)
}
