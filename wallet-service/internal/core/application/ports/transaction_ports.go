package ports

import (
	"context"

	"goldvault/wallet-service/internal/core/domain/entity"
)

type (
	TransactionPersistence interface {
		AddTransaction(ctx context.Context, transaction *entity.Transaction) error
		GetTransactionsByAssetID(ctx context.Context, assetID int64) ([]entity.Transaction, error)
		GetTransactionsByWalletID(ctx context.Context, walletID int64) ([]entity.Transaction, error)
		GetTransactionsWithPagination(ctx context.Context, limit, offset int) ([]entity.Transaction, error)
	}

	TransactionDomainService interface {
		GetWalletTransactions(ctx context.Context, userID int64) ([]entity.Transaction, error)
		GetTransactionsWithPagination(ctx context.Context, limit, offset int) ([]entity.Transaction, error)
	}
)
