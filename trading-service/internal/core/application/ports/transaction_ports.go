package ports

import (
	"context"
	"database/sql"

	"goldvault/trading-service/internal/core/domain/entity"
)

type (
	TransactionPersistence interface {
		CreateTransaction(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) error
		GetTransactionByUserID(ctx context.Context, userID int64) ([]*entity.Transaction, error)
		UpdateTransactionStatus(ctx context.Context, transactionID int64, status string) error
		GetTransactions(ctx context.Context, limit, offset int64) ([]*entity.Transaction, error)
	}

	TransactionDomainService interface {
		LogTransaction(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) error
		UpdateTransactionStatus(ctx context.Context, transactionID int64, status string) error
		GetUserTransaction(ctx context.Context, userID int64) ([]*entity.Transaction, error)
		GetTransactions(ctx context.Context, limit, offset int64) ([]*entity.Transaction, error)
	}
)
