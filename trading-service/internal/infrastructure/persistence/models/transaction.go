package models

import (
	"time"

	"goldvault/trading-service/internal/core/domain/entity"
)

type Transaction struct {
	ID              int64
	UserID          int64
	AssetType       string
	Quantity        float64
	Price           float64
	TransactionType string // Buy or Sell
	Status          string // Pending, Completed, Cancelled
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

const (
	TransactionTableName   = "transactions"
	TransactionColumns     = "id, user_id, asset_type, quantity, price, transaction_type, status, created_at, updated_at"
	TransactionColumnsNoID = "user_id, asset_type, quantity, price, transaction_type, status, created_at, updated_at"
)

func (t *Transaction) Scan(scanner Scanner) error {
	return scanner.Scan(
		&t.ID,
		&t.UserID,
		&t.AssetType,
		&t.Quantity,
		&t.Price,
		&t.TransactionType,
		&t.Status,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
}

func ToTransactionDB(entity *entity.Transaction) *Transaction {
	return &Transaction{
		ID:              entity.ID,
		UserID:          entity.UserID,
		AssetType:       entity.AssetType,
		Quantity:        entity.Quantity,
		Price:           entity.Price,
		TransactionType: entity.TransactionType.String(),
		Status:          entity.Status.String(),
		CreatedAt:       entity.CreatedAt,
		UpdatedAt:       entity.UpdatedAt,
	}
}

func (t *Transaction) ToEntity() *entity.Transaction {
	return &entity.Transaction{
		ID:              t.ID,
		UserID:          t.UserID,
		AssetType:       t.AssetType,
		Quantity:        t.Quantity,
		Price:           t.Price,
		TransactionType: entity.TransactionType(t.TransactionType),
		Status:          entity.TransactionStatus(t.Status),
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
	}
}
