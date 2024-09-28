package services

import (
	"context"

	"goldvault/wallet-service/internal/core/application/ports"
	"goldvault/wallet-service/internal/core/domain/entity"
)

type TransactionService struct {
	transactionPersistence ports.TransactionPersistence
	walletPersistence      ports.WalletPersistence
}

func NewTransactionDomainService(
	transactionPersistence ports.TransactionPersistence,
	walletPersistence ports.WalletPersistence,
) ports.TransactionDomainService {
	return &TransactionService{
		transactionPersistence: transactionPersistence,
		walletPersistence:      walletPersistence,
	}
}

func (t *TransactionService) GetWalletTransactions(ctx context.Context, userID int64) ([]entity.Transaction, error) {
	// check if the wallet belongs to the user
	wallet, err := t.walletPersistence.GetWalletByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return t.transactionPersistence.GetTransactionsByWalletID(ctx, wallet.ID)
}

func (t *TransactionService) GetTransactionsWithPagination(ctx context.Context, limit, offset int) ([]entity.Transaction, error) {
	return t.transactionPersistence.GetTransactionsWithPagination(ctx, limit, offset)
}
