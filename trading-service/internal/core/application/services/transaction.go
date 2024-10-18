package services

import (
	"context"
	"net/http"
	"time"

	"goldvault/trading-service/internal/core/application/ports"
	"goldvault/trading-service/internal/core/domain/entity"
	"goldvault/trading-service/pkg/serr"
)

type TransactionService struct {
	transactionDomain ports.TransactionDomainService
}

func NewTransactionService(transactionDomain ports.TransactionDomainService) *TransactionService {
	return &TransactionService{transactionDomain: transactionDomain}
}

func (t *TransactionService) GetUserTransactions(ctx context.Context, userID int64) ([]*entity.Transaction, error) {
	txs, err := t.transactionDomain.GetUserTransaction(ctx, userID)
	if err != nil {
		return nil, serr.ServiceErr("TransactionService.GetUserTransactions", err.Error(), err, http.StatusInternalServerError)
	}

	return txs, nil
}

func (t *TransactionService) GetTransactions(ctx context.Context, limit, offset int64) ([]*entity.Transaction, error) {
	txs, err := t.transactionDomain.GetTransactions(ctx, limit, offset)
	if err != nil {
		return nil, serr.ServiceErr("TransactionService.GetTransactions", err.Error(), err, http.StatusInternalServerError)
	}

	return txs, nil
}

func (t *TransactionService) GetSingleDaySummary(ctx context.Context, date time.Time, assetType string) (*entity.TransactionsSummary, error) {
	summary, err := t.transactionDomain.GetSingleDaySummary(ctx, date, assetType)
	if err != nil {
		return nil, serr.ServiceErr("TransactionService.GetSingleDaySummary", err.Error(), err, http.StatusInternalServerError)
	}

	return summary, nil
}

func (t *TransactionService) GetTotalSummary(ctx context.Context, assetType string) (*entity.TransactionsSummary, error) {
	summary, err := t.transactionDomain.GetTotalSummary(ctx, assetType)
	if err != nil {
		return nil, serr.ServiceErr("TransactionService.GetTotalSummary", err.Error(), err, http.StatusInternalServerError)
	}

	return summary, nil
}
