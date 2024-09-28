package models

import (
	"time"

	"goldvault/wallet-service/internal/core/domain/entity"
)

type Wallet struct {
	ID        int64
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	WalletsColumns     = "id, user_id, created_at, updated_at"
	WalletsColumnsNoID = "user_id, created_at, updated_at"
)

func ToWalletDB(entity *entity.Wallet) (*Wallet, error) {
	return &Wallet{
		ID:        entity.ID,
		UserID:    entity.UserID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}, nil
}

func (w *Wallet) Scan(scanner Scanner) error {
	return scanner.Scan(
		&w.ID,
		&w.UserID,
		&w.CreatedAt,
		&w.UpdatedAt,
	)
}

func (w *Wallet) ToEntity() *entity.Wallet {
	return &entity.Wallet{
		ID:        w.ID,
		UserID:    w.UserID,
		CreatedAt: w.CreatedAt,
		UpdatedAt: w.UpdatedAt,
	}
}
