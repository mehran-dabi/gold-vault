package ports

import (
	"context"
	"database/sql"
	"time"

	"goldvault/trading-service/internal/core/domain/entity"
)

type (
	TransactionPersistence interface {
		CreateTransaction(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) error
		GetTransactionByUserID(ctx context.Context, userID int64) ([]*entity.Transaction, error)
		UpdateTransactionStatus(ctx context.Context, transactionID int64, status string) error
		GetTransactions(ctx context.Context, limit, offset int64) ([]*entity.Transaction, error)
		GetTransactionSummaryForSingleDay(ctx context.Context, date time.Time, assetType string) (*entity.TransactionsSummary, error)
		GetTotalTransactionsSummary(ctx context.Context, assetType string) (*entity.TransactionsSummary, error)
	}

	TransactionDomainService interface {
		LogTransaction(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) error
		UpdateTransactionStatus(ctx context.Context, transactionID int64, status string) error
		GetUserTransaction(ctx context.Context, userID int64) ([]*entity.Transaction, error)
		GetTransactions(ctx context.Context, limit, offset int64) ([]*entity.Transaction, error)
		GetSingleDaySummary(ctx context.Context, date time.Time, assetType string) (*entity.TransactionsSummary, error)
		GetTotalSummary(ctx context.Context, assetType string) (*entity.TransactionsSummary, error)
	}
)
