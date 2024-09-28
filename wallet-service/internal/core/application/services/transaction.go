package services

import (
	"context"
	"fmt"
	"net/http"

	"goldvault/wallet-service/internal/core/application/ports"
	"goldvault/wallet-service/internal/core/domain/entity"
	"goldvault/wallet-service/pkg/serr"
)

type TransactionService struct {
	transactionDomainService ports.TransactionDomainService
}

func NewTransactionService(transactionDomainService ports.TransactionDomainService) *TransactionService {
	return &TransactionService{transactionDomainService: transactionDomainService}
}

func (t *TransactionService) GetWalletTransactions(ctx context.Context, userID int64) ([]entity.Transaction, error) {
	fmt.Println("context: ", ctx, "userID: ", userID, "t", t)
	transactions, err := t.transactionDomainService.GetWalletTransactions(ctx, userID)
	if err != nil {
		return nil, serr.ServiceErr("TransactionApplicationService.GetWalletTransactions", err.Error(), err, http.StatusInternalServerError)
	}

	return transactions, nil
}

func (t *TransactionService) GetTransactionsWithPagination(ctx context.Context, limit, offset int) ([]entity.Transaction, error) {
	transactions, err := t.transactionDomainService.GetTransactionsWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, serr.ServiceErr("TransactionApplicationService.GetTransactionsWithPagination", err.Error(), err, http.StatusInternalServerError)
	}

	return transactions, nil
}
