package persistence

import (
	"context"
	"database/sql"

	"goldvault/trading-service/internal/core/application/ports"
	"goldvault/trading-service/internal/core/domain/entity"
	"goldvault/trading-service/internal/infrastructure/persistence/models"
	"goldvault/trading-service/internal/infrastructure/persistence/queries"
	"goldvault/trading-service/pkg/serr"
)

type TransactionPersistence struct {
	db *sql.DB
}

func NewTransactionPersistence(db *sql.DB) ports.TransactionPersistence {
	return &TransactionPersistence{db: db}
}

func (t *TransactionPersistence) CreateTransaction(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) error {
	transactionDB := models.ToTransactionDB(transaction)
	result, err := tx.ExecContext(
		ctx,
		queries.InsertTransaction,
		transactionDB.UserID,
		transactionDB.AssetType,
		transactionDB.Quantity,
		transactionDB.Price,
		transactionDB.TransactionType,
		transactionDB.Status,
	)
	if err != nil {
		return serr.DBError("CreateTransaction", "transactions", err)
	}

	transactionID, err := result.LastInsertId()
	if err != nil {
		return serr.DBError("CreateTransaction", "transactions", err)
	}

	transaction.ID = transactionID

	return nil
}

func (t *TransactionPersistence) GetTransactionByUserID(ctx context.Context, userID int64) ([]*entity.Transaction, error) {
	rows, err := t.db.QueryContext(ctx, queries.GetUserTransactions, userID)
	if err != nil {
		return nil, serr.DBError("GetTransactionByUserID", "transactions", err)
	}
	defer rows.Close()

	var transactions []*entity.Transaction
	for rows.Next() {
		var transactionDB models.Transaction
		err := transactionDB.Scan(rows)
		if err != nil {
			return nil, serr.DBError("GetTransactionByUserID", "transactions", err)
		}
		transactions = append(transactions, transactionDB.ToEntity())
	}
	return transactions, nil
}

func (t *TransactionPersistence) UpdateTransactionStatus(ctx context.Context, transactionID int64, status string) error {
	_, err := t.db.ExecContext(ctx, queries.UpdateTransactionStatus, status, transactionID)
	if err != nil {
		return serr.DBError("UpdateTransactionStatus", "transactions", err)
	}
	return nil
}

func (t *TransactionPersistence) GetTransactions(ctx context.Context, limit, offset int64) ([]*entity.Transaction, error) {
	rows, err := t.db.QueryContext(ctx, queries.GetAllTransactions, limit, offset)
	if err != nil {
		return nil, serr.DBError("GetTransactions", "transactions", err)
	}
	defer rows.Close()

	var transactions []*entity.Transaction
	for rows.Next() {
		var transactionDB models.Transaction
		err := transactionDB.Scan(rows)
		if err != nil {
			return nil, serr.DBError("GetTransactions", "transactions", err)
		}
		transactions = append(transactions, transactionDB.ToEntity())
	}
	return transactions, nil
}
