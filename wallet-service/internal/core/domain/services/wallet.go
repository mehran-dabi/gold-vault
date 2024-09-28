package services

import (
	"context"
	"time"

	"goldvault/wallet-service/internal/core/application/ports"
	"goldvault/wallet-service/internal/core/domain/entity"
)

type WalletService struct {
	walletPersistence ports.WalletPersistence
}

func NewWalletDomainService(walletPersistence ports.WalletPersistence) ports.WalletDomainService {
	return &WalletService{walletPersistence: walletPersistence}
}

func (w *WalletService) CreateWallet(ctx context.Context, userID int64) (*entity.Wallet, error) {
	wallet, err := w.walletPersistence.GetWalletByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if wallet != nil {
		return wallet, nil
	}

	newWallet := &entity.Wallet{
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = w.walletPersistence.SaveWallet(ctx, newWallet)
	if err != nil {
		return nil, err
	}

	return newWallet, nil
}

func (w *WalletService) GetUserWallet(ctx context.Context, userID int64) (*entity.Wallet, error) {
	return w.walletPersistence.GetWalletByUserID(ctx, userID)
}

func (w *WalletService) GetWalletsWithPagination(ctx context.Context, limit, offset int) ([]*entity.Wallet, error) {
	return w.walletPersistence.GetWalletsWithPagination(ctx, limit, offset)
}
