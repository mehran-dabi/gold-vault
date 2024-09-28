package persistence

import (
	"context"
	"database/sql"

	"goldvault/wallet-service/internal/core/application/ports"
	"goldvault/wallet-service/internal/core/domain/entity"
	"goldvault/wallet-service/internal/infrastructure/persistence/models"
	"goldvault/wallet-service/internal/infrastructure/persistence/queries"
	"goldvault/wallet-service/pkg/serr"
)

type TransactionPersistence struct {
	db *sql.DB
}

func NewTransactionPersistence(db *sql.DB) ports.TransactionPersistence {
	return &TransactionPersistence{db: db}
}

func (t *TransactionPersistence) AddTransaction(ctx context.Context, transaction *entity.Transaction) error {
	dbModel, err := models.ToTransactionDB(transaction)
	if err != nil {
		return serr.DBError("AddTransaction", "transaction", err)
	}

	err = t.db.QueryRowContext(ctx, queries.AddTransaction,
		dbModel.AssetID,
		dbModel.WalletID,
		dbModel.Amount,
		dbModel.Status,
		dbModel.CreatedAt,
	).Scan(&dbModel.ID)
	if err != nil {
		return serr.DBError("AddTransaction", "transaction", err)
	}

	// Update the entity ID after successful creation
	transaction.ID = dbModel.ID

	return nil
}

func (t *TransactionPersistence) GetTransactionsByAssetID(ctx context.Context, assetID int64) ([]entity.Transaction, error) {
	rows, err := t.db.QueryContext(ctx, queries.GetTransactionsByAssetID, assetID)
	if err != nil {
		return nil, serr.DBError("GetTransactionsByAssetID", "transaction", err)
	}
	defer rows.Close()

	transactions := make([]entity.Transaction, 0)
	for rows.Next() {
		var transaction models.Transaction
		err := transaction.Scan(rows)
		if err != nil {
			return nil, serr.DBError("GetTransactionsByAssetID", "transaction", err)
		}
		transactions = append(transactions, *transaction.ToTransactionEntity())
	}
	return transactions, nil
}

func (t *TransactionPersistence) GetTransactionsByWalletID(ctx context.Context, walletID int64) ([]entity.Transaction, error) {
	rows, err := t.db.QueryContext(ctx, queries.GetTransactionsByWalletID, walletID)
	if err != nil {
		return nil, serr.DBError("GetTransactionsByWalletID", "transaction", err)
	}
	defer rows.Close()

	transactions := make([]entity.Transaction, 0)
	for rows.Next() {
		var transaction models.Transaction
		err := transaction.Scan(rows)
		if err != nil {
			return nil, serr.DBError("GetTransactionsByWalletID", "transaction", err)
		}
		transactions = append(transactions, *transaction.ToTransactionEntity())
	}
	return transactions, nil
}

func (t *TransactionPersistence) GetTransactionsWithPagination(ctx context.Context, limit, offset int) ([]entity.Transaction, error) {
	rows, err := t.db.QueryContext(ctx, queries.GetTransactionsWithPagination, limit, offset)
	if err != nil {
		return nil, serr.DBError("GetTransactionsWithPagination", "transaction", err)
	}
	defer rows.Close()

	transactions := make([]entity.Transaction, 0)
	for rows.Next() {
		var transaction models.Transaction
		err := transaction.Scan(rows)
		if err != nil {
			return nil, serr.DBError("GetTransactionsWithPagination", "transaction", err)
		}
		transactions = append(transactions, *transaction.ToTransactionEntity())
	}
	return transactions, nil
}
