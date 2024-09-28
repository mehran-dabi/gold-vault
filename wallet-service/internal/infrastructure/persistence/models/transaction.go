package models

import (
	"time"

	"goldvault/wallet-service/internal/core/domain/entity"
)

type Transaction struct {
	ID          int64
	WalletID    int64
	AssetID     int64
	Type        string
	Amount      float64
	Description string
	Status      string
	CreatedAt   time.Time
}

const (
	TransactionColumns     = "id, wallet_id, asset_id, type, amount, description, created_at"
	TransactionColumnsNoID = "wallet_id, asset_id, type, amount, description, created_at"
)

func ToTransactionDB(transaction *entity.Transaction) (*Transaction, error) {
	return &Transaction{
		ID:          transaction.ID,
		WalletID:    transaction.WalletID,
		AssetID:     transaction.AssetID,
		Type:        string(transaction.Type),
		Amount:      transaction.Amount,
		Description: transaction.Description,
		Status:      string(transaction.Status),
		CreatedAt:   transaction.CreatedAt,
	}, nil
}

func (t *Transaction) ToTransactionEntity() *entity.Transaction {
	return &entity.Transaction{
		ID:          t.ID,
		WalletID:    t.WalletID,
		AssetID:     t.AssetID,
		Type:        entity.TxType(t.Type),
		Amount:      t.Amount,
		Description: t.Description,
		Status:      entity.TxStatus(t.Status),
		CreatedAt:   t.CreatedAt,
	}
}

func (t *Transaction) Scan(scanner Scanner) error {
	return scanner.Scan(
		&t.ID,
		&t.WalletID,
		&t.AssetID,
		&t.Type,
		&t.Amount,
		&t.Description,
		&t.CreatedAt,
	)
}
