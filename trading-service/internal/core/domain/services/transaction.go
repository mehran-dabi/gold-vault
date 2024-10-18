package services

import (
	"context"
	"database/sql"
	"time"

	"goldvault/trading-service/internal/core/application/ports"
	"goldvault/trading-service/internal/core/domain/entity"
)

type TransactionDomainService struct {
	transactionPersistence ports.TransactionPersistence
}

func NewTransactionDomainService(transactionPersistence ports.TransactionPersistence) ports.TransactionDomainService {
	return &TransactionDomainService{transactionPersistence: transactionPersistence}
}

func (t *TransactionDomainService) LogTransaction(ctx context.Context, tx *sql.Tx, transaction *entity.Transaction) error {
	if err := transaction.Validate(); err != nil {
		return err
	}

	err := t.transactionPersistence.CreateTransaction(ctx, tx, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionDomainService) UpdateTransactionStatus(ctx context.Context, transactionID int64, status string) error {
	err := t.transactionPersistence.UpdateTransactionStatus(ctx, transactionID, status)
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionDomainService) GetUserTransaction(ctx context.Context, userID int64) ([]*entity.Transaction, error) {
	transactions, err := t.transactionPersistence.GetTransactionByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *TransactionDomainService) GetTransactions(ctx context.Context, limit, offset int64) ([]*entity.Transaction, error) {
	transactions, err := t.transactionPersistence.GetTransactions(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *TransactionDomainService) GetSingleDaySummary(ctx context.Context, date time.Time, assetType string) (*entity.TransactionsSummary, error) {
	summary, err := t.transactionPersistence.GetTransactionSummaryForSingleDay(ctx, date, assetType)
	if err != nil {
		return nil, err
	}

	return summary, nil
}

func (t *TransactionDomainService) GetTotalSummary(ctx context.Context, assetType string) (*entity.TransactionsSummary, error) {
	summary, err := t.transactionPersistence.GetTotalTransactionsSummary(ctx, assetType)
	if err != nil {
		return nil, err
	}

	return summary, nil

}
